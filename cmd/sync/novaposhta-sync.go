package main

import (
	"sashapekh/country_delivery_service/internal/repositories"
	"sashapekh/country_delivery_service/internal/services/sync/providers/novaposhta_sync"
	"sashapekh/country_delivery_service/pkg/helpers"
	"sashapekh/country_delivery_service/pkg/logger"
	"sashapekh/country_delivery_service/pkg/novaposhta"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	loggerHandler := logger.NewLoggerHandler(logger.Default_file_channel)
	defer loggerHandler.Close()
	localLogger := loggerHandler.GetLogger()

	mongoClient, err := helpers.GetMongoClient()

	if err != nil {

		panic(err)
	}

	np := novaposhta.NewNovaPoshta()
	repo := repositories.NewRepoHandler(mongoClient)

	syncService := novaposhta_sync.New(repo, np, localLogger)

	err = syncService.SyncRegions()

	if err != nil {
		localLogger.Info("sync regions ≤failed")
	}

	err = syncService.SyncSettlmentRegions()

	if err != nil {
		localLogger.Info("sync settlment regions failed")
	}
	err = syncService.SyncSettlments()

	if err != nil {
		localLogger.Info("sync settlments failed")
	}
	err = syncService.SyncWarehouses()

	if err != nil {
		localLogger.Info("sync warehouses failed")
	}
}
