package config

import (
	"os"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	WEB_SERVER_ENABLED = getEnvOrDefaultBool("WEB_SERVER_ENABLED", WEB_SERVER_ENABLED)
	WEB_SERVER_PORT = getEnvOrDefaultInt("WEB_SERVER_PORT", WEB_SERVER_PORT)

	DISCORD_BOT_ENABLED = getEnvOrDefaultBool("DISCORD_BOT_ENABLED", DISCORD_BOT_ENABLED)
	DISCORD_BOT_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
	DISCORD_ADMIN_USERNAMES = adminUsernamesFromEnv()
	DISCORD_COMMAND_IGNORE_CASE = getEnvOrDefaultBool("DISCORD_COMMAND_IGNORE_CASE", DISCORD_COMMAND_IGNORE_CASE)
	DISCORD_COMMAND_PREFIX = getEnvOrDefault("DISCORD_COMMAND_PREFIX", DISCORD_COMMAND_PREFIX)

	PALWORLD_RCON_ENABLED = getEnvOrDefaultBool("PALWORLD_RCON_ENABLED", PALWORLD_RCON_ENABLED)
	PALWORLD_RCON_HOST = getEnvOrDefault("PALWORLD_RCON_HOST", PALWORLD_RCON_HOST)
	PALWORLD_RCON_PORT = getEnvOrDefaultUint("PALWORLD_RCON_PORT", PALWORLD_RCON_PORT)
	PALWORLD_ADMIN_PASSWORD = os.Getenv("PALWORLD_ADMIN_PASSWORD")
	PALWORLD_SERVER_FILE_PATH = palworldServerPathFromEnv()
	PALWORLD_SERVER_EXECUTE_FLAGS = palworldServerExecuteFlagsFromEnv()

	DISCORD_DASHBOARD_CHANNEL_ID = os.Getenv("DISCORD_DASHBOARD_CHANNEL_ID")
	DISCORD_LOG_CHANNEL_ID = os.Getenv("DISCORD_LOG_CHANNEL_ID")
	DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID = os.Getenv("DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID")
	DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID = os.Getenv("DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID")
	DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID = os.Getenv("DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID")

	LANGUAGE = os.Getenv("LANGUAGE")

	return nil
}

func adminUsernamesFromEnv() []string {
	users := []string{}

	for _, user := range strings.Split(os.Getenv("DISCORD_ADMIN_USERNAMES"), ",") {
		users = append(users, strings.TrimSpace(user))
	}
	return users
}

func palworldServerPathFromEnv() string {
	path := os.Getenv("PALWORLD_SERVER_FILE_PATH")
	if path == "" {
		if runtime.GOOS == "windows" {
			return DEFAULT_PALWORLD_WINDOWS_SERVER_PATH
		}

		return DEFAULT_PALWORLD_LINUX_SERVER_PATH
	}

	return path
}

func palworldServerExecuteFlagsFromEnv() []string {
	flags := os.Getenv("PALWORLD_SERVER_EXECUTE_FLAGS")

	if flags == "" {
		return []string{}
	}

	return strings.Split(flags, " ")
}
