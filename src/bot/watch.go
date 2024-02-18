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

	if config.DISCORD_LOG_CHANNEL_ID != "" {
		sendNewPlayers(session)
		sendLeftPlayers(session)
	}

	if config.DISCORD_DASHBOARD_CHANNEL_ID == "" {
		sendOnlinePlayerDashboard(session)
	}
}

func sendNewPlayers(session *discordgo.Session) {
	for _, player := range watchers.NewPlayers {
		session.ChannelMessageSend(config.DISCORD_LOG_CHANNEL_ID, player.Username+" joined the server")
	}
}

func sendLeftPlayers(session *discordgo.Session) {
	for _, player := range watchers.LeftPlayers {
		session.ChannelMessageSend(config.DISCORD_LOG_CHANNEL_ID, player.Username+" left the server")
	}
}

func sendOnlinePlayerDashboard(session *discordgo.Session) {
	if len(watchers.NewPlayers) == 0 && len(watchers.LeftPlayers) == 0 {
		return
	}

	content := ""

	for _, player := range watchers.OnlinePlayers {
		content += fmt.Sprintf("- %s\n", player.Username)
	}

	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Online Players (%d)", len(watchers.OnlinePlayers)),
		Fields: []*discordgo.MessageEmbedField{
			{
				Value:  content,
				Inline: false,
			},
		},
	}

	if config.DISCORD_DASHBAORD_ONLINE_PLAYERS_MESSAGE_ID == "" {
		message, err := session.ChannelMessageSendEmbed(config.DISCORD_DASHBOARD_CHANNEL_ID, embed)
		if err != nil {
			return
		}

		config.DISCORD_DASHBAORD_ONLINE_PLAYERS_MESSAGE_ID = message.ID
		config.SaveEnv()
		return
	}

	_, err := session.ChannelMessageEditEmbed(config.DISCORD_DASHBOARD_CHANNEL_ID, config.DISCORD_DASHBAORD_ONLINE_PLAYERS_MESSAGE_ID, embed)
	if err != nil {
		fmt.Println(err)
		return
	}

}
