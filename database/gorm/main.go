package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primary_key"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product
}

type SerialNumber struct {
	ID        int `gorm:"primary_key"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	//category := Category{Name: "Electronics"}
	//db.Create(&category)

	// create product with category
	//db.Create(
	//	&Product{
	//		Name:       "Macbook",
	//		Price:      1000,
	//		CategoryID: category.ID,
	//	})

	//create serial number
	//db.Create(&SerialNumber{Number: "123456", ProductID: 2})

	// create
	//db.Create(&Product{Name: "Macbook", Price: 1000})

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
	//var products []Product
	//db.Where("price >= ?", 200).Find(&products)
	//fmt.Println(products)

	// update
	//var product Product
	//db.First(&product, 1)
	//product.Name = "Macbook"
	//db.Save(&product)

	//var p2 Product
	//db.First(&p2, 1)
	//fmt.Println(p2)

	// delete
	//db.Delete(&p2)

	//var product Product
	//db.First(&product, 1)
	//db.Delete(&product)

	// select products with categories
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println(" - ", product.Name, category.Name)
		}
	}
}
