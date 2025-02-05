package helpers

import (
	"context"
	"encoding/json"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ExtractRawJSONField(body []byte, need string) ([]byte, error) {
	var topLevel map[string]json.RawMessage

	err := json.Unmarshal(body, &topLevel)

	if err != nil {
		return nil, err
	}

	dataBytes, find := topLevel[need]

	if !find {
		return nil, err
	}

	return dataBytes, nil
}

func MapToJSON(data map[string]string) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func GetMongoClient() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return nil, err
	}
	return client, nil
}
