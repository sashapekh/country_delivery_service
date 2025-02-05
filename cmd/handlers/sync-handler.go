package handlers

import "sashapekh/country_delivery_service/internal/repositories"

type SyncHandler struct {
	RepoHandler *repositories.RepoHandler
}
