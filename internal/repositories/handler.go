package repositories

import (
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
