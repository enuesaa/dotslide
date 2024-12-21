package main

import (
	"log"

	flag "github.com/spf13/pflag"
	"github.com/enuesaa/dotslide/internal"
)

var (
	port = *flag.IntP("port", "p", 3000, "Port")
	workdir = *flag.StringP("workdir", "w", ".", "Workdir")
)

func main() {
	flag.Parse()

	app := internal.New()
	app.Port = port
	app.Workdir = workdir

	if err := app.Run(); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}
