package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
)

func Response(message string, username string) string {
	command, _ := strings.CutPrefix(message, config.DISCORD_COMMAND_PREFIX+" ")
	isAdmin := utils.IsAdmin(username)

	if command == "help" {
		return i18n.Help(config.DISCORD_COMMAND_PREFIX, isAdmin)
	}

	// Under commands, only admins can execute
	if !isAdmin {
		return i18n.UnknownCommand
	}

	if command == "kick" {
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

	return i18n.UnknownCommand
}
