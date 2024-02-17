package commands

import (
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
)

func help() string {
	return i18n.Help(config.DISCORD_COMMAND_PREFIX)
}
