package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// 创建表
func create(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

func insert(db *gorm.DB) {
	db.Create(&Product{Code: "D42", Price: 100})
}
func find(db *gorm.DB) {
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Get first matched record
	db.Take(&product, "code = ?", "D42") // find product with code D42
	fmt.Println(product)
}

func update(db *gorm.DB, product Product) {
	db.Model(&product).Update("Price", 200)
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
}

func delete(db *gorm.DB, product Product) {
	db.Delete(&product, 1)
}

func main() {
	dsn := "root:5527379@tcp(localhost:3306)/fish-master?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// create(db)
	// insert(db)
	find(db)
	// update(db, Product{Code: "D42", Price: 100})
	// delete(db, Product{Code: "D42", Price: 100})
}
