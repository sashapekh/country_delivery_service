package helpers

import "encoding/json"

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
