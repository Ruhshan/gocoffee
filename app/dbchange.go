package app

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBInterface struct {
	DB *gorm.DB
}

func (i *DBInterface) InitDB() {
	var err error
	i.DB, err = gorm.Open("postgres", Connection)
	if err != nil {
		panic("failed to connect database")
	}

}

func (i *DBInterface) Migrate() {

	fmt.Println("Running Migration for product")
	i.DB.AutoMigrate(&Product{})

}

func (i *DBInterface) DeleteField(model string, field string) {

	switch model {

	case "Product":
		err := i.DB.Model(&Product{}).DropColumn(field).Error
		if err != nil {
			panic(err)
		}

	}

}
