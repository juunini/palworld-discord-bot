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
	FailedToBroadcastCommand = "Failed to Broadcast command"
	FailedToKickCommand = "Failed to Kick command"
	FailedToBanCommand = "Failed to Ban command"
	WrongParameters = "Wrong parameters"
	Help = func(commandPrefix string, isAdmin bool) string {
		message := fmt.Sprintf("`%s help` - Show help\n", commandPrefix)

		if !isAdmin {
			return message
		}

		message += fmt.Sprintf("`%s kick <SteamID>` - Kick <SteamID> from server. Can connect again.\n", commandPrefix)
		message += fmt.Sprintf("`%s ban <SteamID>` - Ban <SteamID> from server. Can't connect again.\n", commandPrefix)
		message += fmt.Sprintf("`%s broadcast <message>` - Send <message> to all users as SYSTEM message.\n", commandPrefix)
		message += fmt.Sprintf("`%s shutdown <seconds> <message>` - Shutdown server after <seconds> with <message>.\n", commandPrefix)
		message += fmt.Sprintf("`%s doExit` - Force exit server.\n", commandPrefix)
		message += fmt.Sprintf("`%s save` - Save server.", commandPrefix)
		return message
	}
	UnknownCommand = "Unknown command"
}
