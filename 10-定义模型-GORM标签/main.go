package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义模型，使用GORM标签
type User struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"column:username;type:varchar(50);not null"`
}

// 指定数据库表名
func (u User) TableName() string {
	return "user"
}

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动迁移模式，根据模型自动创建数据库表
	db.AutoMigrate(&User{})

	// 新增记录
	db.Create(&User{
		Name: "zhangsan",
	})
}
