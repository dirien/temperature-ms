package main

import (
	"log"

	"it.schwarz/city/app"

	"it.schwarz/city/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := config.NewConfig()
	app.ConfigAndRunApp(config)
}
