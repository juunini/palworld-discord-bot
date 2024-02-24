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
	FailedToShowPlayerCommand = "Failed to ShowPlayer command"
	FailedToStartServerCommand = "Failed to start server"
	SuccessToStartServerCommand = "Server started"
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
		message += fmt.Sprintf("`%s startServer` - Start server.", commandPrefix)
		return message
	}
	UnknownCommand = "Unknown command"

	WebConfig = "Web Config"
	EnableWebServer = "Enable Web Server"
	EnableWebServerDisablePrompt = "Do you want to disable web server? You need to modify .env file to enable it again."
	WebServerPort = "Web Server Port"
	WebServerPortTooltip = "Set the port of web server. It is recommended to use a number above 60000."

	Language = "Language"
	SelectLanguage = "Select Language"

	DiscordBotConfig = "Discord Bot Config"
	EnableDiscordBot = "Enable Discord Bot"
	DiscordBotToken = "Discord Bot Token"
	DiscordBotTokenTooltip = "Enter the token of the Discord bot. You can find out how to get the token by clicking the ? button you are currently hovering over."
	DiscordAdminUsernames = "Discord Admin Usernames"
	DiscordAdminUsernamesTooltip = "Enter the nicknames of Discord admins. If there are multiple, separate them with a comma."
	DiscordCommandCaseSensitive = "Discord Command Case Sensitive"
	DiscordCommandCaseSensitiveTooltip = "Set whether to distinguish the case of Discord commands."
	DiscordCommandPrefix = "Discord Bot Call Command"
	DiscordCommandPrefixTooltip = "Set the command to call the Discord bot."

	PalworldConfig = "Palworld Config"
	EnablePalworldRcon = "Enable Palworld RCON (Please set RCONEnable value in PalWorldSettings.ini to true in advance.)"
	PalworldRconHost = "Palworld RCON Host"
	PalworldRconHostTooltip = "Enter the address of the server where the Palworld server is running."
	PalworldRconPort = "Palworld RCON Port"
	PalworldRconPortTooltip = "Enter the port of the server where the Palworld server is running."
	PalworldAdminPassword = "Palworld Admin Password"
	PalworldAdminPasswordTooltip = "Enter the password of the Palworld server."
	PalworldServerFilePath = "Palworld Server File Path"
	PalworldServerFilePathTooltip = "Enter the path of the Palworld server file."
	PalworldServerExecuteFlags = "Palworld Server Execute Flags"
	PalworldServerExecuteFlagsTooltip = "Enter the flags of the Palworld server."

	Close = "Close"
	Confirm = "Confirm"
	Save = "Save"
}
