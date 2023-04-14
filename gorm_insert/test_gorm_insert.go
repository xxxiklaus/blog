package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:880818@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func createTable() {
	db.AutoMigrate(&User{})
}

func insert1() {
	user := User{
		Name:     "tom",
		Age:      20,
		Birthday: time.Now(),
	}
	result := db.Create(&user)                                   //通过数据的指针来创建
	fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected) //返回插入记录的条数
	fmt.Printf("user.ID: %v\n", user.ID)                         //返回插入数的主键
}

func main() {
	// createTable()
	insert1()
}
