package config

import (
	"fmt"
	"log"
	"os"
)

type configWHI struct {
	TelegramToken string
	ServerURL     string
	MisType       string
}

func New() *configWHI {
	return &configWHI{
		TelegramToken: getEnv("TELEGRAM_TOKEN"),
		ServerURL:     getEnv("SERVER_URL"),
		MisType:       getEnv("MIS_TYPE"),
	}
}

func getEnv(key string) string {
	value, exits := os.LookupEnv(key)

	if !exits || value == "" {
		error_string := fmt.Sprintf("Environment variable %s does not exist or empty", key)
		log.Fatal(error_string)
	}

	return value
}