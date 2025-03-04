package data

import (
	"context"
	"fmt"
	"qqlx/base/apierrs"
	"qqlx/base/conf"
	"time"

	"github.com/redis/go-redis/v9"
)

var NeverExpires time.Duration = 0

// Redis redis 客户端
type Redis struct {
	client     *redis.Client
	expireTime time.Duration
	keyPrefix  string
}

func NewRedis(client *redis.Client) (*Redis, func()) {
	expireTime, err := conf.GetRedisExpireTime()
	if err != nil {
		panic(fmt.Sprintf("get redis expire time faild, err: %s", err.Error()))
	}

	claeup := func() {
		client.Close()
	}
	return &Redis{
		client:     client,
		expireTime: expireTime,
		keyPrefix:  conf.GetRedisKeyPrefix(),
	}, claeup
}

func CreateRDB(ctx context.Context) *redis.Client {
	switch conf.GetRdisMode() {
	case "sentinel":
		return initSentinelRedis(ctx)
	case "single":
		return initSingleRedis(ctx)
	default:
		panic("unsupported redis mode, please check the configuration redis.mode")
	}
}

func initSingleRedis(ctx context.Context) *redis.Client {
	host := conf.GetRdisHost()
	prot := conf.GetRdisPort()
	address := fmt.Sprintf("%s:%s", host, prot)
	if address == "" {
		panic("redis address is empty")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: conf.GetRdisPassword(),
		DB:       conf.GetRdisDB(),
	})
	s := rdb.Ping(ctx).Err()
	if s != nil {
		panic(s)
	}
	return rdb
}

func initSentinelRedis(ctx context.Context) *redis.Client {
	sentinelHosts := conf.GetRdisSentinelHosts()

	if len(sentinelHosts) == 0 {
		panic("redis sentinel is empty")
	}
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       conf.GetRdisMasterName(),
		SentinelAddrs:    sentinelHosts,
		Password:         conf.GetRdisPassword(),
		SentinelPassword: conf.GetRdisSentinelPassword(),
		RouteByLatency:   true,
		DB:               conf.GetRdisDB(),
	})
	s := rdb.Ping(ctx).Err()
	if s != nil {
		panic(s)
	}
	return rdb
}

func (c *Redis) GetString(ctx context.Context, key string) (string, error) {
	saveKey := fmt.Sprintf("%s_%s", c.keyPrefix, key)
	if v, err := c.client.Get(ctx, saveKey).Result(); err != nil {
		return "", apierrs.NewRedisGetErr(fmt.Errorf("redis getting %s key failed: %w", saveKey, err))
	} else {
		return v, nil
	}
}

func (c *Redis) SetString(ctx context.Context, key string, value string, expireTime *time.Duration) error {
	saveKey := fmt.Sprintf("%s_%s", c.keyPrefix, key)
	if expireTime == nil {
		if err := c.client.Set(ctx, saveKey, value, c.expireTime).Err(); err != nil {
			return apierrs.NewRedisSetErr(fmt.Errorf("redis setting %s key failed: %w", saveKey, err))
		}
		return nil
	}
	if expireTime == &NeverExpires {
		if err := c.client.Set(ctx, saveKey, value, 0).Err(); err != nil {
			return apierrs.NewRedisSetErr(fmt.Errorf("redis setting %s key failed: %w", saveKey, err))
		}
		return nil
	}
	if err := c.client.Set(ctx, saveKey, value, *expireTime).Err(); err != nil {
		if err := c.client.Set(ctx, saveKey, value, 0).Err(); err != nil {
			return apierrs.NewRedisSetErr(fmt.Errorf("redis setting %s key failed: %w", saveKey, err))
		}
		return nil
	}
	return nil
}

func (c *Redis) GetInt64(ctx context.Context, key string) (int64, error) {
	saveKey := fmt.Sprintf("%s_%s", c.keyPrefix, key)
	v, err := c.client.Get(ctx, saveKey).Int64()
	if err != nil {
		return 0, apierrs.NewRedisGetErr(fmt.Errorf("redis getting %s key failed: %w", saveKey, err))
	}

	return v, nil
}

func (c *Redis) SetInt64(ctx context.Context, key string, value int64, expireTime *time.Duration) error {
	saveKey := fmt.Sprintf("%s_%s", c.keyPrefix, key)

	if expireTime == nil {
		if err := c.client.Set(ctx, saveKey, value, c.expireTime).Err(); err != nil {
			return apierrs.NewRedisSetErr(fmt.Errorf("redis setting %s key failed: %w", saveKey, err))
		}
		return nil
	}
	if expireTime == &NeverExpires {
		if err := c.client.Set(ctx, saveKey, value, 0).Err(); err != nil {
			return apierrs.NewRedisSetErr(fmt.Errorf("redis setting %s key failed: %w", saveKey, err))
		}
		return nil
	}
	if err := c.client.Set(ctx, saveKey, value, *expireTime).Err(); err != nil {
		return apierrs.NewRedisSetErr(fmt.Errorf("redis setting %s key failed: %w", saveKey, err))
	}
	return nil
}

func (c *Redis) Del(ctx context.Context, key string) error {
	saveKey := fmt.Sprintf("%s_%s", c.keyPrefix, key)
	if err := c.client.Del(ctx, saveKey).Err(); err != nil {
		return apierrs.NewRedisDelErr(fmt.Errorf("redis deleting %s key failed: %w", saveKey, err))
	}
	return nil
}

func (c *Redis) Flush(ctx context.Context) error {
	if err := c.client.FlushDB(ctx).Err(); err != nil {
		return apierrs.NewRedisFlushErr(fmt.Errorf("redis flushing failed: %w", err))
	}
	return nil
}
