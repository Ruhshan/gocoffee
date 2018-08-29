package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code    string
	Price   *int    `gorm:"not null"`
	Company *string `gorm:"not null"`
}

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gotest password=1234")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	defer db.Close()

	//db.AutoMigrate(&Product{})

	//db.Model(&Product{}).DropColumn("price")
	// if err != nil {
	// 	fmt.Println("what is error!")
	// }

	//Create
	company := "poogle"
	price := 10
	err = db.Create(&Product{Code: "L2217", Company: &company, Price: &price}).Error
	if err != nil {
		panic(err)
	}

}
