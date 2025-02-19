package sync

import (
	"sashapekh/country_delivery_service/internal/services/sync/providers"
)

func (h *SyncServiceHanlder) SyncRegions() error {

	err := providers.SyncRegionsNovaposhta(
		h.novaposhta,
		h.RepoHandler,
	)

	if err != nil {
		h.logger.Error(err.Error())
		return err
	}

	return nil
}
