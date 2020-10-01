package main

import (
	"log"

	"it.schwarz/landmark/app"

	"it.schwarz/landmark/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := config.NewConfig()
	app.ConfigAndRunApp(config)
}
