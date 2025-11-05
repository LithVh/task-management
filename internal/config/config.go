package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	CORS     CORSConfig
}

type ServerConfig struct {
	Port string
	Env  string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

type CORSConfig struct {
	Origin string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Port: getEnvVal("PORT", "8080"),
			Env:  getEnvVal("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnvVal("DB_HOST", "localhost"),
			Port:     getEnvVal("DB_PORT", "5432"),
			User:     getEnvVal("DB_HOST", "postgres"),
			Password: getEnvVal("DB_HOST", "0"),
			DBName:   getEnvVal("DB_HOST", "task_management"),
		},
		JWT: JWTConfig{
			Secret:      getEnvVal("JWT_SECRET", "f3b1c2d4e5f67890123456789abcdef0123456789abcdef0123456789abcdef"),
			ExpireHours: getEnvIntVal("JWT_EXPIRE_HOURS", 24),
		},
		CORS: CORSConfig{
			Origin: getEnvVal("CORS_ORIGIN", "*"),
		},
	}
}

func getEnvVal(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func getEnvIntVal(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	res, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("value %d is not a valid number, switching to default value", res)
	}
	return res
}
