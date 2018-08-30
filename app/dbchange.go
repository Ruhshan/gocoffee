package app

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Migrate() {
	db, err := gorm.Open("postgres", Connection)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	fmt.Println("Running Migration for product")
	db.AutoMigrate(&Product{})

}

func DeleteField(model string, field string) {
	field = strings.ToLower(field)

	db, err := gorm.Open("postgres", Connection)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	switch model {

	case "Product":
		err = db.Model(&Product{}).DropColumn(field).Error
		if err != nil {
			panic(err)
		}

	}

}
