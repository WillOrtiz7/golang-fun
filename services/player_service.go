package services

import (
	"awesomeProject/models"
	"encoding/json"
	"errors"
	"log/slog"
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

func GetPlayerById(id string) (*models.Player, error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	for _, player := range players {
		if strconv.Itoa(player.Id) == id {
			_, err := json.Marshal(player)
			if err != nil {
				logger.Error("Error encoding player to JSON", err)
				return nil, err
			}
			return &player, nil
		}
	}
	return nil, errors.New("PLAYER_NOT_FOUND")
}
