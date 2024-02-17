package config

import (
	"fmt"
	"os"

	"github.com/juunini/palworld-discord-bot/src/console_decoration"
)

func askConfigToUser() {
	fmt.Printf("%s.env file not found. Please provide the following information:%s\n\n", console_decoration.RED, console_decoration.RESET)

	scanEnv("Discord Bot Token", &DISCORD_BOT_TOKEN)
	scanEnv("Discord Admin Usernames (comma separated)", &DISCORD_ADMIN_USERNAMES_STRING)
	scanEnvWithDefaultValue("Discord Command Prefix", &DISCORD_COMMAND_PREFIX, DISCORD_COMMAND_PREFIX)
	scanEnvWithDefaultValue("Palworld RCON Host", &PALWORLD_RCON_HOST, PALWORLD_RCON_HOST)
	scanEnvWithDefaultValue("Palworld RCON Port", &PALWORLD_RCON_PORT, PALWORLD_RCON_PORT)
	scanEnv("Palworld RCON Password", &PALWORLD_RCON_PASSWORD)

	env := envString()
	fmt.Printf(
		"Created .env file with the following content:\n\n%s%s%s\n\n",
		console_decoration.BOLD,
		env,
		console_decoration.RESET,
	)

	os.WriteFile(".env", []byte(env), 0644)

	Init()
}
