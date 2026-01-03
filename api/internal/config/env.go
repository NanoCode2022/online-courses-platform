package config

import (
	"os"
)

type Config struct {
	AppEnv          string
	Port            string
	MongoURI        string
	MongoDB         string
	SupabaseJWKSURL string
}

func Load() *Config {
	return &Config{
		AppEnv:          getEnv("APP_ENV", "development"),
		Port:            getEnv("PORT", "8081"),
		MongoURI:        getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:         getEnv("MONGO_DB", "online_courses"),
		SupabaseJWKSURL: getEnv("SUPABASE_JWKS_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
