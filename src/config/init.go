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

	DISCORD_BOT_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
	DISCORD_ADMIN_USERS_STRING = os.Getenv("DISCORD_ADMIN_USERS")
	DISCORD_ADMIN_USERS = trimSpaceAdminUsers()
}

func trimSpaceAdminUsers() (users []string) {
	users = []string{}

	for _, user := range strings.Split(DISCORD_ADMIN_USERS_STRING, ",") {
		users = append(users, strings.TrimSpace(user))
	}
	return
}
