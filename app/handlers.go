package app

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func (i *DBInterface) GetProductList(w rest.ResponseWriter, r *rest.Request) {

	products := []Product{}
	// := []Customer{}

	i.DB.Preload("Customer").Find(&products)

	w.WriteJson(&products)

}

func (i *DBInterface) GetProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")

	product := Product{}

	if i.DB.Preload("Customer").First(&product, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&product)
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

func (i *DBInterface) PostCustomer(w rest.ResponseWriter, r *rest.Request) {

	customer := Customer{}

	if err := r.DecodeJsonPayload(&customer); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := i.DB.Save(&customer).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&customer)
}
