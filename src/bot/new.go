package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
)

func New(token string) (session *discordgo.Session) {
	discord, err := discordgo.New("Bot " + config.DISCORD_BOT_TOKEN)
	if err != nil {
		console_decoration.PrintError("error creating Discord session\n")
		panic(err)
	}

	return discord
}
