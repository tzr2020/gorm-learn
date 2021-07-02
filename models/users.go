package models

import (
	"gorm-learn/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Age      uint8
}

// 等效
// type User struct {
// 	ID        uint `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// 	Username  string         `gorm:"unique"`
// 	Password  string
// 	Age       uint8
// }

// 迁移模式
func init() {
	utils.DB.AutoMigrate(&User{})
}
