package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                   string
	DBURL                  string
	GinMode                string
	FirebaseCredentialPath string
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	return &Config{
		Port:                   getEnv("PORT", "8080"),
		DBURL:                  getEnv("DB_URL", ""),
		GinMode:                getEnv("GIN_MODE", "debug"),
		FirebaseCredentialPath: getEnv("FIREBASE_CREDENTIAL_PATH", "./serviceAccountKey.json"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
