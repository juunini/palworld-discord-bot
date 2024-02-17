package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/config"
)

func Response(message string) string {
	command, _ := strings.CutPrefix(message, config.DISCORD_COMMAND_PREFIX+" ")

	if command == "help" {
		return Help()
	}

	return "Unknown command"
}
