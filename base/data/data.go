package data

import (
	"fmt"
	"qqlx/base/conf"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() (*gorm.DB, func(), error) {
	dsn := conf.GetMysqlDsn()
	var DBLogger logger.Interface
	// 开启mysql日志
	if viper.GetBool("mysql.debug") {
		zap.S().Debug("enable debug mode on the database")
		DBLogger = logger.Default.LogMode(logger.Info)
	}

	dbInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   DBLogger,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("exception in initializing mysql database, %w", err)
	}

	// 确保数据库连接已建立
	sqlDB, err := dbInstance.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to obtain database connection, %w", err)
	}

	// 尝试Ping数据库以确保连接有效
	err = sqlDB.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to obtain database connection, %w", err)
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
	// 设置数据库的最大打开连接数
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))

	zap.S().Info("mysql database initialization completed")
	return dbInstance,
		func() {
			sqlDB.Close()
		},
		nil
}
