package config

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	DatabaseEnv   string = "DATABASE_URL"
	DataSourceEnv string = "DATA_SOURCE"
	ServerPortEnv string = "SERVER_PORT"
)

const (
	MemorySource   string = "memory"
	PostgresSource string = "postgres"
)

type Config struct {
	DatabaseUrl    string
	DataSourceType string
	ServerPort     string
}

func DefaultConfig() Config {
	return Config{
		DatabaseUrl:    "",
		DataSourceType: MemorySource,
		ServerPort:     "8080",
	}
}

func validDataSource(value string) bool {
	return value == MemorySource || value == PostgresSource
}

func BuildConfig() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}, MissingEnvFile{}
	}

	config := DefaultConfig()
	if value, exists := os.LookupEnv(DataSourceEnv); exists {
		if !validDataSource(value) {
			return Config{}, IncorrectEnvVarValue{
				EnvVar:  DataSourceEnv,
				Options: []string{MemorySource, PostgresSource},
			}
		}
		config.DataSourceType = value
	}

	if value, exists := os.LookupEnv(ServerPortEnv); exists {
		config.ServerPort = value
	}

	if config.DataSourceType == PostgresSource {
		if value, exists := os.LookupEnv(DatabaseEnv); exists {
			config.DatabaseUrl = value
		} else {
			return Config{}, MissingEnvVar{EnvVar: DatabaseEnv}
		}
	}

	return config, nil
}
