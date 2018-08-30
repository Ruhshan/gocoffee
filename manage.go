package main

import (
	"os"

	app "../gocoffee/app"
)

func main() {

	args := os.Args[1:]

	command := args[0]

	switch command {

	case "migrate":
		m := (*app.DBInterface).Migrate()

	case "delete":
		model := args[1]
		field := args[2]

		(DBInterface).DeleteField(model, field)
	case "runserver":
		App.Serve()

	}

}
