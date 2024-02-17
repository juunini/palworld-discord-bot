package i18n

import (
	"fmt"
)

func en() {
	BotRunningStart = "Bot is now running. Press CTRL-C to exit."
	FailedToConnectRconServer = "Failed to connect RCON server"
	FailedToSaveCommand = "Failed to Save command"
	FailedToDoExitCommand = "Failed to DoExit command"
	FailedToShutdownCommand = "Failed to Shutdown command"
	WrongParameters = "Wrong parameters"
	Help = func(commandPrefix string) string {
		return fmt.Sprintf(`%s help - Show help`, commandPrefix)
	}
	UnknownCommand = "Unknown command"
}
