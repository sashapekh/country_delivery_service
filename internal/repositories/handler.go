package repositories

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type RepoHandler struct {
	Client *mongo.Client
}

func NewRepoHandler(client *mongo.Client) *RepoHandler {
	return &RepoHandler{
		Client: client,
	}
}

func (h *RepoHandler) getDatabase() *mongo.Database {
	return h.Client.Database(os.Getenv("DATABASE_NAME"))
}
