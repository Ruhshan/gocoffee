package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ant0ine/go-json-rest/rest"
)

func (i *DBInterface) GetProductList(w http.ResponseWriter, r *http.Request) {

	products := []Product{}
	i.DB.Find(&products)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

func (i *DBInterface) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	product := Product{}

	if err := i.DB.First(&product, id).Error; err != nil {

		http.Error(w, `{"Error": "not found"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (i *DBInterface) PostProduct(w rest.ResponseWriter, r *rest.Request) {

	product := Product{}

	if err := r.DecodeJsonPayload(&product); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := i.DB.Save(&product).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&product)
}
