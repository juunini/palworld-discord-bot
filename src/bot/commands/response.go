package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
)

func Response(message string, username string) string {
	isAdmin := utils.IsAdmin(username)

	command, found := strings.CutPrefix(message, config.DISCORD_COMMAND_PREFIX+" ")
	if command == "help" || !found {
		return i18n.Help(config.DISCORD_COMMAND_PREFIX, isAdmin)
	}

	// Under commands, only admins can execute
	if !isAdmin {
		return i18n.UnknownCommand
	}

	if strings.HasPrefix(command, "kick") {
		return kick(command)
	} else if strings.HasPrefix(command, "ban") {
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

	return i18n.UnknownCommand
}
