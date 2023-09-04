package config

import (
	"fmt"
	"log"
	"os"
)

const (
	pg_container = "postgres"
	pg_network   = "bridge"
	pg_database  = "loyalmed"
	pg_user      = "loyalmed"
	pg_pass      = "1q2w3e"
	pg_port      = "5432"
)

type ConfigWHI struct {
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

func New() *ConfigWHI {
	return &ConfigWHI{
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
			DbName:     getEnvOrDefault("PG_DB_NAME", pg_database),
			DbUser:     getEnvOrDefault("PG_USER", pg_user),
			DbPassword: getEnvOrDefault("PG_PASSWORD", pg_pass),
			DbPort:     getEnvOrDefault("PG_PORT", pg_port),
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
