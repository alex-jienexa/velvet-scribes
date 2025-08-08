package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
}

// LoadConfig загружает переменные конфигурации из файла .env и сохраняет их в структуру AppConfig.
func LoadConfig(envFiles ...string) (AppConfig, error) {
	godotenv.Load(envFiles...)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Println("переменная port не установлена, используем значение по умолчанию: ", port)
	}
	return AppConfig{
		Port: port,
	}, nil
}
