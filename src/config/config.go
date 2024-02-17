package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
)

var DISCORD_BOT_TOKEN = ""

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		askConfigToUser()
		return
	}

	DISCORD_BOT_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
}

func askConfigToUser() {
	var discordBotToken string

	fmt.Printf("%s.env file not found. Please provide the following information:%s\n\n", console_decoration.RED, console_decoration.RESET)

	fmt.Printf("Discord Bot Token: %s", console_decoration.BOLD)
	fmt.Scanf("%s", &discordBotToken)
	fmt.Printf("%s", console_decoration.RESET)

	createdEnv := fmt.Sprintf(`DISCORD_BOT_TOKEN=%s`, discordBotToken)
	fmt.Printf("Created .env file with the following content:\n\n%s\n\n", createdEnv)

	os.WriteFile(".env", []byte(createdEnv), 0644)

	Init()
}
