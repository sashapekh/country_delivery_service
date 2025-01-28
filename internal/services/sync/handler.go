package sync

import (
	"sashapekh/country_delivery_service/internal/repositories"
	"sashapekh/country_delivery_service/pkg/novaposhta"
)

type SyncServiceHanlder struct {
	RepoHandler *repositories.RepoHandler
	novaposhta  *novaposhta.Novaposhta
}

func NewSyncServiceHanlder(repoHandler *repositories.RepoHandler, novaposhta *novaposhta.Novaposhta) *SyncServiceHanlder {
	return &SyncServiceHanlder{
		RepoHandler: repoHandler,
		novaposhta:  novaposhta,
	}
}
