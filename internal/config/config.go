package config

import "os"

type Config struct {
	Port         string
	AppEnv       string
	BaseURL      string
	ContactEmail string
	ResendAPIKey string
	EmailFrom    string
}

func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", "3000"),
		AppEnv:       getEnv("APP_ENV", "development"),
		BaseURL:      getEnv("BASE_URL", "https://logam.gold"),
		ContactEmail: getEnv("CONTACT_EMAIL", "wiliamjones@pm.me"),
		ResendAPIKey: getEnv("RESEND_API_KEY", ""),
		EmailFrom:    getEnv("EMAIL_FROM", "Logam Gold <onboarding@resend.dev>"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
