package novaposhta

type SettlementsResponse struct {
	Success  bool         `json:"success"`
	Data     []Settlement `json:"data"`
	Errors   []string     `json:"errors"`
	Warnings []string     `json:"warnings"`
	Info     struct {
		TotalCount int `json:"totalCount"`
	} `json:"info"`
}
