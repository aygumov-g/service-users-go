package config

import "os"

type Config struct {
	AppPort string
}

func Load() *Config {
	return &Config{
		AppPort: os.Getenv("APP_PORT"),
	}
}
