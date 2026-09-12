package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"encoding/json"
	"net/http"
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
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	player, err := services.GetPlayerById(id)
	if err != nil {
		switch err.Error() {
		case "PLAYER_NOT_FOUND":
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)

	return
}
