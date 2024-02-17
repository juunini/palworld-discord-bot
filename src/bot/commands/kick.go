package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
)

func kick(command string) string {
	steamID, found := strings.CutPrefix(command, "kick ")

	if !found {
		return i18n.WrongParameters
	}

	client, err := utils.RconClient()
	if err != nil {
		return i18n.FailedToConnectRconServer
	}

	defer client.Disconnect()

	response, err := client.KickPlayer(steamID)
	if err != nil {
		return i18n.FailedToKickCommand
	}

	return response
}
