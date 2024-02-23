package main

import (
	"time"

	"github.com/juunini/palworld-discord-bot/src/bot"
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
)

func init() {
	config.Load()
	i18n.SetLanguage(config.LANGUAGE)
}

func main() {
	session := bot.New(config.DISCORD_BOT_TOKEN)

	session.AddHandler(bot.Handler)

	go func() {
		for {
			bot.Watch(session)
			time.Sleep(10 * time.Second)
		}
	}()

	bot.Connect(session)
}
