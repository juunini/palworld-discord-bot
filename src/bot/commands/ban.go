package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
)

func ban(command string) string {
	steamID, found := strings.CutPrefix(command, "ban ")

	if !found {
		return i18n.WrongParameters
	}

	client, err := utils.RconClient()
	if err != nil {
		return i18n.FailedToConnectRconServer
	}

	defer client.Disconnect()

	response, err := client.BanPlayer(steamID)
	if err != nil {
		return i18n.FailedToBanCommand
	}

	return response
}
