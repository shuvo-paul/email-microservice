// Package config handles application configuration loading
package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port int
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type Config struct {
	AppEnv       string
	ServerConfig ServerConfig
	SMTPConfig   SMTPConfig
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		ServerConfig: ServerConfig{
			Port: getEnvAsInt("SERVER_PORT", 4000),
		},
		SMTPConfig: SMTPConfig{
			Host:     getEnv("SMTP_HOST", ""),
			Port:     getEnvAsInt("SMTP_PORT", 587),
			Username: getEnv("SMTP_USERNAME", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
			From:     getEnv("SMTP_FROM", ""),
		},
	}

	validate(cfg)

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")

	i, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return i
}

func validate(cfg *Config) {
	if cfg.SMTPConfig.Host == "" {
		log.Fatal("SMTP_HOST is required")
	}

	if cfg.SMTPConfig.Username == "" {
		log.Fatal("SMTP_USERNAME is required")
	}

	if cfg.SMTPConfig.Password == "" {
		log.Fatal("SMTP_PASSWORD is required")
	}

	if cfg.SMTPConfig.From == "" {
		log.Fatal("SMTP_FROM is required")
	}
}
