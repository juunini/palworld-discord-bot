package config

import (
	"fmt"
	"os"

	"github.com/juunini/palworld-discord-bot/src/console_decoration"
)

func askConfigToUser() {
	fmt.Printf("%s.env file not found. Please provide the following information:%s\n\n", console_decoration.RED, console_decoration.RESET)

	scanDiscordBotToken()
	scanDiscordAdminUsernames()
	scanPalworldRconHost()
	scanPalworldRconPort()
	scanPalworldRconPassword()

	env := fmt.Sprintf(
		`DISCORD_BOT_TOKEN=%s
DISCORD_ADMIN_USERNAMES=%s
PALWORLD_RCON_HOST=%s
PALWORLD_RCON_PORT=%s
PALWORLD_RCON_PASSWORD=%s`,
		DISCORD_BOT_TOKEN,
		DISCORD_ADMIN_USERNAMES_STRING,
		PALWORLD_RCON_HOST,
		PALWORLD_RCON_PORT,
		PALWORLD_RCON_PASSWORD,
	)
	fmt.Printf(
		"Created .env file with the following content:\n\n%s%s%s\n\n",
		console_decoration.BOLD,
		env,
		console_decoration.RESET,
	)

	os.WriteFile(".env", []byte(env), 0644)

	Init()
}
