package commands

import (
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	palworldrcon "github.com/juunini/palworld-rcon"
)

func doExit(client *palworldrcon.Client) string {
	response, err := client.DoExit()
	if err != nil {
		console_decoration.PrintError(i18n.FailedToDoExitCommand + ": " + err.Error())
		return i18n.FailedToDoExitCommand
	}

	return response
}
