package bot

import "github.com/bwmarrin/discordgo"

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isBotMessage(s, m) {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func isBotMessage(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return m.Author.ID == s.State.User.ID
}
