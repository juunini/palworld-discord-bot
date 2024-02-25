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
	Help = func(params HelpParams, isAdmin bool) string {
		message := fmt.Sprintf("`%s %s` - Show help\n", params.CommandPrefix, params.HelpAlias)

		if !isAdmin {
			return message
		}

		message += fmt.Sprintf("`%s %s <SteamID>` - Kick <SteamID> from server. Can connect again.\n", params.CommandPrefix, params.KickAlias)
		message += fmt.Sprintf("`%s %s <SteamID>` - Ban <SteamID> from server. Can't connect again.\n", params.CommandPrefix, params.BanAlias)
		message += fmt.Sprintf("`%s %s <message>` - Send <message> to all users as SYSTEM message.\n", params.CommandPrefix, params.BroadcastAlias)
		message += fmt.Sprintf("`%s %s <seconds> <message>` - Shutdown server after <seconds> with <message>.\n", params.CommandPrefix, params.ShutdownAlias)
		message += fmt.Sprintf("`%s %s` - Force exit server.\n", params.CommandPrefix, params.DoExitAlias)
		message += fmt.Sprintf("`%s %s` - Save server.", params.CommandPrefix, params.SaveAlias)
		message += fmt.Sprintf("`%s %s` - Start server.", params.CommandPrefix, params.StartServerAlias)
		return message
	}
	UnknownCommand = "Unknown command"

	WebServerOpeningMessage = "If you cannot access the web server, please modify the .env file and run it again."
	WebServerShutdownMessage = "Web server has been shut down."

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
	DiscordBotTokenDescription = "Enter the token of the Discord bot. You can find out how to get the token at https://github.com/juunini/palworld-discord-bot/wiki/How-to-create-Discord-Bot%3F-%5BEnglish%5D"
	DiscordAdminUsernames = "Discord Admin Usernames"
	DiscordAdminUsernamesTooltip = "Enter the 'Discord username' of the administrator. If there are multiple, separate them with a comma."
	DiscordCommandCaseSensitive = "Discord Command Case Sensitive"
	DiscordCommandCaseSensitiveTooltip = "Set whether to distinguish the case of Discord commands."
	DiscordCommandPrefix = "Discord Bot Call Command"
	DiscordCommandPrefixTooltip = "Set the command to call the Discord bot."

	PalworldConfig = "Palworld Config"
	EnablePalworldRcon = "Enable Palworld RCON"
	EnablePalworldRconTooltip = "Enable the Palworld RCON function."
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

	DiscordChannelConfig = "Discord Channel Config"
	DiscordDashboardChannelID = "Discord Dashboard Channel ID"
	DiscordDashboardChannelIDTooltip = "Enter the channel ID for using the dashboard function. If not entered, the function is not used. You can find out how to check the channel ID by clicking the ? button you are currently hovering over."
	DiscordDashboardChannelIDDescription = "Enter the channel ID for using the dashboard function. If not entered, the function is not used. You can find out how to check the channel ID at https://github.com/juunini/palworld-discord-bot/wiki/How-to-get-Channel-ID%3F-%5BEnglish%5D"
	DiscordLogChannelID = "Discord Log Channel ID"
	DiscordLogChannelIDTooltip = "Enter the channel ID to check the user's connection/disconnection record. If not entered, the function is not used. You can find out how to check the channel ID by clicking the ? button you are currently hovering over."
	DiscordLogChannelIDDescription = "Enter the channel ID to check the user's connection/disconnection record. If not entered, the function is not used. You can find out how to check the channel ID at https://github.com/juunini/palworld-discord-bot/wiki/How-to-get-Channel-ID%3F-%5BEnglish%5D"
	DiscordDashboardOnlinePlayersMessageID = "Discord Dashboard Online Players Message ID"
	DiscordDashboardOnlinePlayersMessageIDTooltip = "Automatically set value. If not necessary, do not modify."
	DiscordDashboardPalworldConfigMessageID = "Discord Dashboard Palworld Config Message ID"
	DiscordDashboardPalworldConfigMessageIDTooltip = "Automatically set value. If not necessary, do not modify."
	DiscordDashboardBotConfigMessageID = "Discord Dashboard Bot Config Message ID"
	DiscordDashboardBotConfigMessageIDTooltip = "Automatically set value. If not necessary, do not modify."

	DiscordCommandAliases = "Discord Command Aliases"
	DiscordCommandAliasHelpTooltip = "Help command customizing"
	DiscordCommandAliasKickTooltip = "Kick command customizing"
	DiscordCommandAliasBanTooltip = "Ban command customizing"
	DiscordCommandAliasBroadcastTooltip = "Broadcast all users command customizing"
	DiscordCommandAliasShutdownTooltip = "Shutdown after few seconds command customizing"
	DiscordCommandAliasDoExitTooltip = "Server force shutdown command customizing"
	DiscordCommandAliasSaveTooltip = "Save command customizing"
	DiscordCommandAliasStartServerTooltip = "StartServer command customizing"

	Close = "Close"
	Confirm = "Confirm"
	Save = "Save"

	ConfigUpdateSuccess = "Config update success"
	ConfigUpdateFailed = "Config update failed"

	OnlinePlayers = "Online Players"
	NoticeNonEnglishNickname = "As of version 0.1.4.1, if a user has a nickname set in a language other than English, their SteamID information may not be displayed correctly, and commands such as kick and ban may not work properly."
}
