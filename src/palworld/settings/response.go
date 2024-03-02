package settings

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
)

func Response(command string) []string {
	after, found := strings.CutPrefix(command, config.DISCORD_COMMAND_ALIAS_SERVER_SETTINGS+" ")
	if !found || strings.TrimSpace(after) == "" {
		return i18n.SettingsHelp(config.DISCORD_COMMAND_PREFIX, config.DISCORD_COMMAND_ALIAS_SERVER_SETTINGS)
	}

	if strings.HasPrefix(after, "DayTimeSpeedRate ") {
		value, _ := strings.CutPrefix(after, "DayTimeSpeedRate ")
		return []string{setDayTimeSpeedRate(value)}
	} else if strings.HasPrefix(after, "NightTimeSpeedRate ") {
		value, _ := strings.CutPrefix(after, "NightTimeSpeedRate ")
		return []string{setNightTimeSpeedRate(value)}
	} else if strings.HasPrefix(after, "ExpRate ") {
		value, _ := strings.CutPrefix(after, "ExpRate ")
		return []string{setExpRate(value)}
	} else if strings.HasPrefix(after, "PalCaptureRate ") {
		value, _ := strings.CutPrefix(after, "PalCaptureRate ")
		return []string{setPalCaptureRate(value)}
	} else if strings.HasPrefix(after, "PalSpawnNumRate ") {
		value, _ := strings.CutPrefix(after, "PalSpawnNumRate ")
		return []string{setPalSpawnNumRate(value)}
	} else if strings.HasPrefix(after, "PalDamageRateAttack ") {
		value, _ := strings.CutPrefix(after, "PalDamageRateAttack ")
		return []string{setPalDamageRateAttack(value)}
	} else if strings.HasPrefix(after, "PalDamageRateDefense ") {
		value, _ := strings.CutPrefix(after, "PalDamageRateDefense ")
		return []string{setPalDamageRateDefense(value)}
	} else if strings.HasPrefix(after, "PlayerDamageRateAttack ") {
		value, _ := strings.CutPrefix(after, "PlayerDamageRateAttack ")
		return []string{setPlayerDamageRateAttack(value)}
	} else if strings.HasPrefix(after, "PlayerDamageRateDefense ") {
		value, _ := strings.CutPrefix(after, "PlayerDamageRateDefense ")
		return []string{setPlayerDamageRateDefense(value)}
	} else if strings.HasPrefix(after, "PlayerStomachDecreaceRate ") {
		value, _ := strings.CutPrefix(after, "PlayerStomachDecreaceRate ")
		return []string{setPlayerStomachDecreaceRate(value)}
	} else if strings.HasPrefix(after, "PlayerStaminaDecreaceRate ") {
		value, _ := strings.CutPrefix(after, "PlayerStaminaDecreaceRate ")
		return []string{setPlayerStaminaDecreaceRate(value)}
	} else if strings.HasPrefix(after, "PlayerAutoHPRegeneRate ") {
		value, _ := strings.CutPrefix(after, "PlayerAutoHPRegeneRate ")
		return []string{setPlayerAutoHPRegeneRate(value)}
	} else if strings.HasPrefix(after, "PlayerAutoHpRegeneRateInSleep ") {
		value, _ := strings.CutPrefix(after, "PlayerAutoHpRegeneRateInSleep ")
		return []string{setPlayerAutoHpRegeneRateInSleep(value)}
	} else if strings.HasPrefix(after, "PalStomachDecreaceRate ") {
		value, _ := strings.CutPrefix(after, "PalStomachDecreaceRate ")
		return []string{setPalStomachDecreaceRate(value)}
	} else if strings.HasPrefix(after, "PalStaminaDecreaceRate ") {
		value, _ := strings.CutPrefix(after, "PalStaminaDecreaceRate ")
		return []string{setPalStaminaDecreaceRate(value)}
	} else if strings.HasPrefix(after, "PalAutoHPRegeneRate ") {
		value, _ := strings.CutPrefix(after, "PalAutoHPRegeneRate ")
		return []string{setPalAutoHPRegeneRate(value)}
	} else if strings.HasPrefix(after, "PalAutoHpRegeneRateInSleep ") {
		value, _ := strings.CutPrefix(after, "PalAutoHpRegeneRateInSleep ")
		return []string{setPalAutoHpRegeneRateInSleep(value)}
	} else if strings.HasPrefix(after, "BuildObjectDamageRate ") {
		value, _ := strings.CutPrefix(after, "BuildObjectDamageRate ")
		return []string{setBuildObjectDamageRate(value)}
	} else if strings.HasPrefix(after, "BuildObjectDeteriorationDamageRate ") {
		value, _ := strings.CutPrefix(after, "BuildObjectDeteriorationDamageRate ")
		return []string{setBuildObjectDeteriorationDamageRate(value)}
	} else if strings.HasPrefix(after, "CollectionDropRate ") {
		value, _ := strings.CutPrefix(after, "CollectionDropRate ")
		return []string{setCollectionDropRate(value)}
	} else if strings.HasPrefix(after, "CollectionObjectHpRate ") {
		value, _ := strings.CutPrefix(after, "CollectionObjectHpRate ")
		return []string{setCollectionObjectHpRate(value)}
	} else if strings.HasPrefix(after, "CollectionObjectRespawnSpeedRate ") {
		value, _ := strings.CutPrefix(after, "CollectionObjectRespawnSpeedRate ")
		return []string{setCollectionObjectRespawnSpeedRate(value)}
	} else if strings.HasPrefix(after, "EnemyDropItemRate ") {
		value, _ := strings.CutPrefix(after, "EnemyDropItemRate ")
		return []string{setEnemyDropItemRate(value)}
	} else if strings.HasPrefix(after, "DeathPenalty ") {
		value, _ := strings.CutPrefix(after, "DeathPenalty ")
		return []string{setDeathPenalty(value)}
	} else if strings.HasPrefix(after, "EnablePlayerToPlayerDamage ") {
		value, _ := strings.CutPrefix(after, "EnablePlayerToPlayerDamage ")
		return []string{setEnablePlayerToPlayerDamage(value)}
	} else if strings.HasPrefix(after, "EnableFriendlyFire ") {
		value, _ := strings.CutPrefix(after, "EnableFriendlyFire ")
		return []string{setEnableFriendlyFire(value)}
	} else if strings.HasPrefix(after, "EnableInvaderEnemy ") {
		value, _ := strings.CutPrefix(after, "EnableInvaderEnemy ")
		return []string{setEnableInvaderEnemy(value)}
	} else if strings.HasPrefix(after, "ActiveUNKO ") {
		value, _ := strings.CutPrefix(after, "ActiveUNKO ")
		return []string{setActiveUNKO(value)}
	} else if strings.HasPrefix(after, "EnableAimAssistPad ") {
		value, _ := strings.CutPrefix(after, "EnableAimAssistPad ")
		return []string{setEnableAimAssistPad(value)}
	} else if strings.HasPrefix(after, "EnableAimAssistKeyboard ") {
		value, _ := strings.CutPrefix(after, "EnableAimAssistKeyboard ")
		return []string{setEnableAimAssistKeyboard(value)}
	} else if strings.HasPrefix(after, "DropItemMaxNum ") {
		value, _ := strings.CutPrefix(after, "DropItemMaxNum ")
		return []string{setDropItemMaxNum(value)}
	} else if strings.HasPrefix(after, "DropItemMaxNum_UNKO ") {
		value, _ := strings.CutPrefix(after, "DropItemMaxNum_UNKO ")
		return []string{setDropItemMaxNum_UNKO(value)}
	} else if strings.HasPrefix(after, "BaseCampMaxNum ") {
		value, _ := strings.CutPrefix(after, "BaseCampMaxNum ")
		return []string{setBaseCampMaxNum(value)}
	} else if strings.HasPrefix(after, "BaseCampWorkerMaxNum ") {
		value, _ := strings.CutPrefix(after, "BaseCampWorkerMaxNum ")
		return []string{setBaseCampWorkerMaxNum(value)}
	} else if strings.HasPrefix(after, "DropItemAliveMaxHours ") {
		value, _ := strings.CutPrefix(after, "DropItemAliveMaxHours ")
		return []string{setDropItemAliveMaxHours(value)}
	} else if strings.HasPrefix(after, "AutoResetGuildNoOnlinePlayers ") {
		value, _ := strings.CutPrefix(after, "AutoResetGuildNoOnlinePlayers ")
		return []string{setAutoResetGuildNoOnlinePlayers(value)}
	} else if strings.HasPrefix(after, "AutoResetGuildTimeNoOnlinePlayers ") {
		value, _ := strings.CutPrefix(after, "AutoResetGuildTimeNoOnlinePlayers ")
		return []string{setAutoResetGuildTimeNoOnlinePlayers(value)}
	} else if strings.HasPrefix(after, "GuildPlayerMaxNum ") {
		value, _ := strings.CutPrefix(after, "GuildPlayerMaxNum ")
		return []string{setGuildPlayerMaxNum(value)}
	} else if strings.HasPrefix(after, "PalEggDefaultHatchingTime ") {
		value, _ := strings.CutPrefix(after, "PalEggDefaultHatchingTime ")
		return []string{setPalEggDefaultHatchingTime(value)}
	} else if strings.HasPrefix(after, "WorkSpeedRate ") {
		value, _ := strings.CutPrefix(after, "WorkSpeedRate ")
		return []string{setWorkSpeedRate(value)}
	} else if strings.HasPrefix(after, "IsMultiplay ") {
		value, _ := strings.CutPrefix(after, "IsMultiplay ")
		return []string{setIsMultiplay(value)}
	} else if strings.HasPrefix(after, "IsPvP ") {
		value, _ := strings.CutPrefix(after, "IsPvP ")
		return []string{setIsPvP(value)}
	} else if strings.HasPrefix(after, "CanPickupOtherGuildDeathPenaltyDrop ") {
		value, _ := strings.CutPrefix(after, "CanPickupOtherGuildDeathPenaltyDrop ")
		return []string{setCanPickupOtherGuildDeathPenaltyDrop(value)}
	} else if strings.HasPrefix(after, "EnableNonLoginPenalty ") {
		value, _ := strings.CutPrefix(after, "EnableNonLoginPenalty ")
		return []string{setEnableNonLoginPenalty(value)}
	} else if strings.HasPrefix(after, "EnableFastTravel ") {
		value, _ := strings.CutPrefix(after, "EnableFastTravel ")
		return []string{setEnableFastTravel(value)}
	} else if strings.HasPrefix(after, "IsStartLocationSelectByMap ") {
		value, _ := strings.CutPrefix(after, "IsStartLocationSelectByMap ")
		return []string{setIsStartLocationSelectByMap(value)}
	} else if strings.HasPrefix(after, "ExistPlayerAfterLogout ") {
		value, _ := strings.CutPrefix(after, "ExistPlayerAfterLogout ")
		return []string{setExistPlayerAfterLogout(value)}
	} else if strings.HasPrefix(after, "EnableDefenseOtherGuildPlayer ") {
		value, _ := strings.CutPrefix(after, "EnableDefenseOtherGuildPlayer ")
		return []string{setEnableDefenseOtherGuildPlayer(value)}
	} else if strings.HasPrefix(after, "CoopPlayerMaxNum ") {
		value, _ := strings.CutPrefix(after, "CoopPlayerMaxNum ")
		return []string{setCoopPlayerMaxNum(value)}
	} else if strings.HasPrefix(after, "ServerPlayerMaxNum ") {
		value, _ := strings.CutPrefix(after, "ServerPlayerMaxNum ")
		return []string{setServerPlayerMaxNum(value)}
	} else if strings.HasPrefix(after, "ServerName ") {
		value, _ := strings.CutPrefix(after, "ServerName ")
		return []string{setServerName(value)}
	} else if strings.HasPrefix(after, "ServerDescription ") {
		value, _ := strings.CutPrefix(after, "ServerDescription ")
		return []string{setServerDescription(value)}
	} else if strings.HasPrefix(after, "AdminPassword ") {
		value, _ := strings.CutPrefix(after, "AdminPassword ")
		return []string{setAdminPassword(value)}
	} else if strings.HasPrefix(after, "ServerPassword ") {
		value, _ := strings.CutPrefix(after, "ServerPassword ")
		return []string{setServerPassword(value)}
	} else if strings.HasPrefix(after, "PublicPort ") {
		value, _ := strings.CutPrefix(after, "PublicPort ")
		return []string{setPublicPort(value)}
	} else if strings.HasPrefix(after, "PublicIP ") {
		value, _ := strings.CutPrefix(after, "PublicIP ")
		return []string{setPublicIP(value)}
	} else if strings.HasPrefix(after, "RCONEnabled ") {
		value, _ := strings.CutPrefix(after, "RCONEnabled ")
		return []string{setRCONEnabled(value)}
	} else if strings.HasPrefix(after, "RCONPort ") {
		value, _ := strings.CutPrefix(after, "RCONPort ")
		return []string{setRCONPort(value)}
	} else if strings.HasPrefix(after, "Region ") {
		value, _ := strings.CutPrefix(after, "Region ")
		return []string{setRegion(value)}
	} else if strings.HasPrefix(after, "UseAuth ") {
		value, _ := strings.CutPrefix(after, "UseAuth ")
		return []string{setUseAuth(value)}
	} else if strings.HasPrefix(after, "BanListURL ") {
		value, _ := strings.CutPrefix(after, "BanListURL ")
		return []string{setBanListURL(value)}
	}

	return []string{"Unknown setting command"}
}
