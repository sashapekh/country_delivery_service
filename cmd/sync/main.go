package main

import (
	"sashapekh/country_delivery_service/internal/repositories"
	"sashapekh/country_delivery_service/internal/services/sync"
	"sashapekh/country_delivery_service/pkg/helpers"
	"sashapekh/country_delivery_service/pkg/logger"
	"sashapekh/country_delivery_service/pkg/novaposhta"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	loggerHandler := logger.NewLoggerHandler(logger.Default_file_channel)
	defer loggerHandler.Close()
	logger := loggerHandler.GetLogger()

	mongoClient, err := helpers.GetMongoClient()

	if err != nil {
		panic(err)
	}

	np := novaposhta.NewNovaPoshta()
	repo := repositories.NewRepoHandler(mongoClient)

	syncService := sync.NewSyncServiceHanlder(repo, np, logger)

	// syncService.SyncRegions()
	// syncService.SyncSettlmentRegions()
	// syncService.SyncSettlments()
	syncService.SyncWarehouses()
}
