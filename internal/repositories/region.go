package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const REGION_COLLECTION_NAME = "regions"

type Region struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Slug      string             `bson:"slug"`
	NpRef     string             `bson:"np_ref"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}

func (h *RepoHandler) InsertOrCreate(region Region) error {

	filter := bson.M{"np_ref": region.NpRef}
	update := bson.M{
		"$set": bson.M{
			"name":       region.Name,
			"slug":       region.Slug,
			"np_ref":     region.NpRef,
			"updated_at": region.UpdatedAt,
		},
		"$setOnInsert": bson.M{
			"created_at": region.CreatedAt,
		},
	}

	_, err := h.getDatabase().Collection(REGION_COLLECTION_NAME).
		UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))

	if err != nil {
		return err
	}

	return nil
}
func (h *RepoHandler) GetAllRegions() ([]Region, error) {
	var regions []Region

	collection := h.getDatabase().Collection(REGION_COLLECTION_NAME)
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var region Region
		err := cursor.Decode(&region)

		if err != nil {
			return nil, err
		}

		regions = append(regions, region)
	}

	return regions, nil
}
func (h *RepoHandler) GetRegionByRef(ref string) (Region, error) {
	filter := bson.M{"np_ref": ref}
	var region Region
	collection := h.getDatabase().Collection(REGION_COLLECTION_NAME)
	err := collection.FindOne(context.Background(), filter).Decode(&region)

	if err != nil {
		return region, err
	}
	return region, nil
}
