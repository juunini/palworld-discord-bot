package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
)

func Index(c *fiber.Ctx) error {
	return c.Render("public/index", fiber.Map{
		"CONFIG": fiber.Map{
			"WEB_SERVER_ENABLED": config.WEB_SERVER_ENABLED,
			"WEB_SERVER_PORT":    config.WEB_SERVER_PORT,

			"LANGUAGE": config.LANGUAGE,

			"DISCORD_BOT_ENABLED":            config.DISCORD_BOT_ENABLED,
			"DISCORD_BOT_TOKEN":              config.DISCORD_BOT_TOKEN,
			"DISCORD_ADMIN_USERNAMES":        strings.Join(config.DISCORD_ADMIN_USERNAMES, ", "),
			"DISCORD_COMMAND_CASE_SENSITIVE": config.DISCORD_COMMAND_CASE_SENSITIVE,
			"DISCORD_COMMAND_PREFIX":         config.DISCORD_COMMAND_PREFIX,

			"PALWORLD_RCON_ENABLED":         config.PALWORLD_RCON_ENABLED,
			"PALWORLD_RCON_HOST":            config.PALWORLD_RCON_HOST,
			"PALWORLD_RCON_PORT":            config.PALWORLD_RCON_PORT,
			"PALWORLD_ADMIN_PASSWORD":       config.PALWORLD_ADMIN_PASSWORD,
			"PALWORLD_SERVER_FILE_PATH":     config.PALWORLD_SERVER_FILE_PATH,
			"PALWORLD_SERVER_EXECUTE_FLAGS": strings.Join(config.PALWORLD_SERVER_EXECUTE_FLAGS, " "),

			"DISCORD_DASHBOARD_CHANNEL_ID":                 config.DISCORD_DASHBOARD_CHANNEL_ID,
			"DISCORD_LOG_CHANNEL_ID":                       config.DISCORD_LOG_CHANNEL_ID,
			"DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID":  config.DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID,
			"DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID": config.DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID,
			"DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID":      config.DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID,

			"DISCORD_COMMAND_ALIAS_KICK":         config.DISCORD_COMMAND_ALIAS_KICK,
			"DISCORD_COMMAND_ALIAS_BAN":          config.DISCORD_COMMAND_ALIAS_BAN,
			"DISCORD_COMMAND_ALIAS_BROADCAST":    config.DISCORD_COMMAND_ALIAS_BROADCAST,
			"DISCORD_COMMAND_ALIAS_SHUTDOWN":     config.DISCORD_COMMAND_ALIAS_SHUTDOWN,
			"DISCORD_COMMAND_ALIAS_DO_EXIT":      config.DISCORD_COMMAND_ALIAS_DO_EXIT,
			"DISCORD_COMMAND_ALIAS_SAVE":         config.DISCORD_COMMAND_ALIAS_SAVE,
			"DISCORD_COMMAND_ALIAS_START_SERVER": config.DISCORD_COMMAND_ALIAS_START_SERVER,
		},
		"i18n": fiber.Map{
			"WebConfig":                    i18n.WebConfig,
			"EnableWebServer":              i18n.EnableWebServer,
			"EnableWebServerDisablePrompt": i18n.EnableWebServerDisablePrompt,
			"WebServerPort":                i18n.WebServerPort,
			"WebServerPortTooltip":         i18n.WebServerPortTooltip,

			"Language":       i18n.Language,
			"Languages":      i18n.Languages,
			"SelectLanguage": i18n.SelectLanguage,

			"DiscordBotConfig":                   i18n.DiscordBotConfig,
			"EnableDiscordBot":                   i18n.EnableDiscordBot,
			"DiscordBotToken":                    i18n.DiscordBotToken,
			"DiscordBotTokenTooltip":             i18n.DiscordBotTokenTooltip,
			"DiscordAdminUsernames":              i18n.DiscordAdminUsernames,
			"DiscordAdminUsernamesTooltip":       i18n.DiscordAdminUsernamesTooltip,
			"DiscordCommandCaseSensitive":        i18n.DiscordCommandCaseSensitive,
			"DiscordCommandCaseSensitiveTooltip": i18n.DiscordCommandCaseSensitiveTooltip,
			"DiscordCommandPrefix":               i18n.DiscordCommandPrefix,
			"DiscordCommandPrefixTooltip":        i18n.DiscordCommandPrefixTooltip,

			"PalworldConfig":                    i18n.PalworldConfig,
			"EnablePalworldRcon":                i18n.EnablePalworldRcon,
			"EnablePalworldRconTooltip":         i18n.EnablePalworldRconTooltip,
			"PalworldRconHost":                  i18n.PalworldRconHost,
			"PalworldRconHostTooltip":           i18n.PalworldRconHostTooltip,
			"PalworldRconPort":                  i18n.PalworldRconPort,
			"PalworldRconPortTooltip":           i18n.PalworldRconPortTooltip,
			"PalworldAdminPassword":             i18n.PalworldAdminPassword,
			"PalworldAdminPasswordTooltip":      i18n.PalworldAdminPasswordTooltip,
			"PalworldServerFilePath":            i18n.PalworldServerFilePath,
			"PalworldServerFilePathTooltip":     i18n.PalworldServerFilePathTooltip,
			"PalworldServerExecuteFlags":        i18n.PalworldServerExecuteFlags,
			"PalworldServerExecuteFlagsTooltip": i18n.PalworldServerExecuteFlagsTooltip,

			"DiscordChannelConfig":                           i18n.DiscordChannelConfig,
			"DiscordDashboardChannelID":                      i18n.DiscordDashboardChannelID,
			"DiscordDashboardChannelIDTooltip":               i18n.DiscordDashboardChannelIDTooltip,
			"DiscordLogChannelID":                            i18n.DiscordLogChannelID,
			"DiscordLogChannelIDTooltip":                     i18n.DiscordLogChannelIDTooltip,
			"DiscordDashboardOnlinePlayersMessageID":         i18n.DiscordDashboardOnlinePlayersMessageID,
			"DiscordDashboardOnlinePlayersMessageIDTooltip":  i18n.DiscordDashboardOnlinePlayersMessageIDTooltip,
			"DiscordDashboardPalworldConfigMessageID":        i18n.DiscordDashboardPalworldConfigMessageID,
			"DiscordDashboardPalworldConfigMessageIDTooltip": i18n.DiscordDashboardPalworldConfigMessageIDTooltip,
			"DiscordDashboardBotConfigMessageID":             i18n.DiscordDashboardBotConfigMessageID,
			"DiscordDashboardBotConfigMessageIDTooltip":      i18n.DiscordDashboardBotConfigMessageIDTooltip,

			"DiscordCommandAliases":                 i18n.DiscordCommandAliases,
			"DiscordCommandAliasKickTooltip":        i18n.DiscordCommandAliasKickTooltip,
			"DiscordCommandAliasBanTooltip":         i18n.DiscordCommandAliasBanTooltip,
			"DiscordCommandAliasBroadcastTooltip":   i18n.DiscordCommandAliasBroadcastTooltip,
			"DiscordCommandAliasShutdownTooltip":    i18n.DiscordCommandAliasShutdownTooltip,
			"DiscordCommandAliasDoExitTooltip":      i18n.DiscordCommandAliasDoExitTooltip,
			"DiscordCommandAliasSaveTooltip":        i18n.DiscordCommandAliasSaveTooltip,
			"DiscordCommandAliasStartServerTooltip": i18n.DiscordCommandAliasStartServerTooltip,

			"Close":   i18n.Close,
			"Confirm": i18n.Confirm,
			"Save":    i18n.Save,

			"ConfigUpdateSuccess": i18n.ConfigUpdateSuccess,
			"ConfigUpdateFailed":  i18n.ConfigUpdateFailed,
		},
	})
}
