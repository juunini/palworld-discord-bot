package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/i18n"
	palworldrcon "github.com/juunini/palworld-rcon"
)

func ban(client *palworldrcon.Client, command string) string {
	steamID, found := strings.CutPrefix(command, "ban ")

	if !found {
		return i18n.WrongParameters
	}

	response, err := client.BanPlayer(steamID)
	if err != nil {
		return i18n.FailedToBanCommand
	}

	return response
}
