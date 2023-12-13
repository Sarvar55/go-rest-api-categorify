package model

type Location struct {
	IpTo        int64  `json:"ip_to"`
	IpFrom      int64  `json:"ip_from"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}

type DomainCategory struct {
	CategoryName string   `bson:"name"`
	Domains      []string `bson:"domains"`
	Urls         []string `bson:"urls"`
}
