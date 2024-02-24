package i18n

var Languages = []map[string]string{
	{"code": "ko", "name": "한국어"},
	{"code": "en", "name": "English"},
}

func SetLanguage(language string) {
	en()

	switch language {
	case "ko":
		ko()
		return
	}
}

var (
	BotRunningStart             string
	FailedToConnectRconServer   string
	FailedToSaveCommand         string
	FailedToDoExitCommand       string
	FailedToShutdownCommand     string
	FailedToBroadcastCommand    string
	FailedToKickCommand         string
	FailedToBanCommand          string
	FailedToShowPlayerCommand   string
	FailedToStartServerCommand  string
	SuccessToStartServerCommand string
	WrongParameters             string
	UnknownCommand              string
	Help                        func(commandPrefix string, isAdmin bool) string
)

var (
	WebServerOpeningMessage string

	WebConfig                    string
	EnableWebServer              string
	EnableWebServerDisablePrompt string
	WebServerPort                string
	WebServerPortTooltip         string

	Language       string
	SelectLanguage string

	DiscordBotConfig                   string
	EnableDiscordBot                   string
	DiscordBotToken                    string
	DiscordBotTokenTooltip             string
	DiscordAdminUsernames              string
	DiscordAdminUsernamesTooltip       string
	DiscordCommandCaseSensitive        string
	DiscordCommandCaseSensitiveTooltip string
	DiscordCommandPrefix               string
	DiscordCommandPrefixTooltip        string

	PalworldConfig                    string
	EnablePalworldRcon                string
	EnablePalworldRconTooltip         string
	PalworldRconHost                  string
	PalworldRconHostTooltip           string
	PalworldRconPort                  string
	PalworldRconPortTooltip           string
	PalworldAdminPassword             string
	PalworldAdminPasswordTooltip      string
	PalworldServerFilePath            string
	PalworldServerFilePathTooltip     string
	PalworldServerExecuteFlags        string
	PalworldServerExecuteFlagsTooltip string

	DiscordChannelConfig                           string
	DiscordDashboardChannelID                      string
	DiscordDashboardChannelIDTooltip               string
	DiscordLogChannelID                            string
	DiscordLogChannelIDTooltip                     string
	DiscordDashboardOnlinePlayersMessageID         string
	DiscordDashboardOnlinePlayersMessageIDTooltip  string
	DiscordDashboardPalworldConfigMessageID        string
	DiscordDashboardPalworldConfigMessageIDTooltip string
	DiscordDashboardBotConfigMessageID             string
	DiscordDashboardBotConfigMessageIDTooltip      string

	DiscordCommandAliases                 string
	DiscordCommandAliasKickTooltip        string
	DiscordCommandAliasBanTooltip         string
	DiscordCommandAliasBroadcastTooltip   string
	DiscordCommandAliasShutdownTooltip    string
	DiscordCommandAliasDoExitTooltip      string
	DiscordCommandAliasSaveTooltip        string
	DiscordCommandAliasStartServerTooltip string

	Close   string
	Confirm string
	Save    string
)
