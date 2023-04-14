/* orm简介
对象关系映射（Object Relational Mapping, 简程ORM）模式是一种为
了解决面向对象与关系数据库（如mysql数据库）
存在的互不匹配的现象技术。简单的说，ORM是通过使用描述对象
和数据库之间映射的元素据,将程序中的对象自动持久化到关系数据库中 */

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
	dsn := "root:880818@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//迁移schema
	db.AutoMigrate(&Product{})

	//create
	db.Create(&Product{Code: "D42", Price: 100})

	/*	//Read
		var product Product
		db.First(&product, 1)               //根据整型主键查找
		db.First(&product, "code=?", "D42") //查找code字段为D42的记录
		fmt.Printf("product: %v\n", product)

			/Update-将product的price更新为200
			db.Model(&product).Update("Price", 200)
			/Update-更新多个字段
			db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) //仅更新非零值字段
			db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

			    /Delete-删除product
			    db.Delete(&product, 1) */

}

/* //自定义驱动
db, err := gorm.Open(mysql.New(mysql.Config{
	DriverName:"my_mysql_driver"
    DSN : "root:880818@tcp(127.0.0.1:9910)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
}), &gorm.Config{})

/Data Source Name,参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
*/
