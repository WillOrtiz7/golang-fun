package utils

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file", err.Error())
	}
}
