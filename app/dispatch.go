package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {

	dbi := DBInterface{}
	dbi.InitDB()
	defer dbi.DB.Close()

	router := mux.NewRouter()

	router.HandleFunc("/product/list/", dbi.GetProductList).Methods("GET")
	router.HandleFunc("/product/{id}/", dbi.GetProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}
