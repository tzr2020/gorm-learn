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
	// f4()
	// f5()
	f6()
}

func f1() {
	var u models.User
	utils.DB.Take(&u)
	fmt.Printf("u, value: %v\n", u)
}

func f2() {
	var u models.User
	// utils.DB.First(&u)
	utils.DB.Last(&u)
	fmt.Printf("u, value: %v\n", u)
}

func f3() {
	var us []models.User
	utils.DB.Find(&us)
	for _, u := range us {
		fmt.Printf("u, value: %v\n", u)
	}
}

func f4() {
	// var u models.User
	var us []models.User

	// utils.DB.Where("username = ?", "zhangsan").First(&u)
	// utils.DB.Where("username in ?", []string{"lisi", "zhangsan"}).Find(&us)
	// utils.DB.Where("username like ?", "li%").Find(&us)
	// utils.DB.Where("username like ? and age > ?", "li%", "10").Find(&us)
	utils.DB.Where("age between ? and ?", "11", "12").Find(&us)

	// fmt.Printf("u, value: %v\n", u)
	for _, u := range us {
		fmt.Printf("u, value: %v\n", u)
	}
}

func f5() {
	var u models.User
	var us []models.User

	// utils.DB.Where(map[string]interface{}{"Username": "zhangsan", "Age": 11}).First(&u)
	// utils.DB.Where(&models.User{Username: "zhangsan", Age: 11}).First(&u)
	utils.DB.Where(&models.User{Username: "zhangsan", Age: 11}, "username").First(&u)

	utils.DB.Where([]int{1, 2, 3}).Find(&us)

	fmt.Printf("u, value: %v\n", u)
	for _, u := range us {
		fmt.Printf("u, value: %v\n", u)
	}
}

func f6() {
	var us []models.User

	// utils.DB.Not("age = ?", "0").Find(&us)
	// utils.DB.Not(map[string]interface{}{"username": []string{"zhangsan", "lisi"}}).Find(&us)
	// utils.DB.Not(&models.User{Username: "zhangsan", Age: 11}).Find(&us)
	utils.DB.Not([]int{1, 2, 3}).Find(&us)

	for _, u := range us {
		fmt.Printf("u, value: %v\n", u)
	}
}
