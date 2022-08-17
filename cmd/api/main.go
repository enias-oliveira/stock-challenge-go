package main

import (
	"log"
	config "stock-challenge-go/pkg/config"
	depinject "stock-challenge-go/pkg/depinject"
)

// @title           Swagger Example API
// @version         1.0
// @description     The GO implementation of the Stooq Stock API challenge.

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	config, cErr := config.LoadConfig()

	if cErr != nil {
		log.Fatal("Error loading config: ", cErr)
	}

	server, depInjErr := depinject.InitializeAPI(config)

	if depInjErr != nil {
		log.Fatal("Error initializing API: ", depInjErr)
	}

	server.Run(config.APIHost + ":" + config.APIPort)
}
