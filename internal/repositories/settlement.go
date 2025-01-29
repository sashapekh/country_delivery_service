package repositories

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const SETTLEMENT_COLLECTION_NAME = "settlements"

type Settlement struct {
	ID                                primitive.ObjectID `bson:"_id,omitempty"`
	AreaDescription                   string             `bson:"area_description"`
	Area                              string             `bson:"area"`
	Ref                               string             `bson:"ref"`
	SettlementType                    string             `bson:"settlement_type"`
	SettlementTypeDescription         string             `bson:"settlement_type_description"`
	SettlementTypeDescriptionTranslit string             `bson:"settlement_type_description_translit"`
	IndexCOATSU1                      string             `bson:"index_coats_u_1"`
	Name                              string             `bson:"name"`
	Region                            string             `bson:"region"`
	RegionsDescription                string             `bson:"regions_description"`
	Latitude                          string             `bson:"latitude"`
	Longitude                         string             `bson:"longitude"`
	CreatedAt                         string             `bson:"created_at,omitempty"`
	UpdatedAt                         string             `bson:"updated_at,omitempty"`
}

func (h *RepoHandler) SettlmentCreateOrInsertViaRef(settlement Settlement) error {
	filter := bson.M{"ref": settlement.Ref}
	update := bson.M{
		"$set": bson.M{
			"area_description":                     settlement.AreaDescription,
			"area":                                 settlement.Area,
			"settlement_type":                      settlement.SettlementType,
			"settlement_type_description":          settlement.SettlementTypeDescription,
			"settlement_type_description_translit": settlement.SettlementTypeDescriptionTranslit,
			"index_coats_u_1":                      settlement.IndexCOATSU1,
			"name":                                 settlement.Name,
			"region":                               settlement.Region,
			"regions_description":                  settlement.RegionsDescription,
			"latitude":                             settlement.Latitude,
			"longitude":                            settlement.Longitude,
			"updated_at":                           time.Now().String(),
		},
		"$setOnInsert": bson.M{
			"created_at": time.Now().String(),
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err := h.getDatabase().Collection(SETTLEMENT_COLLECTION_NAME).UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		fmt.Printf("Error upserting settlement %s: %s\n", settlement.Ref, err)
		return err
	}
	return nil
}

func (h *RepoHandler) GetSettlementCursor() (*mongo.Cursor, error) {
	cursor, err := h.getDatabase().Collection(SETTLEMENT_COLLECTION_NAME).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	return cursor, nil
}
