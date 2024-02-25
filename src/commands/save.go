package commands

import (
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	palworldrcon "github.com/juunini/palworld-rcon"
)

func save(client *palworldrcon.Client) string {
	response, err := client.Save()
	if err != nil {
		console_decoration.PrintError(i18n.FailedToSaveCommand + ": " + err.Error())
		return i18n.FailedToSaveCommand
	}

	return response
}
