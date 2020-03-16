package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mjavier2k/covid19-exporter/pkg/rapidapi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/amoghe/distillog"
)

type rapidAPICollector struct {
	client *rapidapi.Client
}

type Descriptions struct {
	ScrapeSuccessDesc *prometheus.Desc

	ConfirmedCount *prometheus.Desc
	RecoveredCount *prometheus.Desc
	DeathsCount    *prometheus.Desc
}

func NewMetricDescriptions(namespace string) *Descriptions {
	var d Descriptions

	d.ScrapeSuccessDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "scrape_success"),
		"Whether last scrape was successful",
		nil,
		nil,
	)

	d.ConfirmedCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "infected"),
		"Number of infection per country",
		[]string{"country", "province"},
		nil,
	)

	d.RecoveredCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "recovered"),
		"Number of infection per country",
		[]string{"country", "province"},
		nil,
	)

	d.DeathsCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "deaths"),
		"Number of deaths per country",
		[]string{"country", "province"},
		nil,
	)

	return &d
}

func NewCollector() (*rapidAPICollector, error) {
	log.Infof("Initializing New RapidAPI Collector")

	return &rapidAPICollector{
		client: rapidapi.NewHTTPClient(),
	}, nil
}

func (c *rapidAPICollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- MetricDescriptions.ConfirmedCount
	ch <- MetricDescriptions.RecoveredCount
	ch <- MetricDescriptions.DeathsCount
}

func (c *rapidAPICollector) Collect(ch chan<- prometheus.Metric) {
	var scrapeSuccess float64 = 1

	result, err := c.client.GetCovid19Data()

	if err != nil {
		log.Errorln(err)
	}

	for _, country := range result.Data.Covid19Stats {
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ConfirmedCount,
			prometheus.GaugeValue,
			country.Confirmed,
			country.Country,
			country.Province,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.RecoveredCount,
			prometheus.GaugeValue,
			country.Recovered,
			country.Country,
			country.Province,
		)

		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.DeathsCount,
			prometheus.GaugeValue,
			country.Recovered,
			country.Country,
			country.Province,
		)
	}

	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.ScrapeSuccessDesc,
		prometheus.GaugeValue,
		scrapeSuccess,
	)
}

var (
	MetricDescriptions = NewMetricDescriptions("covid19")
)

func resolvePort() string {
	port, ok := os.LookupEnv("EXPORTER_PORT")
	if !ok {
		port = "9999"
	}

	return port
}

func main() {
	listenAddr := fmt.Sprintf("0.0.0.0:%v", resolvePort())

	rapidAPIExporter, _ := NewCollector()
	prometheus.MustRegister(rapidAPIExporter)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "UP")
	})

	log.Infof("Booted and listening on %v/metrics\n", listenAddr)
	err := http.ListenAndServe(listenAddr, nil)

	if err != nil {
		log.Errorln(err)
	}
}
