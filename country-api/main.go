package main

import (
	"log"

	"it.schwarz/country/app"

	"it.schwarz/country/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := config.NewConfig()
	app.ConfigAndRunApp(config)
}
