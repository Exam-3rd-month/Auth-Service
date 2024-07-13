package main

import (
	"log"

	"auth-service/api"
	"auth-service/internal/config"
	"auth-service/internal/service"
	"auth-service/internal/storage"
)

func main() {
	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	storage, err := storage.New(configs)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(service.New(*storage))

	log.Fatal(api.RUN(configs))
}
