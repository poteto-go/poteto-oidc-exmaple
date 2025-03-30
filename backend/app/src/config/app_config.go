package config

import (
	"github.com/joho/godotenv"
)

type ApplicationConfig struct {
	AppPort       string
	JWKsUrl       string
	ClientId      string
	ClientSecret  string
	RedirectURL   string
	TokenEndpoint string
}

var AppConfig ApplicationConfig

func InitAppConfig() {
	godotenv.Load(".env")

	AppConfig = ApplicationConfig{
		AppPort:       GetEnv("APP_PORT", "3000"),
		JWKsUrl:       GetEnv("JWKS_URL", ""),
		ClientId:      GetEnv("CLIENT_ID", ""),
		ClientSecret:  GetEnv("CLIENT_SECRET", ""),
		RedirectURL:   GetEnv("REDIRECT_URL", ""),
		TokenEndpoint: GetEnv("TOKEN_ENDPOINT", ""),
	}
}
