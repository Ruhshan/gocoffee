package app

import (
	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
)

func ProductListHandler(w rest.ResponseWriter, r *rest.Request) {

	db, err := gorm.Open("postgres", Connection)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	products := []Product{}

	db.Find(&products)

	w.WriteJson(&products)

}
