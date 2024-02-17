package config

import (
	"fmt"
	"os"
	"strconv"
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
	scanEnvWithDefaultValue("Palworld RCON Port", &PALWORLD_RCON_PORT_STRING, PALWORLD_RCON_PORT_STRING)
	scanEnv("Palworld RCON Password", &PALWORLD_RCON_PASSWORD)
	scanEnvWithDefaultValue("Language", &LANGUAGE, LANGUAGE)

	saveEnv()
	printEnv()
	loadEnv()
}

func saveEnv() {
	if err := godotenv.Write(envMap(), ".env"); err != nil {
		fmt.Printf("%sError writing .env file: %s%s\n", console_decoration.RED, err, console_decoration.RESET)
		os.Exit(1)
	}
}

func printEnv() {
	fmt.Printf("\n%sCreated .env file with the following content:%s\n\n", console_decoration.GREEN, console_decoration.RESET)

	for key, value := range envMap() {
		fmt.Printf("%s=%s%s%s\n", key, console_decoration.BOLD, value, console_decoration.RESET)
	}

	fmt.Print("\n")
}

func loadEnv() {
	godotenv.Load(".env")

	DISCORD_BOT_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
	DISCORD_ADMIN_USERNAMES_STRING = os.Getenv("DISCORD_ADMIN_USERNAMES")
	DISCORD_ADMIN_USERNAMES = trimSpaceAdminUsers()
	DISCORD_COMMAND_PREFIX = os.Getenv("DISCORD_COMMAND_PREFIX")
	PALWORLD_RCON_HOST = os.Getenv("PALWORLD_RCON_HOST")
	PALWORLD_RCON_PORT_STRING = os.Getenv("PALWORLD_RCON_PORT")

	var err error
	PALWORLD_RCON_PORT, err = strconv.ParseUint(PALWORLD_RCON_PORT_STRING, 10, 64)
	if err != nil {
		fmt.Printf("%sError parsing PALWORLD_RCON_PORT: %s%s\n", console_decoration.RED, err, console_decoration.RESET)
		os.Exit(1)
	}

	PALWORLD_RCON_PASSWORD = os.Getenv("PALWORLD_RCON_PASSWORD")
	DISCORD_DASHBOARD_CHANNEL_ID = os.Getenv("DISCORD_DASHBOARD_CHANNEL_ID")
	DISCORD_LOG_CHANNEL_ID = os.Getenv("DISCORD_LOG_CHANNEL_ID")
	LANGUAGE = os.Getenv("LANGUAGE")
}

func envMap() map[string]string {
	return map[string]string{
		"DISCORD_BOT_TOKEN":            DISCORD_BOT_TOKEN,
		"DISCORD_ADMIN_USERNAMES":      DISCORD_ADMIN_USERNAMES_STRING,
		"DISCORD_COMMAND_PREFIX":       DISCORD_COMMAND_PREFIX,
		"PALWORLD_RCON_HOST":           PALWORLD_RCON_HOST,
		"PALWORLD_RCON_PORT":           PALWORLD_RCON_PORT_STRING,
		"PALWORLD_RCON_PASSWORD":       PALWORLD_RCON_PASSWORD,
		"DISCORD_DASHBOARD_CHANNEL_ID": DISCORD_DASHBOARD_CHANNEL_ID,
		"DISCORD_LOG_CHANNEL_ID":       DISCORD_LOG_CHANNEL_ID,
		"LANGUAGE":                     LANGUAGE,
	}
}

func trimSpaceAdminUsers() []string {
	users := []string{}

	for _, user := range strings.Split(DISCORD_ADMIN_USERNAMES_STRING, ",") {
		users = append(users, strings.TrimSpace(user))
	}
	return users
}
