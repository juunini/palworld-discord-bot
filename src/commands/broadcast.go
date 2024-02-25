package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	palworldrcon "github.com/juunini/palworld-rcon"
)

func broadcast(client *palworldrcon.Client, command string) string {
	message, found := strings.CutPrefix(command, config.DISCORD_COMMAND_ALIAS_BROADCAST+" ")

	if !found {
		return i18n.WrongParameters
	}

	response, err := client.Broadcast(strings.ReplaceAll(message, " ", "_"))
	if err != nil {
		return i18n.FailedToBroadcastCommand
	}

	return response
}
