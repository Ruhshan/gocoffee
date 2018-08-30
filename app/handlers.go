package app

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func ProductListHandler(w rest.ResponseWriter, r *rest.Request) {

	dbi := DBInterface{}
	dbi.InitDB()
	products := []Product{}
	dbi.DB.Find(&products)
	dbi.DB.Close()
	w.WriteJson(&products)

}
