package sync

import (
	"log/slog"
	"sashapekh/country_delivery_service/internal/repositories"
	"sashapekh/country_delivery_service/pkg/novaposhta"
)

type SyncServiceHanlder struct {
	RepoHandler *repositories.RepoHandler
	novaposhta  *novaposhta.Novaposhta
	logger      *slog.Logger
}

func NewSyncServiceHanlder(
	repoHandler *repositories.RepoHandler,
	novaposhta *novaposhta.Novaposhta,
	slog *slog.Logger,
) *SyncServiceHanlder {
	return &SyncServiceHanlder{
		RepoHandler: repoHandler,
		novaposhta:  novaposhta,
		logger:      slog,
	}
}
