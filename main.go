package main

import (
	"log"

	"github.com/enuesaa/dotslide/internal"
	flag "github.com/spf13/pflag"
)

var (
	port    = flag.IntP("port", "p", 3000, "Port")
	workdir = flag.StringP("workdir", "w", ".", "Workdir")
	capture = flag.Bool("capture", false, "Make Capture")
)

func main() {
	flag.Parse()

	if *capture {
		if err := internal.Capture(); err != nil {
			log.Panicf("Error: %s", err.Error())
		}
	}

	config := internal.Config{
		Port:    *port,
		Workdir: *workdir,
		Capture: *capture,
	}
	app := internal.New(config)

	if err := app.Run(); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}
