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

	if strings.HasPrefix(command, config.DISCORD_COMMAND_ALIAS_START_SERVER) {
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

	if strings.HasPrefix(command, config.DISCORD_COMMAND_ALIAS_KICK) {
		return kick(client, command)
	} else if strings.HasPrefix(command, config.DISCORD_COMMAND_ALIAS_BAN) {
		return ban(client, command)
	} else if strings.HasPrefix(command, config.DISCORD_COMMAND_ALIAS_BROADCAST) {
		return broadcast(client, command)
	} else if strings.HasPrefix(command, config.DISCORD_COMMAND_ALIAS_SHUTDOWN) {
		return shutdown(client, command)
	} else if strings.HasPrefix(command, config.DISCORD_COMMAND_ALIAS_DO_EXIT) {
		return doExit(client)
	} else if strings.HasPrefix(command, config.DISCORD_COMMAND_ALIAS_SAVE) {
		return save(client)
	}

	return i18n.UnknownCommand
}