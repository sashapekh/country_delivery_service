package novaposhta

import (
	"encoding/json"
	"os"
	apiparser "sashapekh/country_delivery_service/pkg/api-parser"
	"sashapekh/country_delivery_service/pkg/helpers"
)

const Endpoint = "https://api.novaposhta.ua/v2.0/json/"

type Novaposhta struct {
	Endpoint string
	client   *apiparser.Client
}

func NewNovaPoshta() *Novaposhta {
	return &Novaposhta{
		Endpoint: Endpoint,
		client:   apiparser.NewClient(),
	}
}

func (np *Novaposhta) GetAllRegions() ([]Region, error) {
	modelName := "AddressGeneral"
	calledMethod := "getSettlementAreas"

	requestBody := map[string]interface{}{
		"apiKey":       os.Getenv("NOVAPOSHTA_API_KEY"),
		"modelName":    modelName,
		"calledMethod": calledMethod,
	}

	req := apiparser.Request{
		Url:    Endpoint,
		Method: "POST",
		Body:   requestBody,
	}

	responseBody, err := np.client.MakeRequest(req)

	if err != nil {
		panic(err)
	}

	dataBytes, err := helpers.ExtractRawJSONField(responseBody, "data")

	if err != nil {
		return nil, err
	}

	var regions []Region

	err = json.Unmarshal(dataBytes, &regions)

	if err != nil {
		return nil, err
	}

	return regions, nil

}

func (np *Novaposhta) GetSettlments(areaRef string) ([]Settlement, error) {
	modelName := "AddressGeneral"
	calledMethod := "getSettlementCountryRegion"

	requestBody := map[string]interface{}{
		"apiKey":       os.Getenv("NOVAPOSHTA_API_KEY"),
		"modelName":    modelName,
		"calledMethod": calledMethod,
		"methodProperties": map[string]string{
			"AreaRef": areaRef,
		},
	}

	req := apiparser.Request{
		Url:    Endpoint,
		Method: "POST",
		Body:   requestBody,
	}

	body, err := np.client.MakeRequest(req)

	if err != nil {
		panic(err)
	}

	dataBytes, err := helpers.ExtractRawJSONField(body, "data")

	if err != nil {
		return nil, err
	}

	var settlements []Settlement

	err = json.Unmarshal(dataBytes, &settlements)

	if err != nil {
		return nil, err
	}

	return settlements, nil
}
