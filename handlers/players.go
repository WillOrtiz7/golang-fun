package handlers

import (
	"awesomeProject/models"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

var players = []models.Player{
	{
		Id:        1,
		FirstName: "James",
		LastName:  "Rodriguez",
	},
	{
		Id:        2,
		FirstName: "Carlos",
		LastName:  "Valderrama",
	},
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	for _, player := range players {
		if strconv.Itoa(player.Id) == id {
			bytes, err := json.Marshal(player)
			if err != nil {
				logger.Error("Error encoding player to JSON", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(bytes)
			if err != nil {
				return
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)

}
