package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	// create
	//db.Create(&Product{Name: "Laptop", Price: 1000})

	// create batch
	//products := []Product{
	//	{Name: "Mouse", Price: 10},
	//	{Name: "Keyboard", Price: 20},
	//	{Name: "Monitor", Price: 200},
	//}
	//db.Create(&products)

	// select one
	//var product Product
	//db.First(&product, 1)
	//fmt.Println(product)

	// select with where
	//var product Product
	//db.First(&product, "name = ?", "Mouse")
	//fmt.Println(product)

	// select all
	//var products []Product
	//db.Find(&products)
	//fmt.Println(products)

	// select with limit
	//var products []Product
	//db.Limit(2).Find(&products)
	//fmt.Println(products)

	// select with offset to pagination
	//var products []Product
	//db.Offset(2).Limit(2).Find(&products)
	//fmt.Println(products)

	// where
	var products []Product
	db.Where("price >= ?", 200).Find(&products)
	fmt.Println(products)
}
