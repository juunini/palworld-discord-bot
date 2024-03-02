package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/juunini/palworld-discord-bot/src/commands"
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/palworld/settings"
	"github.com/juunini/palworld-discord-bot/src/utils"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isBotMessage(s, m) {
		return
	}

	if isCommand(m) {
		isAdmin := utils.IsAdmin(m.Author.Username)

		if !isAdmin {
			return
		}

		if strings.HasPrefix(m.Content, config.DISCORD_COMMAND_PREFIX+" "+config.DISCORD_COMMAND_ALIAS_SERVER_SETTINGS) {
			content, _ := strings.CutPrefix(m.Content, config.DISCORD_COMMAND_PREFIX+" ")
			settings.Response(s, m.ChannelID, content)
			return
		}

		s.ChannelMessageSend(m.ChannelID, commands.Response(m.Content, m.Author.Username))
		return
	}
}

func isBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Author.ID == s.State.User.ID
}

func isCommand(m *discordgo.MessageCreate) bool {
	return strings.HasPrefix(m.Content, config.DISCORD_COMMAND_PREFIX)
}
