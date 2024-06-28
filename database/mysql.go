package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect 连接数据库
func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}
