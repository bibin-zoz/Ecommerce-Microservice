package main

import (
	"github.com/bibin-zoz/ecommerce-api-gateway/pkg/config"
	"github.com/bibin-zoz/ecommerce-api-gateway/pkg/di"
	"log"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}

}