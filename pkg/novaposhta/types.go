package novaposhta

type Region struct {
	Ref         string `json:"Ref"`
	Description string `json:"Description"`
	RegionType  string `json:"RegionType"`
	AreasCenter string `json:"AreasCenter"`
}

type Settlement struct {
	AreasCenter string `json:"AreasCenter"`
	Description string `json:"Description"`
	Ref         string `json:"Ref"`
	RegionType  string `json:"RegionType"`
}
