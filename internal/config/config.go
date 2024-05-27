package config

import (
	"log"
	"main/internal/services/mistakes"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	PostgresDB struct {
		UserName string
		Password string
		DbName   string
		Host     string
		Port     string
		SslMode  string
	}
	Env struct {
		PostgresDB
		ContextTimeout int
		ServerAddress  string
	}
)

func NewEnv(path string) *Env {
	err := godotenv.Load(path)
	if err != nil {
		log.Println(err)
	}

	return &Env{
		PostgresDB: PostgresDB{
			DbName:   getEnv("DB_NAME"),
			UserName: getEnv("DB_USER"),
			Password: getEnv("DB_PASS"),
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			SslMode:  getEnv("DB_SSL_MODE"),
		},
		ContextTimeout: getEnvAsInt("CONTEXT_TIMEOUT"),
		ServerAddress:  getEnv("SERVER_ADDRESS"),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	log.Println(mistakes.ErrEnv(key))

	return ""
}

func getEnvAsInt(key string) int {
	valStr := getEnv(key)
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}

	log.Println(mistakes.ErrEnv(key))

	return 0
}
