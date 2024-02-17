package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		askConfigToUser()
		return
	}

	loadEnv()
}

func trimSpaceAdminUsers() (users []string) {
	users = []string{}

	for _, user := range strings.Split(DISCORD_ADMIN_USERNAMES_STRING, ",") {
		users = append(users, strings.TrimSpace(user))
	}
	return
}

func loadEnv() {
	DISCORD_BOT_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
	DISCORD_ADMIN_USERNAMES_STRING = os.Getenv("DISCORD_ADMIN_USERS")
	DISCORD_ADMIN_USERNAMES = trimSpaceAdminUsers()
	DISCORD_COMMAND_PREFIX = os.Getenv("DISCORD_COMMAND_PREFIX")
	PALWORLD_RCON_HOST = os.Getenv("PALWORLD_RCON_HOST")
	PALWORLD_RCON_PORT = os.Getenv("PALWORLD_RCON_PORT")
	PALWORLD_RCON_PASSWORD = os.Getenv("PALWORLD_RCON_PASSWORD")
}
