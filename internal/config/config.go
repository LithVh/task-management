package config

import (
	"fmt"
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

func Load() (*Config, error) {
	_ = godotenv.Load()

	config := &Config{
		Server: ServerConfig{
			Port: getEnvVal("PORT", "8080"),
			Env:  getEnvVal("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnvVal("DB_HOST", "localhost"),
			Port:     getEnvVal("DB_PORT", "5432"),
			User:     getEnvVal("DB_USER", "postgres"),
			Password: getEnvVal("DB_PASSWORD", ""),
			DBName:   getEnvVal("DB_NAME", "task_management"),
		},
		JWT: JWTConfig{
			Secret:      getEnvVal("JWT_SECRET", ""),
			ExpireHours: getEnvIntVal("JWT_EXPIRE_HOURS", 24),
		},
		CORS: CORSConfig{
			Origin: getEnvVal("CORS_ORIGIN", "*"),
		},
	}
	// fmt.Println(config.Database.Password, config.JWT.Secret)

	if config.Database.Password == "" || config.JWT.Secret == "" {
		return nil, fmt.Errorf("Load - unset confidential info")
	}

	return config, nil
}

func getEnvVal(key, defaultVal string) string {
	val := os.Getenv(key)
	// fmt.Println(key, val)
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
