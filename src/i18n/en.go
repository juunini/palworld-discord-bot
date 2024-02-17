package i18n

import (
	"fmt"

	"github.com/juunini/palworld-discord-bot/src/config"
)

func en() {
	Help = fmt.Sprintf(`%s help - Display help message.`, config.DISCORD_COMMAND_PREFIX)
	UnknownCommand = "Unknown command"
}
