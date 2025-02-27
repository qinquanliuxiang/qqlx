package test

import (
	"context"
	"qqlx/base"
	"qqlx/base/conf"
	"qqlx/base/data"
	"qqlx/pkg/permissions"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Ctx      = context.Background()
	Enforcer *casbin.Enforcer
	Cache    base.Cache
	Close1   func()
	DB       *gorm.DB
	Close2   func()
)

func init() {
	var (
		rdb      *redis.Client
		redisCli *data.Redis
		err      error
	)
	conf.LoadConfig("../../config.yaml")
	Enforcer, err = permissions.InitCasbin("../../rbac_model.conf")
	if err != nil {
		panic(err)
	}
	rdb = data.CreateRDB(Ctx)
	redisCli, Close1 = data.NewRedis(rdb)
	Cache = redisCli
	DB, Close2, err = data.NewDB()
	if err != nil {
		panic(err)
	}
}
