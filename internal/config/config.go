package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/gorilla/securecookie"
)

type Config struct {
	Host           string
	Port           string
	AllowedOrigins []string
	CookieName     string
	CookieSecret   []byte
}

func New() Config {
	host := getEnvDefault("HOST", "0.0.0.0")
	port := getEnvDefault("PORT", "8080")
	return Config{
		Host: host,
		Port: port,
		AllowedOrigins: strings.Split(
			getEnvDefault("ALLOWED_ORIGINS", fmt.Sprintf("http://%s:%s,https://%s:%s", host, port, host, port)), ",",
		),
		CookieName:   getEnvDefault("COOKIE_NAME", "session"),
		CookieSecret: getCookieSecret("COOKIE_SECRET"),
	}
}

func getEnvDefault(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}

func getCookieSecret(key string) []byte {
	if value := os.Getenv(key); value != "" {
		return []byte(value)
	}
	return securecookie.GenerateRandomKey(32)
}
