package config

import (
	"os"
)

type Config struct {
	AppPort string
	DB      postgres
}

func Load() *Config {
	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DB: postgres{
			dBHost:     os.Getenv("POSTGRES_HOST"),
			dBUser:     os.Getenv("POSTGRES_USER"),
			dBPassword: os.Getenv("POSTGRES_PASSWORD"),
			dBName:     os.Getenv("POSTGRES_DB"),
		},
	}
}
