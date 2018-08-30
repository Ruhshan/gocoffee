package app

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func Serve() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/product/list/", GetProductList),
		rest.Get("/product/:id/", GetProduct),
		rest.Post("/product/", PostProduct),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8000", api.MakeHandler()))

}
