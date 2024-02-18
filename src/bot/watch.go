package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/juunini/palworld-discord-bot/src/bot/watchers"
	"github.com/juunini/palworld-discord-bot/src/config"
)

func Watch(session *discordgo.Session) {
	if err := watchers.Players(); err != nil {
		return
	}

	sendNewPlayers(session)
	sendLeftPlayers(session)
}

func sendNewPlayers(session *discordgo.Session) {
	for _, player := range watchers.NewPlayers {
		fmt.Println(player.Username)
		session.ChannelMessageSend(config.DISCORD_LOG_CHANNEL_ID, player.Username+" joined the server")
	}
}

func sendLeftPlayers(session *discordgo.Session) {
	for _, player := range watchers.LeftPlayers {
		session.ChannelMessageSend(config.DISCORD_LOG_CHANNEL_ID, player.Username+" left the server")
	}
}
