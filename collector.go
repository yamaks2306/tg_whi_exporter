package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type tgInfoCollector struct {
	pending_updates  *prometheus.Desc
	request_duration *prometheus.Desc
}

func newTgInfoCollector() *tgInfoCollector {
	return &tgInfoCollector{
		pending_updates: prometheus.NewDesc(
			"pending_telegram_webhook_updates_count",
			"Telegram webhook pending updates count",
			[]string{"client", "mis_type"}, nil,
		),
		request_duration: prometheus.NewDesc(
			"request_duration",
			"Request duration",
			[]string{"client", "mis_type"}, nil,
		),
	}
}

func (collector *tgInfoCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.pending_updates
	ch <- collector.request_duration
}

func (collector *tgInfoCollector) Collect(ch chan<- prometheus.Metric) {

	telegram_token, exits := os.LookupEnv("TELEGRAM_TOKEN")

	if !exits || telegram_token == "" {
		log.Println("Environment variable TELEGRAM_TOKEN does not exist or empty")
		os.Exit(1)
	}

	server_url, exists := os.LookupEnv("SERVER_URL")

	if !exists || server_url == "" {
		log.Println("Environment variable SERVER_URL does not exist or empty")
		os.Exit(1)
	}

	mis_type, exists := os.LookupEnv("MIS_TYPE")

	if !exists || mis_type == "" {
		log.Println("Environment variable MIS_TYPE does not exist or empty")
		os.Exit(1)
	}

	start := time.Now()
	webhookinfo := get_webhook_info(telegram_token).Result.Pending_update_count
	duration := time.Since(start).Seconds()

	ch <- prometheus.MustNewConstMetric(collector.pending_updates, prometheus.GaugeValue, float64(webhookinfo), server_url, mis_type)
	ch <- prometheus.MustNewConstMetric(collector.request_duration, prometheus.GaugeValue, duration, server_url, mis_type)

}

func get_webhook_info(token string) TgWebhookInfo {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getWebhookInfo", token)
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	log.Println(response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("No response from request")
	}
	log.Println(string(body))

	var result TgWebhookInfo
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

type TgWebhookInfo struct {
	Status bool `json:"ok"`
	Result struct {
		URL                    string `json:"url"`
		Has_custom_certificate bool   `json:"has_custom_certificate"`
		Pending_update_count   int    `json:"pending_update_count"`
		Max_connections        int    `json:"max_connections"`
		IP_address             string `json:"ip_address"`
	} `json:"result"`
}
