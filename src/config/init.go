package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		askConfigToUser()
		return
	}

	loadEnv()
}

func askConfigToUser() {
	fmt.Printf("%s.env file not found. Please provide the following information:%s\n\n", console_decoration.RED, console_decoration.RESET)

	scanEnv("Discord Bot Token", &DISCORD_BOT_TOKEN)
	scanEnv("Discord Admin Usernames (comma separated)", &DISCORD_ADMIN_USERNAMES_STRING)
	scanEnvWithDefaultValue("Discord Command Prefix", &DISCORD_COMMAND_PREFIX, DISCORD_COMMAND_PREFIX)
	scanEnvWithDefaultValue("Palworld RCON Host", &PALWORLD_RCON_HOST, PALWORLD_RCON_HOST)
	scanEnvWithDefaultValue("Palworld RCON Port", &PALWORLD_RCON_PORT, PALWORLD_RCON_PORT)
	scanEnv("Palworld RCON Password", &PALWORLD_RCON_PASSWORD)

	saveEnv()
	printEnv()
	Init()
}

func saveEnv() {
	godotenv.Write(envMap(), ".env")
}

func printEnv() {
	fmt.Printf("\n%sCreated .env file with the following content:%s\n\n", console_decoration.GREEN, console_decoration.RESET)

	for key, value := range envMap() {
		fmt.Printf("%s=%s%s%s\n", key, console_decoration.BOLD, value, console_decoration.RESET)
	}

	fmt.Print("\n")
}

func loadEnv() {
	DISCORD_BOT_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
	DISCORD_ADMIN_USERNAMES_STRING = os.Getenv("DISCORD_ADMIN_USERNAMES")
	DISCORD_ADMIN_USERNAMES = trimSpaceAdminUsers()
	DISCORD_COMMAND_PREFIX = os.Getenv("DISCORD_COMMAND_PREFIX")
	PALWORLD_RCON_HOST = os.Getenv("PALWORLD_RCON_HOST")
	PALWORLD_RCON_PORT = os.Getenv("PALWORLD_RCON_PORT")
	PALWORLD_RCON_PASSWORD = os.Getenv("PALWORLD_RCON_PASSWORD")
}

func envMap() map[string]string {
	return map[string]string{
		"DISCORD_BOT_TOKEN":       DISCORD_BOT_TOKEN,
		"DISCORD_ADMIN_USERNAMES": DISCORD_ADMIN_USERNAMES_STRING,
		"DISCORD_COMMAND_PREFIX":  DISCORD_COMMAND_PREFIX,
		"PALWORLD_RCON_HOST":      PALWORLD_RCON_HOST,
		"PALWORLD_RCON_PORT":      PALWORLD_RCON_PORT,
		"PALWORLD_RCON_PASSWORD":  PALWORLD_RCON_PASSWORD,
	}
}

func trimSpaceAdminUsers() []string {
	users := []string{}

	for _, user := range strings.Split(DISCORD_ADMIN_USERNAMES_STRING, ",") {
		users = append(users, strings.TrimSpace(user))
	}
	return users
}
