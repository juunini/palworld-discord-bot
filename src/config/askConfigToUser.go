package config

import (
	"fmt"
	"os"

	"github.com/juunini/palworld-discord-bot/src/console_decoration"
)

func askConfigToUser() {
	fmt.Printf("%s.env file not found. Please provide the following information:%s\n\n", console_decoration.RED, console_decoration.RESET)

	scanDiscordBotToken()
	scanDiscordAdminUsers()

	createdEnv := fmt.Sprintf(`DISCORD_BOT_TOKEN=%s
DISCORD_ADMIN_USERS=%s`,
		DISCORD_BOT_TOKEN,
		DISCORD_ADMIN_USERS_STRING)
	fmt.Printf("Created .env file with the following content:\n\n%s\n\n", createdEnv)

	os.WriteFile(".env", []byte(createdEnv), 0644)

	Init()
}

func scanDiscordBotToken() {
	fmt.Printf("Discord Bot Token: %s", console_decoration.BOLD)
	fmt.Scanf("%s", &DISCORD_BOT_TOKEN)
	fmt.Print(console_decoration.RESET)
}

func scanDiscordAdminUsers() {
	fmt.Printf("Discord Admin Users (comma separated): %s", console_decoration.BOLD)
	fmt.Scanf("%s", &DISCORD_ADMIN_USERS_STRING)
	fmt.Print(console_decoration.RESET)
}
