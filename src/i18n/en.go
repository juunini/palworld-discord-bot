package i18n

import (
	"fmt"

	"github.com/juunini/palworld-discord-bot/src/config"
)

func en() {
	BotRunningStart = "Bot is now running. Press CTRL-C to exit."
	FailedToConnectRconServer = "Failed to connect RCON server"
	FailedToSaveCommand = "Failed to Save command"
	FailedToDoExitCommand = "Failed to DoExit command"
	Help = fmt.Sprintf(`%s help - Display help message.`, config.DISCORD_COMMAND_PREFIX)
	UnknownCommand = "Unknown command"
}
