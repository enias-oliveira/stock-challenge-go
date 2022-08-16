package main

import (
	"log"
	config "stock-challenge-go/pkg/config"
	depinject "stock-challenge-go/pkg/depinject"
)

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
