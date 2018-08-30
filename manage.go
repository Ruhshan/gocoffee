package main

import (
	"os"

	App "../gocoffee/app"
)

func main() {

	cmd := os.Args[1:]

	if cmd[0] == "migrate" {
		App.Migrate()
	}

	if cmd[0] == "delete" {
		model := cmd[1]
		field := cmd[2]

		App.DeleteField(model, field)
	}

}
