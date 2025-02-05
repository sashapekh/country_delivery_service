package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SETTLEMENT_COUNTRY_REGION_COLLECTION_NAME = "settlement_country_regions"

type SettlementCountryRegion struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Slug       string             `bson:"slug"`
	NpRef      string             `bson:"np_ref"`
	RegionType string             `bson:"region_type"`
	RegionID   primitive.ObjectID `bson:"region_id"`
	CreatedAt  string             `bson:"created_at"`
	UpdatedAt  string             `bson:"updated_at"`
	Active     bool               `bson:"active"`
}

func (h *RepoHandler) InsertSettlementCountryRegion(settlementCountryRegion SettlementCountryRegion) error {
	collection := h.getDatabase().Collection(SETTLEMENT_COUNTRY_REGION_COLLECTION_NAME)
	_, err := collection.InsertOne(context.Background(), settlementCountryRegion)
	if err != nil {
		return err
	}
	return nil
}

func (h *RepoHandler) GetAllSettlementCountryRegions() ([]SettlementCountryRegion, error) {
	var settlementCountryRegions []SettlementCountryRegion

	collection := h.getDatabase().Collection(SETTLEMENT_COUNTRY_REGION_COLLECTION_NAME)
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var settlementCountryRegion SettlementCountryRegion
		err := cursor.Decode(&settlementCountryRegion)
		if err != nil {
			return nil, err
		}
		settlementCountryRegions = append(settlementCountryRegions, settlementCountryRegion)
	}

	return settlementCountryRegions, nil
}
