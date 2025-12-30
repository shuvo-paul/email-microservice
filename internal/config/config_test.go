package config_test

import (
	"testing"

	"github.com/shuvo-paul/email-microservice/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLoad_LoadsFromEnv(t *testing.T) {
	t.Setenv("APP_ENV", "test")
	t.Setenv("SERVER_PORT", "3000")
	t.Setenv("SMTP_HOST", "example.org")
	t.Setenv("SMTP_PORT", "22")
	t.Setenv("SMTP_USERNAME", "jhon")
	t.Setenv("SMTP_PASSWORD", "abcd")
	t.Setenv("SMTP_FROM", "jhon@example.org")

	cfg := config.Load()

	assert.Equal(t, "test", cfg.AppEnv)
	assert.Equal(t, 3000, cfg.ServerConfig.Port)
	assert.Equal(t, "example.org", cfg.SMTPConfig.Host)
	assert.Equal(t, 22, cfg.SMTPConfig.Port)
	assert.Equal(t, "jhon", cfg.SMTPConfig.Username)
	assert.Equal(t, "abcd", cfg.SMTPConfig.Password)
	assert.Equal(t, "jhon@example.org", cfg.SMTPConfig.From)
}

func TestLoad_UseDefaults(t *testing.T) {
	t.Setenv("SMTP_HOST", "example.org")
	t.Setenv("SMTP_USERNAME", "jhon")
	t.Setenv("SMTP_PASSWORD", "abcd")
	t.Setenv("SMTP_FROM", "jhon@example.org")

	tests := []struct {
		name       string
		ServerPort string
		SMTPPort   string
	}{
		{
			name: "No env",
		},
		{
			name:       "Invalid env",
			ServerPort: "invalid",
			SMTPPort:   "invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ServerPort != "" {
				t.Setenv("SERVER_PORT", tt.ServerPort)
			}
			if tt.SMTPPort != "" {
				t.Setenv("SMTP_PORT", tt.SMTPPort)
			}

			cfg := config.Load()
			assert.Equal(t, 4000, cfg.ServerConfig.Port)
		})
	}
}
