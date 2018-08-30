package app

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func ProductListHandler(w rest.ResponseWriter, r *rest.Request) {

	dbi := DBInterface{}
	dbi.InitDB()
	defer dbi.DB.Close()

	products := []Product{}
	dbi.DB.Find(&products)

	w.WriteJson(&products)

}
