package rapidapi

import (
	"net/http"
	"time"
)

type Client struct {
	HostHeader   string
	APIKey       string
	HttpEndpoint string
	HttpClient   *http.Client
}

type GISDataResponse struct {
	Error      bool   `json:"error"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		LastChecked  time.Time `json:"lastChecked"`
		Covid19Stats []struct {
			Province   string  `json:"province"`
			Country    string  `json:"country"`
			LastUpdate string  `json:"lastUpdate"`
			Confirmed  float64 `json:"confirmed"`
			Deaths     float64 `json:"deaths"`
			Recovered  float64 `json:"recovered"`
		} `json:"covid19Stats"`
	} `json:"data"`
}
