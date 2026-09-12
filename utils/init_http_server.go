package utils

import (
	"awesomeProject/handlers"
	"log/slog"
	"net/http"
	"os"
)

func InitHttpServer() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mux := http.NewServeMux()

	mux.HandleFunc("GET /player/{id}", handlers.GetPlayer)

	logger.Info("Server running on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Error("Unable to start server", err.Error())
		os.Exit(1)
	}
}
