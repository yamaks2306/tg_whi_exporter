package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yamaks2306/tg_whi_exporter/collector"
	"github.com/yamaks2306/tg_whi_exporter/config"
)

const (
	ENV_PROD   = "env/env.prod"
	ENV_COMMON = "env/env.common"
)

func init() {
	if err := godotenv.Load(ENV_PROD); err != nil {
		log.Fatal("No prod env file found")
	}
	if err := godotenv.Load(ENV_COMMON); err != nil {
		log.Fatal("No common env file found")
	}
}

func main() {

	config := config.New()

	tg_collector := collector.NewTgInfoCollector(*config)
	tg_registry := prometheus.NewRegistry()
	tg_registry.MustRegister(tg_collector)
	tg_prom_handler := promhttp.HandlerFor(tg_registry, promhttp.HandlerOpts{})

	pg_collector := collector.NewPgCollector(*config)
	pg_registry := prometheus.NewRegistry()
	pg_registry.MustRegister(pg_collector)
	pg_prom_handler := promhttp.HandlerFor(pg_registry, promhttp.HandlerOpts{})

	http.Handle("/tg-metrics", tg_prom_handler)
	http.Handle("/pg-metrics", pg_prom_handler)
	log.Println("Start metrics web service")
	http.ListenAndServe(":9900", nil)

}
