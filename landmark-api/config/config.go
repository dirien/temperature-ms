package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//  Config the EnvironmentConfig struct
type Config struct {
	MongoDB     MongoDB
	SnykAPIKey  string
	SnykGroupID string
	SnykAPIUrl  string
	AppPort     string
}

//  MongoDB struct
type MongoDB struct {
	URL              string
	Database         string
	SnykOrganization string
	SnykMember       string
}

// initialize config structure fields from env vars
func (config *Config) initialize() {
	initDevMode()
	config.AppPort = os.Getenv("APP_PORT")
}

func initDevMode() {
	env := os.Getenv("PLATFORM")
	fmt.Println(fmt.Sprintf("env '%s'", env))
	if "" == env {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

// NewConfig create and initialize config
func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}
