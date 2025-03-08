package main

import (
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
	db.Create(&Product{Name: "Laptop", Price: 1000})

	// create batch
	products := []Product{
		{Name: "Mouse", Price: 10},
		{Name: "Keyboard", Price: 20},
		{Name: "Monitor", Price: 200},
	}
	db.Create(&products)
}
