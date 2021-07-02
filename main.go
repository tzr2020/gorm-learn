package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 迁移模式
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "D42", Price: 100})

	// 查询
	var p Product
	db.First(&p, 1)
	db.First(&p, "code=?", "D42")

	// 更新
	db.Model(&p).Update("Price", 200)
	db.Model(&p).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&p).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// 删除
	db.Delete(&p, 1)
}
