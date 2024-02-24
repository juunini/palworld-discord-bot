package web

import (
	"context"
	"log"

	"github.com/juunini/palworld-discord-bot/src/i18n"
)

func Shutdown() error {
	if app == nil {
		return nil
	}

	log.Println(i18n.WebServerShutdownMessage)
	return app.ShutdownWithContext(context.Background())
}
