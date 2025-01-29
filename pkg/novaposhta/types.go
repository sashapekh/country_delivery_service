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

type Settlement struct {
	Ref                               string `json:"Ref"`
	SettlementType                    string `json:"SettlementType"`
	Latitude                          string `json:"Latitude"`
	Longitude                         string `json:"Longitude"`
	Description                       string `json:"Description"`
	DescriptionRu                     string `json:"DescriptionRu"`
	DescriptionTranslit               string `json:"DescriptionTranslit"`
	SettlementTypeDescription         string `json:"SettlementTypeDescription"`
	SettlementTypeDescriptionRu       string `json:"SettlementTypeDescriptionRu"`
	SettlementTypeDescriptionTranslit string `json:"SettlementTypeDescriptionTranslit"`
	Region                            string `json:"Region"`
	RegionsDescription                string `json:"RegionsDescription"`
	RegionsDescriptionRu              string `json:"RegionsDescriptionRu"`
	RegionsDescriptionTranslit        string `json:"RegionsDescriptionTranslit"`
	Area                              string `json:"Area"`
	AreaDescription                   string `json:"AreaDescription"`
	AreaDescriptionRu                 string `json:"AreaDescriptionRu"`
	AreaDescriptionTranslit           string `json:"AreaDescriptionTranslit"`
	Index1                            string `json:"Index1"`
	Index2                            string `json:"Index2"`
	IndexCOATSU1                      string `json:"IndexCOATSU1"`
	Delivery1                         string `json:"Delivery1"`
	Delivery2                         string `json:"Delivery2"`
	Delivery3                         string `json:"Delivery3"`
	Delivery4                         string `json:"Delivery4"`
	Delivery5                         string `json:"Delivery5"`
	Delivery6                         string `json:"Delivery6"`
	Delivery7                         string `json:"Delivery7"`
	SpecialCashCheck                  int    `json:"SpecialCashCheck"`
	RadiusHomeDelivery                string `json:"RadiusHomeDelivery"`
	RadiusExpressPickUp               string `json:"RadiusExpressPickUp"`
	RadiusDrop                        string `json:"RadiusDrop"`
	Warehouse                         string `json:"Warehouse"`
}

type SettlementResultByPage struct {
	Page       string `json:"page"`
	TotalCount int    `json:"totalCount"`
	Items      []Settlement
}

type Warehouse struct {
	SiteKey                      string `json:"SiteKey"`
	Description                  string `json:"Description"`
	DescriptionRu                string `json:"DescriptionRu"`
	ShortAddress                 string `json:"ShortAddress"`
	ShortAddressRu               string `json:"ShortAddressRu"`
	Phone                        string `json:"Phone"`
	TypeOfWarehouse              string `json:"TypeOfWarehouse"`
	Ref                          string `json:"Ref"`
	Number                       string `json:"Number"`
	CityRef                      string `json:"CityRef"`
	CityDescription              string `json:"CityDescription"`
	CityDescriptionRu            string `json:"CityDescriptionRu"`
	SettlementRef                string `json:"SettlementRef"`
	SettlementDescription        string `json:"SettlementDescription"`
	SettlementAreaDescription    string `json:"SettlementAreaDescription"`
	SettlementRegionsDescription string `json:"SettlementRegionsDescription"`
	SettlementTypeDescription    string `json:"SettlementTypeDescription"`
	SettlementTypeDescriptionRu  string `json:"SettlementTypeDescriptionRu"`
	Longitude                    string `json:"Longitude"`
	Latitude                     string `json:"Latitude"`
	PostFinance                  string `json:"PostFinance"`
	BicycleParking               string `json:"BicycleParking"`
	PaymentAccess                string `json:"PaymentAccess"`
	POSTerminal                  string `json:"POSTerminal"`
	InternationalShipping        string `json:"InternationalShipping"`
	SelfServiceWorkplacesCount   string `json:"SelfServiceWorkplacesCount"`
	TotalMaxWeightAllowed        string `json:"TotalMaxWeightAllowed"`
	PlaceMaxWeightAllowed        string `json:"PlaceMaxWeightAllowed"`
}
