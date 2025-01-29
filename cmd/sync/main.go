package main

import (
	"sashapekh/country_delivery_service/internal/repositories"
	"sashapekh/country_delivery_service/internal/services/sync"
	"sashapekh/country_delivery_service/pkg/helpers"
	"sashapekh/country_delivery_service/pkg/novaposhta"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mongoClient, err := helpers.GetMongoClient()

	if err != nil {
		panic(err)
	}

	np := novaposhta.NewNovaPoshta()
	repo := repositories.NewRepoHandler(mongoClient)

	syncService := sync.NewSyncServiceHanlder(repo, np)

	// syncService.SyncRegions()
	// syncService.SyncSettlmentRegions()
	// syncService.SyncSettlments()
	syncService.SyncWarehouses()
}
