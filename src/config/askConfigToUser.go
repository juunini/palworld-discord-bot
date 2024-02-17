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
	scanDiscordCommandPrefix()
	scanPalworldRconHost()
	scanPalworldRconPort()
	scanPalworldRconPassword()

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
