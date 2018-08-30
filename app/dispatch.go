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
		rest.Get("/products/", ProductListHandler),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8000", api.MakeHandler()))

}
