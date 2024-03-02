package commands

import (
	"strings"
	"time"

	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/palworld/settings"
	"github.com/juunini/palworld-discord-bot/src/utils"
	palworldrcon "github.com/juunini/palworld-rcon"
)

func Response(message string, username string) []string {
	isAdmin := utils.IsAdmin(username)

	command, found := strings.CutPrefix(message, config.DISCORD_COMMAND_PREFIX+" ")
	processedCommand := commandForCheck(command)

	if strings.HasPrefix(processedCommand, config.DISCORD_COMMAND_ALIAS_HELP) || !found {
		return []string{i18n.Help(i18n.HelpParams{
			CommandPrefix:       config.DISCORD_COMMAND_PREFIX,
			HelpAlias:           config.DISCORD_COMMAND_ALIAS_HELP,
			KickAlias:           config.DISCORD_COMMAND_ALIAS_KICK,
			BanAlias:            config.DISCORD_COMMAND_ALIAS_BAN,
			BroadcastAlias:      config.DISCORD_COMMAND_ALIAS_BROADCAST,
			ShutdownAlias:       config.DISCORD_COMMAND_ALIAS_SHUTDOWN,
			DoExitAlias:         config.DISCORD_COMMAND_ALIAS_DO_EXIT,
			SaveAlias:           config.DISCORD_COMMAND_ALIAS_SAVE,
			StartServerAlias:    config.DISCORD_COMMAND_ALIAS_START_SERVER,
			ServerSettingsAlias: config.DISCORD_COMMAND_ALIAS_SERVER_SETTINGS,
		}, isAdmin)}
	}

	// Under commands, only admins can execute
	if !isAdmin {
		return []string{i18n.UnknownCommand}
	}

	if strings.HasPrefix(processedCommand, config.DISCORD_COMMAND_ALIAS_START_SERVER) {
		return []string{startServer()}
	}

	if strings.HasPrefix(processedCommand, config.DISCORD_COMMAND_ALIAS_SERVER_SETTINGS) {
		return settings.Response(command)
	}

	client, err := palworldrcon.Connect(
		config.PALWORLD_RCON_HOST,
		config.PALWORLD_RCON_PORT,
		config.PALWORLD_ADMIN_PASSWORD,
		5*time.Second,
	)
	if err != nil {
		return []string{i18n.FailedToConnectRconServer}
	}
	defer client.Disconnect()

	if strings.HasPrefix(processedCommand, config.DISCORD_COMMAND_ALIAS_KICK) {
		return []string{kick(client, command)}
	} else if strings.HasPrefix(processedCommand, config.DISCORD_COMMAND_ALIAS_BAN) {
		return []string{ban(client, command)}
	} else if strings.HasPrefix(processedCommand, config.DISCORD_COMMAND_ALIAS_BROADCAST) {
		return []string{broadcast(client, command)}
	} else if strings.HasPrefix(processedCommand, config.DISCORD_COMMAND_ALIAS_SHUTDOWN) {
		return []string{shutdown(client, command)}
	} else if strings.HasPrefix(processedCommand, config.DISCORD_COMMAND_ALIAS_DO_EXIT) {
		return []string{doExit(client)}
	} else if strings.HasPrefix(processedCommand, config.DISCORD_COMMAND_ALIAS_SAVE) {
		return []string{save(client)}
	}

	return []string{i18n.UnknownCommand}
}

func commandForCheck(command string) string {
	if !config.DISCORD_COMMAND_CASE_SENSITIVE {
		return strings.ToLower(command)
	}
	return command
}
