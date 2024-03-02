package settings

import (
	"fmt"
	"os"
	"strings"

	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
	palworld_settings "github.com/juunini/palworld-settings"
)

func setDayTimeSpeedRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.DayTimeSpeedRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.DayTimeSpeedRate, val)
}

func setNightTimeSpeedRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.NightTimeSpeedRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.NightTimeSpeedRate, val)
}

func setExpRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.ExpRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.ExpRate, val)
}

func setPalCaptureRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalCaptureRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PalCaptureRate, val)
}

func setPalSpawnNumRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalSpawnNumRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PalSpawnNumRate, val)
}

func setPalDamageRateAttack(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalDamageRateAttack = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PalDamageRateAttack, val)
}

func setPalDamageRateDefense(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalDamageRateDefense = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PalDamageRateDefense, val)
}

func setPlayerDamageRateAttack(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerDamageRateAttack = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PlayerDamageRateAttack, val)
}

func setPlayerDamageRateDefense(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerDamageRateDefense = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PlayerDamageRateDefense, val)
}

func setPlayerStomachDecreaceRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerStomachDecreaceRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PlayerStomachDecreaceRate, val)
}

func setPlayerStaminaDecreaceRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerStaminaDecreaceRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PlayerStaminaDecreaceRate, val)
}

func setPlayerAutoHPRegeneRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerAutoHPRegeneRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PlayerAutoHPRegeneRate, val)
}

func setPlayerAutoHpRegeneRateInSleep(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerAutoHpRegeneRateInSleep = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PlayerAutoHpRegeneRateInSleep, val)
}

func setPalStomachDecreaceRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalStomachDecreaceRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PalStomachDecreaceRate, val)
}

func setPalStaminaDecreaceRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalStaminaDecreaceRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PalStaminaDecreaceRate, val)
}

func setPalAutoHPRegeneRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalAutoHPRegeneRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PalAutoHPRegeneRate, val)
}

func setPalAutoHpRegeneRateInSleep(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalAutoHpRegeneRateInSleep = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PalAutoHpRegeneRateInSleep, val)
}

func setBuildObjectDamageRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.BuildObjectDamageRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.BuildObjectDamageRate, val)
}

func setBuildObjectDeteriorationDamageRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.BuildObjectDeteriorationDamageRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.BuildObjectDeteriorationDamageRate, val)
}

func setCollectionDropRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.CollectionDropRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.CollectionDropRate, val)
}

func setCollectionObjectHpRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.CollectionObjectHpRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.CollectionObjectHpRate, val)
}

func setCollectionObjectRespawnSpeedRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.CollectionObjectRespawnSpeedRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.CollectionObjectRespawnSpeedRate, val)
}

func setEnemyDropItemRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnemyDropItemRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.EnemyDropItemRate, val)
}

func setDeathPenalty(value string) string {
	val := strings.TrimSpace(value)
	_, ok := map[string]bool{
		palworld_settings.DEATH_PENALTY_DROP_ALL:                true,
		palworld_settings.DEATH_PENALTY_DROP_ITEM_AND_EQUIPMENT: true,
		palworld_settings.DEATH_PENALTY_DROP_ITEM:               true,
		palworld_settings.DEATH_PENALTY_NONE:                    true,
	}[val]

	if !ok {
		return fmt.Sprintf("Invalid DeathPenalty value: %s", value)
	}

	config.PALWORLD_SERVER_SETTINGS.DeathPenalty = val
	saveSettings()
	return fmt.Sprintf("%s: %s", i18n.DeathPenalty, val)
}

func setEnablePlayerToPlayerDamage(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnablePlayerToPlayerDamage = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.EnablePlayerToPlayerDamage, val)
}

func setEnableFriendlyFire(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableFriendlyFire = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.EnableFriendlyFire, val)
}

func setEnableInvaderEnemy(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableInvaderEnemy = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.EnableInvaderEnemy, val)
}

func setActiveUNKO(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.ActiveUNKO = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.ActiveUNKO, val)
}

func setEnableAimAssistPad(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableAimAssistPad = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.EnableAimAssistPad, val)
}

func setEnableAimAssistKeyboard(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableAimAssistKeyboard = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.EnableAimAssistKeyboard, val)
}

func setDropItemMaxNum(value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.DropItemMaxNum = val
	saveSettings()
	return fmt.Sprintf("%s: %d", i18n.DropItemMaxNum, val)
}

func setDropItemMaxNum_UNKO(value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.DropItemMaxNum_UNKO = val
	saveSettings()
	return fmt.Sprintf("%s: %d", i18n.DropItemMaxNum_UNKO, val)
}

