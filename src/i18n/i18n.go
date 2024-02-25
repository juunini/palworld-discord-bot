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

type HelpParams struct {
	CommandPrefix    string
	HelpAlias        string
	KickAlias        string
	BanAlias         string
	BroadcastAlias   string
	ShutdownAlias    string
	DoExitAlias      string
	SaveAlias        string
	StartServerAlias string
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
	Help                        func(params HelpParams, isAdmin bool) string

	WebServerOpeningMessage  string
	WebServerShutdownMessage string

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
	DiscordBotTokenDescription         string
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
	DiscordDashboardChannelIDDescription           string
	DiscordLogChannelID                            string
	DiscordLogChannelIDTooltip                     string
	DiscordLogChannelIDDescription                 string
	DiscordDashboardOnlinePlayersMessageID         string
	DiscordDashboardOnlinePlayersMessageIDTooltip  string
	DiscordDashboardPalworldConfigMessageID        string
	DiscordDashboardPalworldConfigMessageIDTooltip string
	DiscordDashboardBotConfigMessageID             string
	DiscordDashboardBotConfigMessageIDTooltip      string

	DiscordCommandAliases                 string
	DiscordCommandAliasHelpTooltip        string
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

	ConfigUpdateSuccess string
	ConfigUpdateFailed  string

	OnlinePlayers            string
	NoticeNonEnglishNickname string
)
