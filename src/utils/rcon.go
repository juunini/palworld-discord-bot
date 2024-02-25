package utils

import (
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	palworldrcon "github.com/juunini/palworld-rcon"
)

func RconClient() (palworldrcon.Client, error) {
	client := palworldrcon.Client{
		Host:          config.PALWORLD_RCON_HOST,
		Port:          config.PALWORLD_RCON_PORT,
		AdminPassword: config.PALWORLD_ADMIN_PASSWORD,
	}

	err := client.Connect()
	if err != nil {
		console_decoration.PrintError(i18n.FailedToConnectRconServer + ": " + err.Error())
	}

	return client, err
}
