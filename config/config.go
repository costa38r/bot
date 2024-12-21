package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type RedisConfig struct {
    Addr     string
    Password string
    DB       int
}

type WhatsAppConfig struct {
    Dialect string
    DSN     string
}

type Config struct {
    Redis     RedisConfig
    WhatsApp  WhatsAppConfig
}


func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := &Config{
		Redis: RedisConfig{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		},

		WhatsApp: WhatsAppConfig{
			Dialect: os.Getenv("DB_DIALECT"),
			DSN:    os.Getenv("DB_DSN"),
		},
	}

	
	return cfg
}