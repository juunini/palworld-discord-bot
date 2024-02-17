package main

import (
	"github.com/juunini/palworld-discord-bot/src/bot"
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
)

func init() {
	config.Init()
	i18n.SetLanguage(config.LANGUAGE)
}

func main() {
	session := bot.New(config.DISCORD_BOT_TOKEN)

	session.AddHandler(bot.Handler)

	bot.Connect(session)
}
