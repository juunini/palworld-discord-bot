package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/config"
)

func Response(message string) string {
	command, _ := strings.CutPrefix(message, config.DISCORD_COMMAND_PREFIX+" ")

	if command == "help" {
		return Help()
	} else if command == "kick" {

	} else if command == "ban" {

	} else if strings.HasPrefix(command, "broadcast") {

	} else if strings.HasPrefix(command, "shutdown") {

	} else if command == "force-shutdown" {

	} else if command == "save" {

	}

	return "Unknown command"
}