func setBaseCampMaxNum(value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.BaseCampMaxNum = val
	saveSettings()
	return fmt.Sprintf("%s: %d", i18n.BaseCampMaxNum, val)
}

func setBaseCampWorkerMaxNum(value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.BaseCampWorkerMaxNum = val
	saveSettings()
	return fmt.Sprintf("%s: %d", i18n.BaseCampWorkerMaxNum, val)
}

func setDropItemAliveMaxHours(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.DropItemAliveMaxHours = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.DropItemAliveMaxHours, val)
}

func setAutoResetGuildNoOnlinePlayers(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.AutoResetGuildNoOnlinePlayers = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.AutoResetGuildNoOnlinePlayers, val)
}

func setAutoResetGuildTimeNoOnlinePlayers(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.AutoResetGuildTimeNoOnlinePlayers = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.AutoResetGuildTimeNoOnlinePlayers, val)
}

func setGuildPlayerMaxNum(value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.GuildPlayerMaxNum = val
	saveSettings()
	return fmt.Sprintf("%s: %d", i18n.GuildPlayerMaxNum, val)
}

func setPalEggDefaultHatchingTime(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalEggDefaultHatchingTime = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.PalEggDefaultHatchingTime, val)
}

func setWorkSpeedRate(value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.WorkSpeedRate = val
	saveSettings()
	return fmt.Sprintf("%s: %f", i18n.WorkSpeedRate, val)
}

func setIsMultiplay(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.IsMultiplay = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.IsMultiplay, val)
}

func setIsPvP(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.IsPvP = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.IsPvP, val)
}

func setCanPickupOtherGuildDeathPenaltyDrop(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.CanPickupOtherGuildDeathPenaltyDrop = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.CanPickupOtherGuildDeathPenaltyDrop, val)
}

func setEnableNonLoginPenalty(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableNonLoginPenalty = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.EnableNonLoginPenalty, val)
}

func setEnableFastTravel(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableFastTravel = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.EnableFastTravel, val)
}

func setIsStartLocationSelectByMap(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.IsStartLocationSelectByMap = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.IsStartLocationSelectByMap, val)
}

func setExistPlayerAfterLogout(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.ExistPlayerAfterLogout = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.ExistPlayerAfterLogout, val)
}

func setEnableDefenseOtherGuildPlayer(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableDefenseOtherGuildPlayer = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.EnableDefenseOtherGuildPlayer, val)
}

func setCoopPlayerMaxNum(value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.CoopPlayerMaxNum = val
	saveSettings()
	return fmt.Sprintf("%s: %d", i18n.CoopPlayerMaxNum, val)
}

func setServerPlayerMaxNum(value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.ServerPlayerMaxNum = val
	saveSettings()
	return fmt.Sprintf("%s: %d", i18n.ServerPlayerMaxNum, val)
}

func setServerName(value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.ServerName = val
	saveSettings()
	return fmt.Sprintf("%s: %s", i18n.ServerName, val)
}

func setServerDescription(value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.ServerDescription = val
	saveSettings()
	return fmt.Sprintf("%s: %s", i18n.ServerDescription, val)
}

func setAdminPassword(value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.AdminPassword = val
	saveSettings()
	return fmt.Sprintf("%s: %s", i18n.AdminPassword, val)
}

func setServerPassword(value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.ServerPassword = val
	saveSettings()
	return fmt.Sprintf("%s: %s", i18n.ServerPassword, val)
}

func setPublicPort(value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PublicPort = val
	saveSettings()
	return fmt.Sprintf("%s: %d", i18n.PublicPort, val)
}

func setPublicIP(value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.PublicIP = val
	saveSettings()
	return fmt.Sprintf("%s: %s", i18n.PublicIP, val)
}

func setRCONEnabled(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.RCONEnabled = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.RCONEnabled, val)
}

func setRCONPort(value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.RCONPort = val
	saveSettings()
	return fmt.Sprintf("%s: %d", i18n.RCONPort, val)
}

func setRegion(value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.Region = val
	saveSettings()
	return fmt.Sprintf("%s: %s", i18n.Region, val)
}

func setUseAuth(value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.UseAuth = val
	saveSettings()
	return fmt.Sprintf("%s: %t", i18n.UseAuth, val)
}

func setBanListURL(value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.BanListURL = val
	saveSettings()
	return fmt.Sprintf("%s: %s", i18n.BanListURL, val)
}

func saveSettings() {
	settingsString := config.PALWORLD_SERVER_SETTINGS.ToString()
	os.WriteFile(config.PALWORLD_SERVER_SETTINGS_FILE_PATH, []byte(settingsString), 0644)
}
