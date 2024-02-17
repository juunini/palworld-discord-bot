package config

import (
	"fmt"

	"github.com/juunini/palworld-discord-bot/src/console_decoration"
)

func scanDiscordBotToken() {
	fmt.Printf("Discord Bot Token: %s", console_decoration.BOLD)
	fmt.Scanf("%s", &DISCORD_BOT_TOKEN)
	fmt.Print(console_decoration.RESET)
}

func scanDiscordAdminUsernames() {
	fmt.Printf("Discord Admin Usernames (comma separated): %s", console_decoration.BOLD)
	fmt.Scanf("%s", &DISCORD_ADMIN_USERNAMES_STRING)
	fmt.Print(console_decoration.RESET)
}

func scanDiscordCommandPrefix() {
	fmt.Printf("Discord Command Prefix(default: %s): %s", DISCORD_COMMAND_PREFIX, console_decoration.BOLD)

	var prefix string
	fmt.Scanf("%s", &prefix)
	fmt.Print(console_decoration.RESET)

	if prefix == "" {
		return
	}

	DISCORD_COMMAND_PREFIX = prefix
}

func scanPalworldRconHost() {
	fmt.Printf("Palworld RCON Host: %s", console_decoration.BOLD)
	fmt.Scanf("%s", &PALWORLD_RCON_HOST)
	fmt.Print(console_decoration.RESET)
}

func scanPalworldRconPort() {
	fmt.Printf("Palworld RCON Port: %s", console_decoration.BOLD)
	fmt.Scanf("%s", &PALWORLD_RCON_PORT)
	fmt.Print(console_decoration.RESET)
}

func scanPalworldRconPassword() {
	fmt.Printf("Palworld RCON Password: %s", console_decoration.BOLD)
	fmt.Scanf("%s", &PALWORLD_RCON_PASSWORD)
	fmt.Print(console_decoration.RESET)
}
