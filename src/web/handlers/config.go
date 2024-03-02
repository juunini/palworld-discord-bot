package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/juunini/palworld-discord-bot/src/config"
)

func PatchConfig(shutdownServer func() error) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		changedConfig := make(map[string]interface{})
		if err := c.BodyParser(&changedConfig); err != nil {
			return err
		}

		for key, value := range changedConfig {
			switch key {
			case "WEB_SERVER_ENABLED":
				config.WEB_SERVER_ENABLED = value.(bool)
			case "WEB_SERVER_PORT":
				config.WEB_SERVER_PORT = int(value.(float64))
			case "LANGUAGE":
				config.LANGUAGE = value.(string)
			case "DISCORD_BOT_ENABLED":
				config.DISCORD_BOT_ENABLED = value.(bool)
			case "DISCORD_BOT_TOKEN":
				config.DISCORD_BOT_TOKEN = value.(string)
			case "DISCORD_ADMIN_USERNAMES":
				config.DISCORD_ADMIN_USERNAMES = strings.Split(value.(string), ", ")
			case "DISCORD_COMMAND_CASE_SENSITIVE":
				config.DISCORD_COMMAND_CASE_SENSITIVE = value.(bool)
			case "DISCORD_COMMAND_PREFIX":
				config.DISCORD_COMMAND_PREFIX = value.(string)
			case "PALWORLD_RCON_ENABLED":
				config.PALWORLD_RCON_ENABLED = value.(bool)
			case "PALWORLD_RCON_HOST":
				config.PALWORLD_RCON_HOST = value.(string)
			case "PALWORLD_RCON_PORT":
				config.PALWORLD_RCON_PORT = int(value.(float64))
			case "PALWORLD_ADMIN_PASSWORD":
				config.PALWORLD_ADMIN_PASSWORD = value.(string)
			case "PALWORLD_SERVER_FILE_PATH":
				config.PALWORLD_SERVER_FILE_PATH = value.(string)
			case "PALWORLD_SERVER_EXECUTE_FLAGS":
				config.PALWORLD_SERVER_EXECUTE_FLAGS = strings.Split(value.(string), " ")
			case "PALWORLD_SERVER_SETTINGS_FILE_PATH":
				config.PALWORLD_SERVER_SETTINGS_FILE_PATH = value.(string)
			case "DISCORD_DASHBOARD_CHANNEL_ID":
				config.DISCORD_DASHBOARD_CHANNEL_ID = value.(string)
			case "DISCORD_LOG_CHANNEL_ID":
				config.DISCORD_LOG_CHANNEL_ID = value.(string)
			case "DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID":
				config.DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID = value.(string)
			case "DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID":
				config.DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID = value.(string)
			case "DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID":
				config.DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID = value.(string)
			case "DISCORD_COMMAND_ALIAS_HELP":
				config.DISCORD_COMMAND_ALIAS_HELP = value.(string)
			case "DISCORD_COMMAND_ALIAS_KICK":
				config.DISCORD_COMMAND_ALIAS_KICK = value.(string)
			case "DISCORD_COMMAND_ALIAS_BAN":
				config.DISCORD_COMMAND_ALIAS_BAN = value.(string)
			case "DISCORD_COMMAND_ALIAS_BROADCAST":
				config.DISCORD_COMMAND_ALIAS_BROADCAST = value.(string)
			case "DISCORD_COMMAND_ALIAS_SHUTDOWN":
				config.DISCORD_COMMAND_ALIAS_SHUTDOWN = value.(string)
			case "DISCORD_COMMAND_ALIAS_DO_EXIT":
				config.DISCORD_COMMAND_ALIAS_DO_EXIT = value.(string)
			case "DISCORD_COMMAND_ALIAS_SAVE":
				config.DISCORD_COMMAND_ALIAS_SAVE = value.(string)
			case "DISCORD_COMMAND_ALIAS_START_SERVER":
				config.DISCORD_COMMAND_ALIAS_START_SERVER = value.(string)
			}
		}

		config.Save()

		if err := c.SendStatus(fiber.StatusAccepted); err != nil {
			return err
		}

		if !config.WEB_SERVER_ENABLED {
			go shutdownServer()
		}

		return nil
	}
}
