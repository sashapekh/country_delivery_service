package novaposhta_sync

import (
	"fmt"
	"math"
	"sashapekh/country_delivery_service/internal/repositories"
	"sashapekh/country_delivery_service/pkg/novaposhta"
	"sync"
	"time"
)

const (
	page_size     = 100
	first_page    = "1"
	maxGoroutines = 10
)

var totalCount int

func (h *SyncServiceHanlder) SyncSettlments() error {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxGoroutines)

	responseResult, err := h.novaposhta.GetSettlements(first_page, fmt.Sprintf("%d", page_size))

	if err != nil {
		return err
	}
	totalCount = responseResult.TotalCount
	totalPages := int(math.Ceil(float64(totalCount) / float64(page_size)))

	fmt.Printf("Total pages: %d\n", totalPages)

	// Process the first page
	h.saveSettlements(responseResult.Items)

	for page := 2; page <= totalPages; page++ {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(page int) {
			defer wg.Done()
			defer func() { <-semaphore }()

			time.Sleep(1 * time.Second)
			fmt.Printf("Processing page: %d\n", page)
			responseResult, err := h.novaposhta.GetSettlements(fmt.Sprintf("%d", page), fmt.Sprintf("%d", page_size))
			if err != nil {
				fmt.Printf("Error fetching page %d: %s\n", page, err)
				return
			}
			h.saveSettlements(responseResult.Items)
		}(page)

	}
	wg.Wait()
	return nil
}

func (h *SyncServiceHanlder) saveSettlements(settlements []novaposhta.Settlement) error {
	for _, settlement := range settlements {
		newSettlement := repositories.Settlement{
			Ref:                               settlement.Ref,
			Name:                              settlement.Description,
			Area:                              settlement.Area,
			AreaDescription:                   settlement.AreaDescription,
			Region:                            settlement.Region,
			RegionsDescription:                settlement.RegionsDescription,
			SettlementType:                    settlement.SettlementType,
			SettlementTypeDescription:         settlement.SettlementTypeDescription,
			SettlementTypeDescriptionTranslit: settlement.SettlementTypeDescriptionTranslit,
			IndexCOATSU1:                      settlement.IndexCOATSU1,
			Latitude:                          settlement.Latitude,
			Longitude:                         settlement.Longitude,
			CreatedAt:                         time.Now().String(),
		}

		h.RepoHandler.SettlmentCreateOrInsertViaRef(newSettlement)
	}

	return nil
}
