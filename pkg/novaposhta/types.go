package novaposhta

type Region struct {
	Ref         string `json:"Ref"`
	Description string `json:"Description"`
	RegionType  string `json:"RegionType"`
	AreasCenter string `json:"AreasCenter"`
}

type SettlementCountryRegion struct {
	Ref         string `json:"Ref"`
	Description string `json:"Description"`
	RegionType  string `json:"RegionType"`
	AreasCenter string `json:"AreasCenter"`
}
