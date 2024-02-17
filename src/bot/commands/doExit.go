package commands

import (
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
)

func doExit() string {
	client, err := utils.RconClient()
	if err != nil {
		return i18n.FailedToConnectRconServer
	}

	defer client.Disconnect()

	response, err := client.DoExit()
	if err != nil {
		console_decoration.PrintError(i18n.FailedToDoExitCommand + ": " + err.Error())
		return i18n.FailedToDoExitCommand
	}

	return response
}
