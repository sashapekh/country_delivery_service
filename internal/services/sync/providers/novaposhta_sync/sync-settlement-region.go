package novaposhta_sync

import (
	"fmt"
	"sashapekh/country_delivery_service/internal/repositories"
	"sashapekh/country_delivery_service/pkg/novaposhta"
	"sync"
	"time"

	"github.com/gosimple/slug"
)

func (h *SyncServiceHanlder) SyncSettlmentRegions() error {
	var wg sync.WaitGroup
	regions, err := h.RepoHandler.GetAllRegions()

	if err != nil {
		return err
	}

	for _, region := range regions {
		settlements, err := h.novaposhta.GetSettlementCountryRegions(region.NpRef)
		if err != nil {
			continue
		}

		for _, settlement := range settlements {
			go func(
				settlement novaposhta.SettlementCountryRegion,
				region repositories.Region,
			) {
				err = h.processSettlement(settlement, region.NpRef)
				if err != nil {
					h.logger.Error(err.Error(), "region", region.NpRef)
				}
			}(
				settlement,
				region,
			)

		}
	}

	wg.Wait()
	return nil
}

func (h *SyncServiceHanlder) processSettlement(settlment novaposhta.SettlementCountryRegion, regionRef string) error {
	fmt.Printf("Processing settlement: %s\n", settlment.Description)
	region, err := h.RepoHandler.GetRegionByRef(regionRef)
	if err != nil {
		return err
	}
	newSettlement := repositories.SettlementCountryRegion{
		Name:       settlment.Description,
		Slug:       slug.Make(settlment.Description),
		RegionID:   region.Id,
		RegionType: settlment.RegionType,
		NpRef:      settlment.Ref,
		Active:     true,
		CreatedAt:  time.Now().String(),
		UpdatedAt:  time.Now().String(),
	}

	err = h.RepoHandler.InsertOrUpdateSettlementCountryRegion(newSettlement)

	if err != nil {
		return err
	}

	return nil
}
