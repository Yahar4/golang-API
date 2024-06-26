package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config is for .env varibales
type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

// Envs is reusable, so i don't have to write initConfig every time
var Envs = initConfig()

func initConfig() Config {
	// here we do load .enc variables
	err := godotenv.Load()
	if err != nil {
		return Config{}
	}

	//this is a connection config for DB
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "9090"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "hello_123"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "ecommerce"),
	}
}

// this loads params for DB in main.go (all for safety :thumbup)
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
