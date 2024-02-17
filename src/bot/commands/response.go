package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/config"
)

func Response(message string) string {
	command, _ := strings.CutPrefix(message, config.DISCORD_COMMAND_PREFIX+" ")

	if command == "help" {
		return help()
	} else if command == "kick" {
		return kick(command)
	} else if command == "ban" {
		return ban(command)
	} else if strings.HasPrefix(command, "broadcast") {
		return broadcast(command)
	} else if strings.HasPrefix(command, "shutdown") {
		return shutdown(command)
	} else if command == "doExit" {
		return doExit()
	} else if command == "save" {
		return save()
	}

	return "Unknown command"
}
