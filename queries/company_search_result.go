package queries

type CompanySearchResult struct {
	CIK      string `json:"cik"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	Industry string `json:"industry"`
	Office   string `json:"office"`
}
