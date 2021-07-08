package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User模型
type User struct {
	gorm.Model
	Name string
	Sex  bool
	Age  int8
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

	// 增
	db.Create(&User{
		Name: "zhangsan",
		Sex:  true,
		Age:  11,
	})

	db.Create(&User{
		Name: "lisi",
		Sex:  false,
		Age:  12,
	})

	db.Create(&User{
		Name: "wangwu",
		Sex:  true,
		Age:  13,
	})

	// 查
	// var user User
	// db.First(&user)
	// db.First(&user, "name = ?", "lisi")
	// fmt.Println(user)

	// var users []User
	// db.Find(&users)
	// db.Find(&users, "age < ?", 13)
	// db.Where("age < ?", 13).Find(&users)
	// fmt.Println(users)

	// 改
	// db.Where("name = ?", "wangwu").First(&User{}).Update("age", 14)

	// db.Where("name = ?", "wangwu").First(&User{}).Updates(&User{
	// 	Sex: false,
	// 	Age: 15,
	// })
	// db.Where("name = ?", "wangwu").First(&User{}).Updates(map[string]interface{}{
	// 	"Sex": false,
	// 	"Age": 16,
	// })
	// 使用map可以使用零值修改数据库表记录，而使用struct则不行

	// db.Where("id in (?)", []int{1, 2}).Find(&[]User{}).Update("age", 18)

	// 删
	// 软删除，更新deleteed_at字段
	// db.Delete(&User{}, "id = ?", 1)
	// db.Where("id in (?)", []int{2, 3}).Delete(&User{})
	// 硬删除
	// db.Where("id in (?)", []int{2, 3}).Unscoped().Delete(&User{})
}
