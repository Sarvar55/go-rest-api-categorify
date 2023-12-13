package model

type LocationResponse struct {
	IpTo        int64  `json:"ip_to"`
	IpFrom      int64  `json:"ip_from"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}

type DomainCategoryResponse struct {
	CategoryName string `bson:"name"`
	Domain       string `json:"domain"`
	IpAddress    string `json:"ip_address"`
	CountryName  string `json:"country_name"`
}
