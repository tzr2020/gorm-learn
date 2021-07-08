package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassName string
	// 一对多关系，Student模型的ClassID字段为外键，引用Class模型的ID字段
	Students []Student
}

type Student struct {
	gorm.Model
	StudentName string
	// 一对一关系，Card模型的StudentID字段为外键，引用Student模型的ID字段
	Card Card
	// 一对多关系
	ClassID uint
	// 多对多关系
	Teachers []Teacher `gorm:"many2many:student_teachers;"`
}

type Card struct {
	gorm.Model
	CardName string
	// 一对一关系
	StudentID uint
}

type Teacher struct {
	gorm.Model
	TeacherName string
	// 多对多关系
	Students []Student `gorm:"many2many:student_teachers;"`
}

var (
	db  *gorm.DB
	err error
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Class{}, &Student{}, &Card{}, &Teacher{})
	// 自动创建student_teachers连接表，该表有student_id和teache_id字段为外键，分别引用Student模型和Teacher模型的ID字段

	// CreateStudent()

	CRUDAPI()
}

// 创建学生记录
func CreateStudent() {
	card := Card{
		CardName: "123456",
	}

	teac := Teacher{
		TeacherName: "lisi",
	}

	stud := Student{
		StudentName: "zhangsan",
		Card:        card,
		Teachers:    []Teacher{teac},
	}

	clas := Class{
		ClassName: "2101",
		Students:  []Student{stud},
	}

	db.Create(&clas)
}

// CRUD接口
func CRUDAPI() {
	r := gin.Default()

	// {
	// 	"StudentName":"wangwu",
	// 	"ClassID":1,
	// 	"Card":{
	// 	  	"CardName": "100000"
	// 	},
	// 	"Teachers":[{
	// 	  	"TeacherName":"老师1"
	// 	},{
	// 	  	"TeacherName":"老师2"
	// 	}]
	// }
	r.POST("/student", func(c *gin.Context) {
		var stud Student
		err := c.ShouldBindJSON(&stud)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "请求参数错误",
			})
			return
		}

		err = db.Create(&stud).Error
		if err != nil {
			c.JSON(200, gin.H{
				"message": "服务器内部错误",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "新增学生成功",
		})
	})

	r.GET("/student/:id", func(c *gin.Context) {
		id := c.Param("id")

		var stud Student
		// err := db.First(&stud, "id = ?", id).Error
		err := db.Preload("Card").Preload("Teachers").First(&stud, "id = ?", id).Error // 预加载
		if err != nil {
			c.JSON(200, gin.H{
				"message": "服务器内部错误",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "查询学生成功",
			"data":    stud,
		})
	})

	r.GET("/class/:id", func(c *gin.Context) {
		id := c.Param("id")

		var clas Class
		// err := db.Preload("Students").First(&clas, "id = ?", id).Error // 预加载
		err := db.Preload("Students").Preload("Students.Card").Preload("Students.Teachers").First(&clas, "id = ?", id).Error // 嵌套预加载
		if err != nil {
			c.JSON(200, gin.H{
				"message": "服务器内部错误",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "查询班级成功",
			"data":    clas,
		})
	})

	r.Run()
}
