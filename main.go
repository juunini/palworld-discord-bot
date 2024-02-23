package main

import (
	"time"

	"github.com/juunini/palworld-discord-bot/src/bot"
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/web"
)

func init() {
	config.Load()
	i18n.SetLanguage(config.LANGUAGE)
}

func main() {
	if config.WEB_SERVER_ENABLED {
		go web.Listen(config.WEB_SERVER_PORT)
	}

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
