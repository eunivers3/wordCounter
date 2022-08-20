package models

type Request struct {
	Address string `json:"address"`
	//Addresses          []string `json:"addresses"`
	//CountryRestriction string   `json:"country_restriction"`
	//LanguageOutput     string   `json:"language_output"`
}

type Parameters struct {
	Address string
	APIKey string
}

type Response struct {
	Results []Result `json:"results"`
	Status  string   `json:"status"`
}

type Result struct {
	ID               string   `json:"place_id"`
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}
