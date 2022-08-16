package main

import (
	"log"
	config "stock-challenge-go/pkg/config"
	depInj "stock-challenge-go/pkg/dependency_injection"
)

func main() {
	config, cErr := config.LoadConfig()

	if cErr != nil {
		log.Fatal("Error loading config: ", cErr)
	}

	server, depInjErr := depInj.InitializeAPI(config)

	if depInjErr != nil {
		log.Fatal("Error initializing API: ", depInjErr)
	}

	server.Run(":8080")
}
