package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const WAREHOUSE_COLLECTION_NAME = "warehouses"

type Warehouse struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Ref           string             `bson:"ref"`
	CityRef       string             `bson:"city_ref"`
	TypeWarehouse string             `bson:"type_warehouse"`
	Description   string             `bson:"description"`
	SettlementRef string             `bson:"settlement_ref"`
	Longitude     string             `bson:"longitude"`
	Latitude      string             `bson:"latitude"`
	Active        bool               `bson:"active"`
	CreatedAt     string             `bson:"created_at,omitempty"`
	UpdatedAt     string             `bson:"updated_at,omitempty"`
}

func (h *RepoHandler) WarehouseCreateOrInsertViaRef(warehouse Warehouse) error {
	filter := bson.M{"ref": warehouse.Ref}
	update := bson.M{
		"$set": bson.M{
			"city_ref":       warehouse.CityRef,
			"type_warehouse": warehouse.TypeWarehouse,
			"description":    warehouse.Description,
			"settlement_ref": warehouse.SettlementRef,
			"longitude":      warehouse.Longitude,
			"latitude":       warehouse.Latitude,
			"active":         warehouse.Active,
			"updated_at":     time.Now().String(),
		},
		"$setOnInsert": bson.M{
			"created_at": time.Now().String(),
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err := h.getDatabase().Collection(WAREHOUSE_COLLECTION_NAME).UpdateOne(context.Background(), filter, update, opts)

	if err != nil {
		return err
	}
	return nil
}
