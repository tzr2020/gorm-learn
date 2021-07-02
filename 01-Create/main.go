package main

import (
	"fmt"
	"gorm-learn/models"
	"gorm-learn/utils"
)

func main() {
	// f1()
	// f2()
	// f3()
	f4()
}

func f1() {
	u := models.User{
		Username: "zhangsan",
		Password: "123456",
		Age:      11,
	}

	res := utils.DB.Create(&u)

	fmt.Printf("ID: %v\n", u.ID)                       // 主键
	fmt.Printf("%v\n", res.Error)                      // 错误
	fmt.Printf("RowsAffected: %v\n", res.RowsAffected) // 影响行数
}

func f2() {
	u := models.User{
		Username: "lisi",
		Password: "123456",
		Age:      12,
	}
	utils.DB.Select("username", "age").Create(&u)
}

func f3() {
	u := models.User{
		Username: "wangwu",
		Password: "123456",
		Age:      13,
	}
	utils.DB.Omit("username", "age").Create(&u)
}

func f4() {
	var us = []models.User{{Username: "q"}, {Username: "w"}, {Username: "e"}}

	utils.DB.Create(&us)

	for _, u := range us {
		fmt.Printf("ID: %v\n", u.ID)
	}
}
