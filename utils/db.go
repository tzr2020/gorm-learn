package utils

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "root"
	password = "123456"
	host     = "127.0.0.1"
	port     = "3306"
	dbname   = "test"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// 打开数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(20)           // 设置空闲连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间

	// 测试数据库连接
	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}

	log.Println("数据库连接成功")
}
