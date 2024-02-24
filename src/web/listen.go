package web

import (
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
)

//go:embed public/*
var public embed.FS

var app *fiber.App

func Listen(port int) error {
	app = fiber.New(fiber.Config{
		Views: html.NewFileSystem(http.FS(public), ".html"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("public/index", fiber.Map{
			"CONFIG": fiber.Map{
				"WEB_SERVER_ENABLED": config.WEB_SERVER_ENABLED,
				"WEB_SERVER_PORT":    config.WEB_SERVER_PORT,

				"LANGUAGE": config.LANGUAGE,

				"DISCORD_BOT_ENABLED":         config.DISCORD_BOT_ENABLED,
				"DISCORD_BOT_TOKEN":           config.DISCORD_BOT_TOKEN,
				"DISCORD_ADMIN_USERNAMES":     strings.Join(config.DISCORD_ADMIN_USERNAMES, ", "),
				"DISCORD_COMMAND_IGNORE_CASE": config.DISCORD_COMMAND_IGNORE_CASE,
				"DISCORD_COMMAND_PREFIX":      config.DISCORD_COMMAND_PREFIX,

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

				"Close":   i18n.Close,
				"Confirm": i18n.Confirm,
				"Save":    i18n.Save,
			},
		})
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(public),
		PathPrefix: "public",
		Browse:     true,
	}))

	return app.Listen(fmt.Sprintf(":%d", port))
}
