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

	if config.DISCORD_DASHBOARD_CHANNEL_ID != "" {
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
		content += fmt.Sprintf("- %s\n", player.Username+" ("+player.SteamID+")")
	}

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Online Players (%d)", len(watchers.OnlinePlayers)),
		Description: "0.1.4.1 기준, 영어가 아닌 문자로 닉네임이 설정된 유저가 있을 시 SteamID등의 정보가 제대로 표시되지 않을 수 있고, kick, ban 등의 명령어가 제대로 작동하지 않을 수도 있습니다.",
		Fields: []*discordgo.MessageEmbedField{
			{
				Value:  content,
				Inline: false,
			},
		},
	}

	if config.DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID == "" {
		message, err := session.ChannelMessageSendEmbed(config.DISCORD_DASHBOARD_CHANNEL_ID, embed)
		if err != nil {
			return
		}

		config.DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID = message.ID
		config.Save()
		return
	}

	_, err := session.ChannelMessageEditEmbed(config.DISCORD_DASHBOARD_CHANNEL_ID, config.DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID, embed)
	if err != nil {
		fmt.Println(err)
		return
	}

}
