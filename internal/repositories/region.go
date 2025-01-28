package repositories

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

const REGION_COLLECTION_NAME = "regions"

type Region struct {
	Name      string `bson:"name"`
	Slug      string `bson:"slug"`
	NpRef     string `bson:"np_ref"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}

func (h *RepoHandler) InsertRegion(region Region) error {
	collection := h.Client.Database(os.Getenv("DATABASE_NAME")).Collection("regions")
	_, err := collection.InsertOne(context.Background(), region)
	if err != nil {
		return err
	}
	return nil
}

func (h *RepoHandler) GetRegionByRef(ref string) (Region, error) {
	filter := bson.M{"np_ref": ref}
	var region Region
	collection := h.Client.Database(os.Getenv("DATABASE_NAME")).Collection("regions")
	err := collection.FindOne(context.Background(), filter).Decode(&region)

	if err != nil {
		return region, err
	}
	return region, nil
}
