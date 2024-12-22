package config

import (
	"fmt"
	"os"
	"sync"

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

type OpenAiConfig struct {
	APIKey      string
	AssistantID string
	URLBase     string
}

type Config struct {
	Redis        RedisConfig
	WhatsApp     WhatsAppConfig
	OpenAiConfig OpenAiConfig
}

var (
	cfg  *Config
	once sync.Once
	err  error
)

// Initialize loads the configuration only once.
func Initialize() error {
	once.Do(func() {
		if err = godotenv.Load(); err != nil {
			err = fmt.Errorf("error loading .env file: %w", err)
			return 
		}

		cfg = &Config{
			Redis: RedisConfig{
				Addr:     os.Getenv("REDIS_ADDR"),
				Password: os.Getenv("REDIS_PASSWORD"),
				DB:       0,
			},
			WhatsApp: WhatsAppConfig{
				Dialect: os.Getenv("DB_DIALECT"),
				DSN:     os.Getenv("DB_DSN"),
			},
			OpenAiConfig: OpenAiConfig{
				APIKey:      os.Getenv("OPENAI_API_KEY"),
				AssistantID: os.Getenv("ASSISTANT_ID"),
				URLBase:     os.Getenv("OPENAI_URL_BASE"),
			},
		}
	})

	if err != nil {
		return fmt.Errorf("failed to initialize configuration: %w", err)
	}

	return nil
}

// GetConfig returns the initialized configuration.
func GetConfig() *Config {
	if cfg == nil {
		return nil
	}
	return cfg
}
