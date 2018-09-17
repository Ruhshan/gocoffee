package app

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func Serve() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	dbi := DBInterface{}
	dbi.InitDB()
	defer dbi.DB.Close()

	router, err := rest.MakeRouter(
		rest.Get("/product/list/", dbi.GetProductList),
		rest.Get("/product/:id/", dbi.GetProduct),
		rest.Post("/product/", dbi.PostProduct),
		rest.Post("/customer/", dbi.PostCustomer),
	)
	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8000", api.MakeHandler()))

}
