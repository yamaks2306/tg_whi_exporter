package config

import (
	"fmt"
	"log"
	"os"
)

const (
	pg_container = "postgres"
	pg_network   = "bridge"
)

type Config struct {
	TgConfig       ConfigTg
	PgDockerConfig ConfigPgDocker
	PgConfig       ConfigPg
}

type ConfigPgDocker struct {
	PgContainerName    string
	PgContainerNetwork string
}

type ConfigTg struct {
	TelegramToken string
	ServerURL     string
	MisType       string
}

type ConfigPg struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbPort     string
}

func New() *Config {
	return &Config{
		TgConfig: ConfigTg{
			TelegramToken: getEnv("TELEGRAM_TOKEN"),
			ServerURL:     getEnv("SERVER_URL"),
			MisType:       getEnv("MIS_TYPE"),
		},
		PgDockerConfig: ConfigPgDocker{
			PgContainerName:    getEnvOrDefault("PG_CONTAINER_NAME", pg_container),
			PgContainerNetwork: getEnvOrDefault("PG_CONTAINER_NETWORK", pg_network),
		},
		PgConfig: ConfigPg{
			DbName:     getEnv("POSTGRES_DB"),
			DbUser:     getEnv("POSTGRES_USER"),
			DbPassword: getEnv("POSTGRES_PASSWORD"),
			DbPort:     getEnv("POSTGRES_PORT"),
		},
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

func getEnvOrDefault(key, def string) string {
	value, exists := os.LookupEnv(key)

	if !exists || value == "" {
		return def
	}

	return value
}
