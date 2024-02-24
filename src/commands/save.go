package commands

import (
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
)

func save() string {
	client, err := utils.RconClient()
	if err != nil {
		return i18n.FailedToConnectRconServer
	}

	defer client.Disconnect()

	response, err := client.Save()
	if err != nil {
		console_decoration.PrintError(i18n.FailedToSaveCommand + ": " + err.Error())
		return i18n.FailedToSaveCommand
	}

	return response
}
