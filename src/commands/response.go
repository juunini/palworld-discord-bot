package commands

import (
	"strings"
	"time"

	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
	palworldrcon "github.com/juunini/palworld-rcon"
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

	if command == "startServer" {
		return startServer()
	}

	client, err := palworldrcon.Connect(
		config.PALWORLD_RCON_HOST,
		config.PALWORLD_RCON_PORT,
		config.PALWORLD_ADMIN_PASSWORD,
		5*time.Second,
	)
	if err != nil {
		return i18n.FailedToConnectRconServer
	}
	defer client.Disconnect()

	if strings.HasPrefix(command, "kick") {
		return kick(client, command)
	} else if strings.HasPrefix(command, "ban") {
		return ban(client, command)
	} else if strings.HasPrefix(command, "broadcast") {
		return broadcast(client, command)
	} else if strings.HasPrefix(command, "shutdown") {
		return shutdown(client, command)
	} else if command == "doExit" {
		return doExit(client)
	} else if command == "save" {
		return save(client)
	}

	return i18n.UnknownCommand
}
