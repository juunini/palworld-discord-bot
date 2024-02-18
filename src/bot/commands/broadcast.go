package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
)

func broadcast(command string) string {
	message, found := strings.CutPrefix(command, "broadcast ")

	if !found {
		return i18n.WrongParameters
	}

	client, err := utils.RconClient()
	if err != nil {
		return i18n.FailedToConnectRconServer
	}

	defer client.Disconnect()

	response, err := client.Broadcast(strings.ReplaceAll(message, " ", "_"))
	if err != nil {
		return i18n.FailedToBroadcastCommand
	}

	return response
}
