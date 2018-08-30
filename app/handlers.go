package app

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jinzhu/gorm"
)

type HandlerDBInterface struct {
	DB *gorm.DB
}

func (i *HandlerDBInterface) InitDB() {
	var err error
	i.DB, err = gorm.Open("postgres", Connection)
	if err != nil {
		panic("failed to connect database")
	}

}

func (i *HandlerDBInterface) GetProductList(w rest.ResponseWriter, r *rest.Request) {

	products := []Product{}
	i.DB.Find(&products)

	w.WriteJson(&products)

}

func (i *HandlerDBInterface) GetProduct(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")

	product := Product{}

	if i.DB.First(&product, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&product)
}

func (i *HandlerDBInterface) PostProduct(w rest.ResponseWriter, r *rest.Request) {

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
