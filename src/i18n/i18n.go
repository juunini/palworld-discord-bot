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

	PalworldConfig                        string
	EnablePalworldRcon                    string
	EnablePalworldRconTooltip             string
	PalworldRconHost                      string
	PalworldRconHostTooltip               string
	PalworldRconPort                      string
	PalworldRconPortTooltip               string
	PalworldAdminPassword                 string
	PalworldAdminPasswordTooltip          string
	PalworldServerFilePath                string
	PalworldServerFilePathTooltip         string
	PalworldServerExecuteFlags            string
	PalworldServerExecuteFlagsTooltip     string
	PalworldServerSettingsFilePath        string
	PalworldServerSettingsFilePathTooltip string

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

	DiscordCommandAliases                    string
	DiscordCommandAliasHelpTooltip           string
	DiscordCommandAliasKickTooltip           string
	DiscordCommandAliasBanTooltip            string
	DiscordCommandAliasBroadcastTooltip      string
	DiscordCommandAliasShutdownTooltip       string
	DiscordCommandAliasDoExitTooltip         string
	DiscordCommandAliasSaveTooltip           string
	DiscordCommandAliasStartServerTooltip    string
	DiscordCommandAliasServerSettingsTooltip string

	Close   string
	Confirm string
	Save    string

	ConfigUpdateSuccess string
	ConfigUpdateFailed  string

	OnlinePlayers            string
	NoticeNonEnglishNickname string
)

var (
	Difficulty                                    string
	DifficultyDescription                         string
	DayTimeSpeedRate                              string
	NightTimeSpeedRate                            string
	ExpRate                                       string
	PalCaptureRate                                string
	PalSpawnNumRate                               string
	PalDamageRateAttack                           string
	PalDamageRateDefense                          string
	PlayerDamageRateAttack                        string
	PlayerDamageRateDefense                       string
	PlayerStomachDecreaceRate                     string
	PlayerStaminaDecreaceRate                     string
	PlayerAutoHPRegeneRate                        string
	PlayerAutoHpRegeneRateInSleep                 string
	PalStomachDecreaceRate                        string
	PalStaminaDecreaceRate                        string
	PalAutoHPRegeneRate                           string
	PalAutoHpRegeneRateInSleep                    string
	BuildObjectDamageRate                         string
	BuildObjectDeteriorationDamageRate            string
	BuildObjectDeteriorationDamageRateDescription string
	CollectionDropRate                            string
	CollectionDropRateDescription                 string
	CollectionObjectHpRate                        string
	CollectionObjectHpRateDescription             string
	CollectionObjectRespawnSpeedRate              string
	EnemyDropItemRate                             string
	DeathPenalty                                  string
	EnablePlayerToPlayerDamage                    string
	EnableFriendlyFire                            string
	EnableInvaderEnemy                            string
	ActiveUNKO                                    string
	EnableAimAssistPad                            string
	EnableAimAssistKeyboard                       string
	DropItemMaxNum                                string
	DropItemMaxNumDescription                     string
	DropItemMaxNum_UNKO                           string
	BaseCampMaxNum                                string
	BaseCampWorkerMaxNum                          string
	DropItemAliveMaxHours                         string
	AutoResetGuildNoOnlinePlayers                 string
	AutoResetGuildTimeNoOnlinePlayers             string
	GuildPlayerMaxNum                             string
	PalEggDefaultHatchingTime                     string
	WorkSpeedRate                                 string
	IsMultiplay                                   string
	IsPvP                                         string
	CanPickupOtherGuildDeathPenaltyDrop           string
	EnableNonLoginPenalty                         string
	EnableFastTravel                              string
	IsStartLocationSelectByMap                    string
	ExistPlayerAfterLogout                        string
	EnableDefenseOtherGuildPlayer                 string
	CoopPlayerMaxNum                              string
	CoopPlayerMaxNumDescription                   string
	ServerPlayerMaxNum                            string
	ServerName                                    string
	ServerDescription                             string
	AdminPassword                                 string
	ServerPassword                                string
	PublicPort                                    string
	PublicIP                                      string
	RCONEnabled                                   string
	RCONPort                                      string
	Region                                        string
	UseAuth                                       string
	BanListURL                                    string

	DEATH_PENALTY_DROP_ALL                string
	DEATH_PENALTY_DROP_ITEM_AND_EQUIPMENT string
	DEATH_PENALTY_DROP_ITEM               string
	DEATH_PENALTY_NONE                    string
)
