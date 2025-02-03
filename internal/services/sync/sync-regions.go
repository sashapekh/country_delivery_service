package sync

import (
	"fmt"
	"sashapekh/country_delivery_service/internal/repositories"
	"time"

	"github.com/gosimple/slug"
)

func (h *SyncServiceHanlder) SyncRegions() error {
	h.logger.Info("SyncRegions() started")

	regions, err := h.novaposhta.GetAllRegions()

	if err != nil {
		h.logger.Error("Error while getting regions from NovaPoshta", "error", err)
		return err
	}

	for _, region := range regions {

		mongoRegion := repositories.Region{
			Name:      region.Description,
			Slug:      slug.Make(region.Description),
			NpRef:     region.Ref,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		}

		h.logger.Info(fmt.Sprintf("Inserting region %s", region.Description))

		err := h.RepoHandler.InsertRegion(mongoRegion)

		if err != nil {
			return err
		}
	}
	return nil
}
