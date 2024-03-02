package i18n

import (
	"fmt"

	palworld_settings "github.com/juunini/palworld-settings"
)

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
	CommandPrefix       string
	HelpAlias           string
	KickAlias           string
	BanAlias            string
	BroadcastAlias      string
	ShutdownAlias       string
	DoExitAlias         string
	SaveAlias           string
	StartServerAlias    string
	ServerSettingsAlias string
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

	Close        string
	Confirm      string
	Save         string
	DefaultValue string

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

	SettingsHelp = func(CommandPrefix, settingsAlias string) []string {
		messages := []string{}

		message := fmt.Sprintf(
			"- `%s %s DayTimeSpeedRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, DayTimeSpeedRate,
			palworld_settings.DAY_TIME_SPEED_RATE_MIN,
			float64(palworld_settings.DAY_TIME_SPEED_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.DayTimeSpeedRate,
		)
		message += fmt.Sprintf(
			"- `%s %s NightTimeSpeedRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, NightTimeSpeedRate,
			palworld_settings.NIGHT_TIME_SPEED_RATE_MIN,
			float64(palworld_settings.NIGHT_TIME_SPEED_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.NightTimeSpeedRate,
		)
		message += fmt.Sprintf(
			"- `%s %s ExpRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, ExpRate,
			palworld_settings.EXP_RATE_MIN,
			float64(palworld_settings.EXP_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.ExpRate,
		)
		message += fmt.Sprintf(
			"- `%s %s PalCaptureRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PalCaptureRate,
			palworld_settings.PAL_CAPTURE_RATE_MIN,
			float64(palworld_settings.PAL_CAPTURE_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PalCaptureRate,
		)
		message += fmt.Sprintf(
			"- `%s %s PalSpawnNumRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PalSpawnNumRate,
			palworld_settings.PAL_SPAWN_NUM_RATE_MIN,
			float64(palworld_settings.PAL_SPAWN_NUM_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PalSpawnNumRate,
		)
		message += fmt.Sprintf(
			"- `%s %s PalDamageRateAttack <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PalDamageRateAttack,
			palworld_settings.PAL_DAMAGE_RATE_ATTACK_MIN,
			float64(palworld_settings.PAL_DAMAGE_RATE_ATTACK_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PalDamageRateAttack,
		)
		message += fmt.Sprintf(
			"- `%s %s PalDamageRateDefense <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PalDamageRateDefense,
			palworld_settings.PAL_DAMAGE_RATE_DEFENSE_MIN,
			float64(palworld_settings.PAL_DAMAGE_RATE_DEFENSE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PalDamageRateDefense,
		)
		message += fmt.Sprintf(
			"- `%s %s PlayerDamageRateAttack <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PlayerDamageRateAttack,
			palworld_settings.PLAYER_DAMAGE_RATE_ATTACK_MIN,
			float64(palworld_settings.PLAYER_DAMAGE_RATE_ATTACK_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PlayerDamageRateAttack,
		)
		message += fmt.Sprintf(
			"- `%s %s PlayerDamageRateDefense <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PlayerDamageRateDefense,
			palworld_settings.PLAYER_DAMAGE_RATE_DEFENSE_MIN,
			float64(palworld_settings.PLAYER_DAMAGE_RATE_DEFENSE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PlayerDamageRateDefense,
		)
		message += fmt.Sprintf(
			"- `%s %s PlayerStomachDecreaceRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PlayerStomachDecreaceRate,
			palworld_settings.PLAYER_STOMACH_DECREASE_RATE_MIN,
			float64(palworld_settings.PLAYER_STOMACH_DECREASE_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PlayerStomachDecreaceRate,
		)

		messages = append(messages, message)

		message = fmt.Sprintf(
			"- `%s %s PlayerStaminaDecreaceRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PlayerStaminaDecreaceRate,
			palworld_settings.PLAYER_STAMINA_DECREASE_RATE_MIN,
			float64(palworld_settings.PLAYER_STAMINA_DECREASE_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PlayerStaminaDecreaceRate,
		)
		message += fmt.Sprintf(
			"- `%s %s PlayerAutoHPRegeneRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PlayerAutoHPRegeneRate,
			palworld_settings.PLAYER_AUTO_HP_REGENE_RATE_MIN,
			float64(palworld_settings.PLAYER_AUTO_HP_REGENE_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PlayerAutoHPRegeneRate,
		)
		message += fmt.Sprintf(
			"- `%s %s PlayerAutoHpRegeneRateInSleep <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PlayerAutoHpRegeneRateInSleep,
			palworld_settings.PLAYER_AUTO_HP_REGENE_RATE_IN_SLEEP_MIN,
			float64(palworld_settings.PLAYER_AUTO_HP_REGENE_RATE_IN_SLEEP_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PlayerAutoHpRegeneRateInSleep,
		)
		message += fmt.Sprintf(
			"- `%s %s PalStomachDecreaceRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PalStomachDecreaceRate,
			palworld_settings.PAL_STOMACH_DECREASE_RATE_MIN,
			float64(palworld_settings.PAL_STOMACH_DECREASE_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PalStomachDecreaceRate,
		)
		message += fmt.Sprintf(
			"- `%s %s PalStaminaDecreaceRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PalStaminaDecreaceRate,
			palworld_settings.PAL_STAMINA_DECREASE_RATE_MIN,
			float64(palworld_settings.PAL_STAMINA_DECREASE_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PalStaminaDecreaceRate,
		)
		message += fmt.Sprintf(
			"- `%s %s PalAutoHPRegeneRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PalAutoHPRegeneRate,
			palworld_settings.PAL_AUTO_HP_REGENE_RATE_MIN,
			float64(palworld_settings.PAL_AUTO_HP_REGENE_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PalAutoHPRegeneRate,
		)
		message += fmt.Sprintf(
			"- `%s %s PalAutoHpRegeneRateInSleep <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PalAutoHpRegeneRateInSleep,
			palworld_settings.PAL_AUTO_HP_REGENE_RATE_IN_SLEEP_MIN,
			float64(palworld_settings.PAL_AUTO_HP_REGENE_RATE_IN_SLEEP_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.PalAutoHpRegeneRateInSleep,
		)
		message += fmt.Sprintf(
			"- `%s %s BuildObjectDamageRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, BuildObjectDamageRate,
			palworld_settings.BUILD_OBJECT_DAMAGE_RATE_MIN,
			float64(palworld_settings.BUILD_OBJECT_DAMAGE_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.BuildObjectDamageRate,
		)
		message += fmt.Sprintf(
			"- `%s %s BuildObjectDeteriorationDamageRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, BuildObjectDeteriorationDamageRate,
			float64(palworld_settings.BUILD_OBJECT_DETERIORATION_DAMAGE_RATE_MIN),
			float64(palworld_settings.BUILD_OBJECT_DETERIORATION_DAMAGE_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.BuildObjectDeteriorationDamageRate,
		)
		message += fmt.Sprintf(
			"- `%s %s CollectionDropRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, CollectionDropRate,
			palworld_settings.COLLECTION_DROP_RATE_MIN,
			float64(palworld_settings.COLLECTION_DROP_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.CollectionDropRate,
		)

		messages = append(messages, message)

		message = fmt.Sprintf(
			"- `%s %s CollectionObjectHpRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, CollectionObjectHpRate,
			palworld_settings.COLLECTION_OBJECT_HP_RATE_MIN,
			float64(palworld_settings.COLLECTION_OBJECT_HP_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.CollectionObjectHpRate,
		)
		message += fmt.Sprintf(
			"- `%s %s CollectionObjectRespawnSpeedRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias, CollectionObjectRespawnSpeedRate,
			palworld_settings.COLLECTION_OBJECT_RESPAWN_SPEED_RATE_MIN,
			float64(palworld_settings.COLLECTION_OBJECT_RESPAWN_SPEED_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.CollectionObjectRespawnSpeedRate,
		)
		message += fmt.Sprintf(
			"- `%s %s EnemyDropItemRate <value>` - %s (%.1f ~ %.1f, %s: %.1f)\n",
			CommandPrefix, settingsAlias,
			EnemyDropItemRate,
			palworld_settings.ENEMY_DROP_ITEM_RATE_MIN,
			float64(palworld_settings.ENEMY_DROP_ITEM_RATE_MAX),
			DefaultValue,
			palworld_settings.DEFAULT.EnemyDropItemRate,
		)
		message += fmt.Sprintf(
			"- `%s %s DeathPenalty <value>` - %s\n",
			CommandPrefix, settingsAlias, DeathPenalty,
		)
		message += fmt.Sprintf(
			"  - `%s`: %s (%s)\n",
			palworld_settings.DEATH_PENALTY_DROP_ALL, DEATH_PENALTY_DROP_ALL, DefaultValue,
		)
		message += fmt.Sprintf(
			"  - `%s`: %s\n",
			palworld_settings.DEATH_PENALTY_DROP_ITEM_AND_EQUIPMENT, DEATH_PENALTY_DROP_ITEM_AND_EQUIPMENT,
		)
		message += fmt.Sprintf(
			"  - `%s`: %s\n",
			palworld_settings.DEATH_PENALTY_DROP_ITEM, DEATH_PENALTY_DROP_ITEM,
		)
		message += fmt.Sprintf(
			"  - `%s`: %s\n",
			palworld_settings.DEATH_PENALTY_NONE, DEATH_PENALTY_NONE,
		)

		messages = append(messages, message)

		message = fmt.Sprintf(
			"- `%s %s EnablePlayerToPlayerDamage <true or false>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, EnablePlayerToPlayerDamage,
			DefaultValue, palworld_settings.DEFAULT.EnablePlayerToPlayerDamage,
		)
		message += fmt.Sprintf(
			"- `%s %s EnableFriendlyFire <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, EnableFriendlyFire,
			DefaultValue, palworld_settings.DEFAULT.EnableFriendlyFire,
		)
		message += fmt.Sprintf(
			"- `%s %s EnableInvaderEnemy <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, EnableInvaderEnemy,
			DefaultValue, palworld_settings.DEFAULT.EnableInvaderEnemy,
		)
		message += fmt.Sprintf(
			"- `%s %s ActiveUNKO <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, ActiveUNKO,
			DefaultValue, palworld_settings.DEFAULT.ActiveUNKO,
		)
		message += fmt.Sprintf(
			"- `%s %s EnableAimAssistPad <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, EnableAimAssistPad,
			DefaultValue, palworld_settings.DEFAULT.EnableAimAssistPad,
		)
		message += fmt.Sprintf(
			"- `%s %s EnableAimAssistKeyboard <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, EnableAimAssistKeyboard,
			DefaultValue, palworld_settings.DEFAULT.EnableAimAssistKeyboard,
		)
		message += fmt.Sprintf(
			"- `%s %s DropItemMaxNum <value>` - %s (%d ~ %d, %s: %d)\n",
			CommandPrefix, settingsAlias, DropItemMaxNum,
			palworld_settings.DROP_ITEM_MAX_NUM_MIN,
			palworld_settings.DROP_ITEM_MAX_NUM_MAX,
			DefaultValue,
			palworld_settings.DEFAULT.DropItemMaxNum,
		)
		message += fmt.Sprintf(
			"- `%s %s DropItemMaxNum_UNKO <value>` - %s (%s: %d)\n",
			CommandPrefix, settingsAlias, DropItemMaxNum_UNKO,
			DefaultValue, palworld_settings.DEFAULT.DropItemMaxNum_UNKO,
		)
		message += fmt.Sprintf(
			"- `%s %s BaseCampMaxNum <value>` - %s (%s: %d)\n",
			CommandPrefix, settingsAlias, BaseCampMaxNum,
			DefaultValue, palworld_settings.DEFAULT.BaseCampMaxNum,
		)
		message += fmt.Sprintf(
			"- `%s %s BaseCampWorkerMaxNum <value>` - %s (%d ~ %d, %s: %d)\n",
			CommandPrefix, settingsAlias, BaseCampWorkerMaxNum,
			palworld_settings.BASE_CAMP_WORKER_MAX_NUM_MIN, palworld_settings.BASE_CAMP_WORKER_MAX_NUM_MAX,
			DefaultValue, palworld_settings.DEFAULT.BaseCampWorkerMaxNum,
		)

		messages = append(messages, message)

		message = fmt.Sprintf(
			"- `%s %s DropItemAliveMaxHours <value>` - %s (%s: %.1f)\n",
			CommandPrefix, settingsAlias, DropItemAliveMaxHours,
			DefaultValue, palworld_settings.DEFAULT.DropItemAliveMaxHours,
		)
		message += fmt.Sprintf(
			"- `%s %s AutoResetGuildNoOnlinePlayers <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, AutoResetGuildNoOnlinePlayers,
			DefaultValue, palworld_settings.DEFAULT.AutoResetGuildNoOnlinePlayers,
		)
		message += fmt.Sprintf(
			"- `%s %s AutoResetGuildTimeNoOnlinePlayers <value>` - %s (%s: %.1f)\n",
			CommandPrefix, settingsAlias, AutoResetGuildTimeNoOnlinePlayers,
			DefaultValue, palworld_settings.DEFAULT.AutoResetGuildTimeNoOnlinePlayers,
		)
		message += fmt.Sprintf(
			"- `%s %s GuildPlayerMaxNum <value>` - %s (%d ~ %d, %s: %d)\n",
			CommandPrefix, settingsAlias, GuildPlayerMaxNum,
			palworld_settings.GUILD_PLAYER_MAX_NUM_MIN, palworld_settings.GUILD_PLAYER_MAX_NUM_MAX,
			DefaultValue, palworld_settings.DEFAULT.GuildPlayerMaxNum,
		)
		message += fmt.Sprintf(
			"- `%s %s PalEggDefaultHatchingTime <value>` - %s (%d ~ %d, %s: %.1f)\n",
			CommandPrefix, settingsAlias, PalEggDefaultHatchingTime,
			palworld_settings.PAL_EGG_DEFAULT_HATCHING_TIME_MIN, palworld_settings.PAL_EGG_DEFAULT_HATCHING_TIME_MAX,
			DefaultValue, palworld_settings.DEFAULT.PalEggDefaultHatchingTime,
		)
		message += fmt.Sprintf(
			"- `%s %s WorkSpeedRate <value>` - %s (%s: %.1f)\n",
			CommandPrefix, settingsAlias, WorkSpeedRate,
			DefaultValue, palworld_settings.DEFAULT.WorkSpeedRate,
		)
		message += fmt.Sprintf(
			"- `%s %s IsMultiplay <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, IsMultiplay,
			DefaultValue, palworld_settings.DEFAULT.IsMultiplay,
		)
		message += fmt.Sprintf(
			"- `%s %s IsPvP <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, IsPvP,
			DefaultValue, palworld_settings.DEFAULT.IsPvP,
		)
		message += fmt.Sprintf(
			"- `%s %s CanPickupOtherGuildDeathPenaltyDrop <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, CanPickupOtherGuildDeathPenaltyDrop,
			DefaultValue, palworld_settings.DEFAULT.CanPickupOtherGuildDeathPenaltyDrop,
		)
		message += fmt.Sprintf(
			"- `%s %s EnableNonLoginPenalty <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, EnableNonLoginPenalty,
			DefaultValue, palworld_settings.DEFAULT.EnableNonLoginPenalty,
		)

		messages = append(messages, message)

		message = fmt.Sprintf(
			"- `%s %s EnableFastTravel <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, EnableFastTravel,
			DefaultValue, palworld_settings.DEFAULT.EnableFastTravel,
		)
		message += fmt.Sprintf(
			"- `%s %s IsStartLocationSelectByMap <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, IsStartLocationSelectByMap,
			DefaultValue, palworld_settings.DEFAULT.IsStartLocationSelectByMap,
		)
		message += fmt.Sprintf(
			"- `%s %s ExistPlayerAfterLogout <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, ExistPlayerAfterLogout,
			DefaultValue, palworld_settings.DEFAULT.ExistPlayerAfterLogout,
		)
		message += fmt.Sprintf(
			"- `%s %s EnableDefenseOtherGuildPlayer <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, EnableDefenseOtherGuildPlayer,
			DefaultValue, palworld_settings.DEFAULT.EnableDefenseOtherGuildPlayer,
		)
		message += fmt.Sprintf(
			"- `%s %s CoopPlayerMaxNum <value>` - %s (%s: %d)\n",
			CommandPrefix, settingsAlias, CoopPlayerMaxNum,
			DefaultValue, palworld_settings.DEFAULT.CoopPlayerMaxNum,
		)
		message += fmt.Sprintf(
			"- `%s %s ServerPlayerMaxNum <value>` - %s (%s: %d)\n",
			CommandPrefix, settingsAlias, ServerPlayerMaxNum,
			DefaultValue, palworld_settings.DEFAULT.ServerPlayerMaxNum,
		)
		message += fmt.Sprintf(
			"- `%s %s ServerName <value>` - %s (%s: %s)\n",
			CommandPrefix, settingsAlias, ServerName,
			DefaultValue, palworld_settings.DEFAULT.ServerName,
		)
		message += fmt.Sprintf(
			"- `%s %s ServerDescription <value>` - %s\n",
			CommandPrefix, settingsAlias, ServerDescription,
		)
		message += fmt.Sprintf(
			"- `%s %s AdminPassword <value>` - %s\n",
			CommandPrefix, settingsAlias, AdminPassword,
		)
		message += fmt.Sprintf(
			"- `%s %s ServerPassword <value>` - %s\n",
			CommandPrefix, settingsAlias, ServerPassword,
		)

		messages = append(messages, message)

		message = fmt.Sprintf(
			"- `%s %s PublicPort <value>` - %s (%s: %d)\n",
			CommandPrefix, settingsAlias, PublicPort,
			DefaultValue, palworld_settings.DEFAULT.PublicPort,
		)
		message += fmt.Sprintf(
			"- `%s %s PublicIP <value>` - %s\n",
			CommandPrefix, settingsAlias, PublicIP,
		)
		message += fmt.Sprintf(
			"- `%s %s RCONEnabled <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, RCONEnabled,
			DefaultValue, palworld_settings.DEFAULT.RCONEnabled,
		)
		message += fmt.Sprintf(
			"- `%s %s RCONPort <value>` - %s (%s: %d)\n",
			CommandPrefix, settingsAlias, RCONPort,
			DefaultValue, palworld_settings.DEFAULT.RCONPort,
		)
		message += fmt.Sprintf(
			"- `%s %s Region <value>` - %s\n",
			CommandPrefix, settingsAlias, Region,
		)
		message += fmt.Sprintf(
			"- `%s %s UseAuth <value>` - %s (%s: %t)\n",
			CommandPrefix, settingsAlias, UseAuth,
			DefaultValue, palworld_settings.DEFAULT.UseAuth,
		)
		message += fmt.Sprintf(
			"- `%s %s BanListURL <value>` - %s (%s: %s)\n",
			CommandPrefix, settingsAlias, BanListURL,
			DefaultValue, palworld_settings.DEFAULT.BanListURL,
		)

		messages = append(messages, message)

		return messages
	}
)
