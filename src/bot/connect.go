package bot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/web"
)

func Connect(session *discordgo.Session) {
	if err := session.Open(); err != nil {
		console_decoration.PrintError("error opening connection")
		panic(err)
	}

	console_decoration.PrintSuccess(i18n.BotRunningStart)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	session.Close()
	// FIXME: move upper code and under line code to main.go
	web.Shutdown()
}
