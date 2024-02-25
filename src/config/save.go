package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/juunini/palworld-discord-bot/src/i18n"
)

func Save() error {
	env := "WEB_SERVER_ENABLED=" + strconv.FormatBool(WEB_SERVER_ENABLED) + "\n"
	env += "# " + i18n.WebServerPortTooltip + "\n"
	env += "WEB_SERVER_PORT=" + strconv.Itoa(WEB_SERVER_PORT) + "\n\n"

	env += "DISCORD_BOT_ENABLED=" + strconv.FormatBool(DISCORD_BOT_ENABLED) + "\n"
	env += "# " + i18n.DiscordBotTokenDescription + "\n"
	env += "DISCORD_BOT_TOKEN=" + DISCORD_BOT_TOKEN + "\n"
	env += "# " + i18n.DiscordAdminUsernamesTooltip + "\n"
	env += "DISCORD_ADMIN_USERNAMES=" + strings.Join(DISCORD_ADMIN_USERNAMES, ",") + "\n"
	env += "# " + i18n.DiscordCommandCaseSensitiveTooltip + "\n"
	env += "DISCORD_COMMAND_CASE_SENSITIVE=" + strconv.FormatBool(DISCORD_COMMAND_CASE_SENSITIVE) + "\n"
	env += "# " + i18n.DiscordCommandPrefixTooltip + "\n"
	env += "DISCORD_COMMAND_PREFIX=" + DISCORD_COMMAND_PREFIX + "\n\n"

	env += "# " + i18n.EnablePalworldRconTooltip + "\n"
	env += "PALWORLD_RCON_ENABLED=" + strconv.FormatBool(PALWORLD_RCON_ENABLED) + "\n"
	env += "# " + i18n.PalworldRconHostTooltip + "\n"
	env += "PALWORLD_RCON_HOST=" + PALWORLD_RCON_HOST + "\n"
	env += "# " + i18n.PalworldRconPortTooltip + "\n"
	env += "PALWORLD_RCON_PORT=" + strconv.Itoa(int(PALWORLD_RCON_PORT)) + "\n"
	env += "# " + i18n.PalworldAdminPasswordTooltip + "\n"
	env += "PALWORLD_ADMIN_PASSWORD=" + PALWORLD_ADMIN_PASSWORD + "\n"
	env += "# " + i18n.PalworldServerFilePathTooltip + "\n"
	env += "PALWORLD_SERVER_FILE_PATH=" + PALWORLD_SERVER_FILE_PATH + "\n"
	env += "# " + i18n.PalworldServerExecuteFlagsTooltip + "\n"
	env += "PALWORLD_SERVER_EXECUTE_FLAGS=" + strings.Join(PALWORLD_SERVER_EXECUTE_FLAGS, " ") + "\n\n"

	env += "# " + i18n.DiscordDashboardChannelIDDescription + "\n"
	env += "DISCORD_DASHBOARD_CHANNEL_ID=" + DISCORD_DASHBOARD_CHANNEL_ID + "\n"
	env += "# " + i18n.DiscordLogChannelIDDescription + "\n"
	env += "DISCORD_LOG_CHANNEL_ID=" + DISCORD_LOG_CHANNEL_ID + "\n"
	env += "# " + i18n.DiscordDashboardPalworldConfigMessageIDTooltip + "\n"
	env += "DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID=" + DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID + "\n"
	env += "# " + i18n.DiscordDashboardOnlinePlayersMessageIDTooltip + "\n"
	env += "DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID=" + DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID + "\n"
	env += "# " + i18n.DiscordDashboardBotConfigMessageIDTooltip + "\n"
	env += "DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID=" + DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID + "\n\n"

	env += "# " + i18n.DiscordCommandAliasHelpTooltip + "\n"
	env += "DISCORD_COMMAND_ALIAS_HELP=" + DISCORD_COMMAND_ALIAS_HELP + "\n"
	env += "# " + i18n.DiscordCommandAliasKickTooltip + "\n"
	env += "DISCORD_COMMAND_ALIAS_KICK=" + DISCORD_COMMAND_ALIAS_KICK + "\n"
	env += "# " + i18n.DiscordCommandAliasBanTooltip + "\n"
	env += "DISCORD_COMMAND_ALIAS_BAN=" + DISCORD_COMMAND_ALIAS_BAN + "\n"
	env += "# " + i18n.DiscordCommandAliasBroadcastTooltip + "\n"
	env += "DISCORD_COMMAND_ALIAS_BROADCAST=" + DISCORD_COMMAND_ALIAS_BROADCAST + "\n"
	env += "# " + i18n.DiscordCommandAliasShutdownTooltip + "\n"
	env += "DISCORD_COMMAND_ALIAS_SHUTDOWN=" + DISCORD_COMMAND_ALIAS_SHUTDOWN + "\n"
	env += "# " + i18n.DiscordCommandAliasDoExitTooltip + "\n"
	env += "DISCORD_COMMAND_ALIAS_DO_EXIT=" + DISCORD_COMMAND_ALIAS_DO_EXIT + "\n"
	env += "# " + i18n.DiscordCommandAliasSaveTooltip + "\n"
	env += "DISCORD_COMMAND_ALIAS_SAVE=" + DISCORD_COMMAND_ALIAS_SAVE + "\n"
	env += "# " + i18n.DiscordCommandAliasStartServerTooltip + "\n"
	env += "DISCORD_COMMAND_ALIAS_START_SERVER=" + DISCORD_COMMAND_ALIAS_START_SERVER + "\n\n"

	env += "LANGUAGE=" + LANGUAGE + "\n"

	return os.WriteFile(".env", []byte(env), 0644)
}
