package conf

import (
	"fmt"
	"time"
	"qqlx/base/constant"

	"github.com/spf13/viper"
)

func GetLogFormat() string {
	return viper.GetString("server.logFormat")
}
func GetLogLevel() string {
	return viper.GetString("server.logLevel")
}

// Server
func GetServerBind() string {
	return viper.GetString("server.bind")
}

func GetProjectName() string {
	return viper.GetString("server.projectName")
}

// JWT
func GetJwtSecret() string {
	return viper.GetString("jwt.secret")
}

func GetJwtIssuer() string {
	return viper.GetString("jwt.issuer")
}

func GetJwtExpirationTime() time.Duration {
	timeOut := viper.GetString("jwt.expirationTime")
	expir, err := time.ParseDuration(timeOut)
	if err != nil {
		expir, _ = time.ParseDuration(constant.DefaultJwtExpireTime)
	}
	return expir
}

// Mysql
func GetCasbinDsn() string {
	user := viper.GetString("mysql.username")
	pas := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	database := viper.GetString("mysql.database")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pas, host, port, database)
}

func GetMysqlDsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=10000ms",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.database"),
	)
	return dsn
}

// Redis
func GetRdisPassword() string {
	return viper.GetString("redis.password")
}
func GetRdisMasterName() string {
	return viper.GetString("redis.sentinel.masterName")
}
func GetRdisSentinelPassword() string {
	return viper.GetString("redis.sentinel.password")
}
func GetRdisSentinelHosts() []string {
	return viper.GetStringSlice("redis.sentinel.hosts")
}

func GetRdisHost() string {
	return viper.GetString("redis.host")
}
func GetRdisPort() string {
	return viper.GetString("redis.port")
}
func GetRdisDB() int {
	return viper.GetInt("redis.db")
}

func GetRdisMode() string {
	return viper.GetString("redis.mode")
}

func GetRedisExpireTime() (time.Duration, error) {
	expireTime := viper.GetString("redis.expireTime")
	if expireTime == "" {
		expireTime = constant.DefaultRedisExpireTime
	}
	return time.ParseDuration(expireTime)
}

func GetRedisKeyPrefix() string {
	return viper.GetString("redis.keyPrefix")
}
