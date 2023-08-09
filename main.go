package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	if err := godotenv.Load("env/env.prod"); err != nil {
		log.Println("No env file found")
	}
}

func main() {

	collector := newTgInfoCollector()
	registry := prometheus.NewRegistry()
	registry.MustRegister(collector)

	promHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	http.Handle("/metrics", promHandler)
	log.Println("Start metrics web service")
	http.ListenAndServe(":8080", nil)

}
