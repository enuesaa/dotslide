package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/dotslide/internal"
	flag "github.com/spf13/pflag"
)

var (
	portFlag    = flag.IntP("port", "p", 3000, "Port")
	workdirFlag = flag.StringP("workdir", "w", ".", "Workdir")
	captureFlag = flag.Bool("capture", false, "Make Capture")
	helpFlag    = flag.BoolP("help", "h", false, "Print help message")
	versionFlag = flag.BoolP("version", "v", false, "Print app version")
)

func main() {
	flag.Parse()

	if *helpFlag {
		fmt.Println("Usage of dotslide:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *versionFlag {
		fmt.Println("dotslide v0.0.1")
		os.Exit(0)
	}

	config := internal.Config{
		Port:    *portFlag,
		Workdir: *workdirFlag,
		Capture: *captureFlag,
	}
	app := internal.New(config)

	if err := app.Run(); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}
