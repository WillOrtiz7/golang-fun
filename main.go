package main

import (
	"awesomeProject/handlers"
	"awesomeProject/utils"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	utils.InitEnv()

	utils.InitDatabase()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /player/{id}", handlers.GetPlayer)

	logger.Info("Server running on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
