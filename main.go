package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/dotslide/internal"
	flag "github.com/spf13/pflag"
)

var (
	port    = flag.IntP("port", "p", 3000, "Port")
	workdir = flag.StringP("workdir", "w", ".", "Workdir")
	capture = flag.Bool("capture", false, "Make Capture")
	help    = flag.BoolP("help", "h", false, "Print help message")
	version = flag.BoolP("version", "v", false, "Print app version")
)

func main() {
	flag.Parse()

	if *help {
		fmt.Println("Usage of dotslide:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *version {
		fmt.Println("dotslide v0.0.1")
		os.Exit(0)
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
