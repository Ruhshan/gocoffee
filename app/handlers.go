package app

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetProductList(w rest.ResponseWriter, r *rest.Request) {

	dbi := DBInterface{}
	dbi.InitDB()
	defer dbi.DB.Close()

	products := []Product{}
	dbi.DB.Find(&products)

	w.WriteJson(&products)

}

func GetProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")

	dbi := DBInterface{}
	dbi.InitDB()
	defer dbi.DB.Close()

	product := Product{}

	if dbi.DB.First(&product, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&product)
}

func PostProduct(w rest.ResponseWriter, r *rest.Request) {

	dbi := DBInterface{}
	dbi.InitDB()
	defer dbi.DB.Close()

	product := Product{}

	if err := r.DecodeJsonPayload(&product); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbi.DB.Save(&product).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&product)
}
