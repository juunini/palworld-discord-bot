package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/i18n"
	palworldrcon "github.com/juunini/palworld-rcon"
)

func kick(client *palworldrcon.Client, command string) string {
	steamID, found := strings.CutPrefix(command, "kick ")

	if !found {
		return i18n.WrongParameters
	}

	response, err := client.KickPlayer(steamID)
	if err != nil {
		return i18n.FailedToKickCommand
	}

	return response
}
