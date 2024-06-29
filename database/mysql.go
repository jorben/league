package database

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"league/config"
	"league/log"
	"time"
)

var mysqlInstance *gorm.DB

// GetInstance 获取DB实例
func GetInstance() *gorm.DB {
	return mysqlInstance
}

type MyGormLogger struct {
	LogLevel logger.LogLevel
}

// InitMysql 新建mysql连接实例
func InitMysql(cfg config.DbConfig) (err error) {

	mysqlInstance, err = gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{
		Logger: &MyGormLogger{LogLevel: logger.Info},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}

	sqlDB, err := mysqlInstance.DB()
	if err != nil {
		return
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)

	return sqlDB.Ping()
}

// Close 关闭mysql连接
func Close() {
	if sqlDB, err := mysqlInstance.DB(); err == nil {
		_ = sqlDB.Close()
	}

}

func (g MyGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := g
	newlogger.LogLevel = level
	return &newlogger
}

func (g MyGormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Info {
		log.Infof(s, i...)
	}

}

func (g MyGormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Warn {
		log.Warnf(s, i...)
	}
}

func (g MyGormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Error {
		log.Errorf(s, i...)
	}
}

func (g MyGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.LogLevel >= logger.Info {
		elapsed := time.Since(begin)
		sql, rows := fc()
		log.WithField(
			"Sql", sql,
			"Rows", rows,
			"Cost", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
		).Debug("SQL Trace")
	}
	return
}
