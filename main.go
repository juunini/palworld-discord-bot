package main

import (
	"github.com/juunini/palworld-discord-bot/src/bot"
	"github.com/juunini/palworld-discord-bot/src/config"
)

func init() {
	config.Init()
}

func main() {
	session := bot.New(config.DISCORD_BOT_TOKEN)

	session.AddHandler(bot.Handler)

	bot.Connect(session)
}
