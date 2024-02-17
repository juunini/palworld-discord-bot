package watchers

import (
	"encoding/csv"
	"strings"

	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
)

var LeftPlayers = map[string]Player{}
var OnlinePlayers = map[string]Player{}
var NewPlayers = []Player{}

type Player struct {
	Username string
	PlayerID string
	SteamID  string
}

func Players() error {
	client, err := utils.RconClient()
	if err != nil {
		console_decoration.PrintError(i18n.FailedToConnectRconServer)
		return err
	}

	response, err := client.ShowPlayers()
	if err != nil {
		console_decoration.PrintError(i18n.FailedToShowPlayerCommand)
		return err
	}

	rows, err := csv.NewReader(strings.NewReader(response)).ReadAll()
	if err != nil {
		return err
	}

	LeftPlayers = OnlinePlayers
	OnlinePlayers = map[string]Player{}

	for _, row := range rows[1:] {
		player := Player{
			Username: row[0],
			PlayerID: row[1],
			SteamID:  row[2],
		}
		OnlinePlayers[player.SteamID] = player

		_, ok := LeftPlayers[player.SteamID]
		if ok {
			delete(LeftPlayers, player.SteamID)
			continue
		}

		NewPlayers = append(NewPlayers, player)
	}

	return nil
}
