package main

import (
	"os"

	app "../gocoffee/app"
)

func main() {

	args := os.Args[1:]

	command := args[0]

	dbi := app.DBInterface{}
	dbi.InitDB()
	defer dbi.DB.Close()

	switch command {

	case "migrate":

		dbi.Migrate()

	case "delete":
		/*
			go run manage.go Product customer
		*/
		model := args[1]
		field := args[2]

		dbi.DeleteField(model, field)
	case "runserver":
		app.Serve()

	}

}
