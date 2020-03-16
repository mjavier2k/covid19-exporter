package rapidapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/amoghe/distillog"
)

const (
	namespace  = "covid19"
	sourceUrl  = "https://covid-19-coronavirus-statistics.p.rapidapi.com/v1/stats"
	hostHeader = "covid-19-coronavirus-statistics.p.rapidapi.com"
)

func NewHTTPClient() *Client {
	log.Infoln("Initialize RapidAPI Client")

	insecure, err := strconv.ParseBool(os.Getenv("INSECURE_SKIP_VERIFY"))
	if err != nil {
		log.Warningln("Could not parse environment variable INSECURE_SKIP_VERIFY. Defaulting to INSECURE_SKIP_VERIFY=false")
		insecure = false
	}
	if insecure {
		log.Warningln("TLS certificate verification is currently disabled - This is not recommended.")
	}

	log.Infoln("RAPIDAPI_URL:", os.Getenv("RAPIDAPI_URL"))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	return &Client{
		HttpClient: &http.Client{
			Transport: tr,
			Timeout:   30 * time.Second,
		},
		HttpEndpoint: os.Getenv("RAPIDAPI_URL"),
		HostHeader:   os.Getenv("RAPIDAPI_HOST"),
		APIKey:       os.Getenv("RAPIDAPI_KEY"),
	}
}

func (c *Client) GetCovid19Data() (GISDataResponse, error) {

	req, error := http.NewRequest("GET", c.HttpEndpoint, nil)

	if error != nil {
		log.Errorln(error)
	}

	req.Header.Add("x-rapidapi-host", c.HostHeader)
	req.Header.Add("x-rapidapi-key", c.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if error != nil {
		log.Errorln(error)
	}

	defer resp.Body.Close()

	r := GISDataResponse{}
	if resp.StatusCode != 200 {
		return r, fmt.Errorf("Received a non-200 response code from remote server: %v", error)
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
