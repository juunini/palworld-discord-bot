package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
)

func Connect(session *discordgo.Session) {
	if err := session.Open(); err != nil {
		console_decoration.PrintError("error opening connection")
		panic(err)
	}
}
