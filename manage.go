package main

import (
	"os"

	App "../gocoffee/app"
)

func main() {

	args := os.Args[1:]

	command := args[0]

	switch command {

	case "migrate":
		App.Migrate()

	case "delete":
		model := args[1]
		field := args[2]

		App.DeleteField(model, field)
	case "runserver":
		App.Serve()

	}

}
