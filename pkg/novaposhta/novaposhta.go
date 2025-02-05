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
		return nil, err
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

func (np *Novaposhta) GetSettlementCountryRegions(areaRef string) ([]SettlementCountryRegion, error) {

	body, err := np.makeRequest(
		"AddressGeneral",
		"getSettlementCountryRegion",
		map[string]string{
			"AreaRef": areaRef,
		},
	)

	if err != nil {
		return nil, err
	}

	dataBytes, err := helpers.ExtractRawJSONField(body, "data")

	if err != nil {
		return nil, err
	}

	var settlements []SettlementCountryRegion

	err = json.Unmarshal(dataBytes, &settlements)

	if err != nil {
		return nil, err
	}

	return settlements, nil
}

func (np *Novaposhta) GetSettlements(page string, limit string) (SettlementResultByPage, error) {
	body, err := np.makeRequest(
		"AddressGeneral",
		"getSettlements",

		map[string]string{
			"Page":  page,
			"Limit": limit,
		},
	)

	if err != nil {
		return SettlementResultByPage{}, err
	}

	var settmentsResponse SettlementsResponse

	err = json.Unmarshal(body, &settmentsResponse)

	if err != nil {
		return SettlementResultByPage{}, err
	}
	var settlements []Settlement
	bodyExtract, err := helpers.ExtractRawJSONField(body, "data")

	if err != nil {
		return SettlementResultByPage{}, err
	}

	err = json.Unmarshal(bodyExtract, &settlements)

	if err != nil {
		return SettlementResultByPage{}, err
	}
	return SettlementResultByPage{
		Page:       page,
		TotalCount: settmentsResponse.Info.TotalCount,
		Items:      settlements,
	}, nil
}

func (np *Novaposhta) GetWarehousesByCityRef(ref string) ([]Warehouse, error) {

	request, err := np.makeRequest(
		"Address",
		"getWarehouses",
		map[string]string{
			"SettlementRef": ref,
		})

	if err != nil {
		return nil, err
	}

	var warehouses []Warehouse
	bodyExtract, err := helpers.ExtractRawJSONField(request, "data")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyExtract, &warehouses)

	if err != nil {
		return nil, err
	}

	return warehouses, nil

}
func (np *Novaposhta) makeRequest(modelName string, calledMethod string, properties map[string]string) ([]byte, error) {
	requestBody := map[string]interface{}{
		"apiKey":           os.Getenv("NOVAPOSHTA_API_KEY"),
		"modelName":        modelName,
		"calledMethod":     calledMethod,
		"methodProperties": properties,
	}

	req := apiparser.Request{
		Url:    Endpoint,
		Method: "POST",
		Body:   requestBody,
	}

	body, err := np.client.MakeRequest(req)

	if err != nil {
		return nil, err
	}

	return body, nil
}
