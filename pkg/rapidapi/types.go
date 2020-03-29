package rapidapi

import (
	"net/http"
)

type Client struct {
	HostHeader   string
	APIKey       string
	HttpEndpoint string
	HttpClient   *http.Client
}

type GISDataResponse []struct {
	FIPS          string  `json:"FIPS"`
	Admin2        string  `json:"Admin2"`
	ProvinceState string  `json:"Province_State"`
	CountryRegion string  `json:"Country_Region"`
	LastUpdate    string  `json:"Last_Update"`
	Lat           string  `json:"Lat"`
	Long          string  `json:"Long"`
	Confirmed     float64 `json:"Confirmed"`
	Deaths        float64 `json:"Deaths"`
	Recovered     float64 `json:"Recovered"`
	Active        float64 `json:"Active"`
	CombinedKey   string  `json:"Combined_Key"`
}
