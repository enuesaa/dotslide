package main

import (
	"log"

	"github.com/enuesaa/dotslide/web"
)


func main() {
	app := web.NewReouter()

	if err := app.Start(":3000"); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}
