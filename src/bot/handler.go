package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/juunini/palworld-discord-bot/src/bot/commands"
	"github.com/juunini/palworld-discord-bot/src/config"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isBotMessage(s, m) {
		return
	}

	if isCommand(m) {
		s.ChannelMessageSend(m.ChannelID, commands.Response(m.Content))
		return
	}
}

func isBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Author.ID == s.State.User.ID
}

func isCommand(m *discordgo.MessageCreate) bool {
	return strings.HasPrefix(m.Content, config.DISCORD_COMMAND_PREFIX)
}
