package main

import (
	"log"

	"github.com/enuesaa/dotslide/internal"
	flag "github.com/spf13/pflag"
)

var (
	port = flag.IntP("port", "p", 3000, "Port")
	workdir = flag.StringP("workdir", "w", ".", "Workdir")
)

func main() {
	flag.Parse()

	config := internal.Config{
		Port: *port,
		Workdir: *workdir,
	}
	app := internal.New(config)

	if err := app.Run(); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}
