package utils

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitDatabase() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("Unable to connect to database", err.Error())
		os.Exit(1)
	}
	defer conn.Close(context.Background())

}
