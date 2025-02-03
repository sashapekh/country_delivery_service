package sync

import (
	"context"
	"sashapekh/country_delivery_service/internal/repositories"
	"sync"
)

const (
	max_settlement_job       = 10
	max_warehouses_inert_job = 100
)

func (h *SyncServiceHanlder) SyncWarehouses() error {
	var wg sync.WaitGroup
	jobChan := make(chan struct{}, max_settlement_job)

	settlmentCursor, err := h.RepoHandler.GetSettlementCursor()

	if err != nil {
		return err
	}

	defer settlmentCursor.Close(context.Background())

	for settlmentCursor.Next(context.Background()) {
		var settlement repositories.Settlement

		err := settlmentCursor.Decode(&settlement)
		if err != nil {
			continue
		}

		wg.Add(1)
		jobChan <- struct{}{}
		go func(ref string) {
			defer wg.Done()
			defer func() { <-jobChan }()
			h.getWarehousesByRef(settlement.Ref)
		}(settlement.Ref)

	}
	wg.Wait()
	return nil
}

func (h *SyncServiceHanlder) getWarehousesByRef(cityRef string) error {
	var wg sync.WaitGroup
	jobChan := make(chan struct{}, max_warehouses_inert_job)

	warehouses, err := h.novaposhta.GetWarehousesByCityRef(cityRef)
	if err != nil {
		return err
	}

	for _, warehouse := range warehouses {

		repoWarehouse := repositories.Warehouse{
			Ref:           warehouse.Ref,
			CityRef:       warehouse.CityRef,
			Description:   warehouse.Description,
			TypeWarehouse: warehouse.TypeOfWarehouse,
			SettlementRef: warehouse.SettlementRef,
			Longitude:     warehouse.Longitude,
			Latitude:      warehouse.Latitude,
			Active:        true,
		}
		wg.Add(1)
		jobChan <- struct{}{}
		go func(repoWarehouse repositories.Warehouse) {
			defer wg.Done()
			defer func() { <-jobChan }()
			err := h.RepoHandler.WarehouseCreateOrInsertViaRef(repoWarehouse)
			if err != nil {
				return
			}
		}(repoWarehouse)
	}
	wg.Wait()
	return nil
}
