package config

import (
	"os"
)

type Config struct {
	AppEnv            string
	Port              string
	MongoURI          string
	MongoDB           string
	SupabaseJWTSecret string
}

func Load() *Config {
	return &Config{
		AppEnv:            getEnv("APP_ENV", "development"),
		Port:              getEnv("PORT", "8081"),
		MongoURI:          getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:           getEnv("MONGO_DB", "online_courses"),
		SupabaseJWTSecret: getEnv("SUPABASE_JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
