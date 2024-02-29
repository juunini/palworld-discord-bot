package palworld_config

import (
	"strconv"
	"strings"
)

func Parse(configString string) *Config {
	extractedConfig := configString[strings.Index(configString, "Difficulty="):strings.LastIndex(configString, ")")]

	return &Config{
		Difficulty:                          getConfigValue(extractedConfig, "Difficulty"),
		DayTimeSpeedRate:                    getConfigValueFloat(extractedConfig, "DayTimeSpeedRate"),
		NightTimeSpeedRate:                  getConfigValueFloat(extractedConfig, "NightTimeSpeedRate"),
		ExpRate:                             getConfigValueFloat(extractedConfig, "ExpRate"),
		PalCaptureRate:                      getConfigValueFloat(extractedConfig, "PalCaptureRate"),
		PalSpawnNumRate:                     getConfigValueFloat(extractedConfig, "PalSpawnNumRate"),
		PalDamageRateAttack:                 getConfigValueFloat(extractedConfig, "PalDamageRateAttack"),
		PalDamageRateDefense:                getConfigValueFloat(extractedConfig, "PalDamageRateDefense"),
		PlayerDamageRateAttack:              getConfigValueFloat(extractedConfig, "PlayerDamageRateAttack"),
		PlayerDamageRateDefense:             getConfigValueFloat(extractedConfig, "PlayerDamageRateDefense"),
		PlayerStomachDecreaceRate:           getConfigValueFloat(extractedConfig, "PlayerStomachDecreaceRate"),
		PlayerStaminaDecreaceRate:           getConfigValueFloat(extractedConfig, "PlayerStaminaDecreaceRate"),
		PlayerAutoHPRegeneRate:              getConfigValueFloat(extractedConfig, "PlayerAutoHPRegeneRate"),
		PlayerAutoHpRegeneRateInSleep:       getConfigValueFloat(extractedConfig, "PlayerAutoHpRegeneRateInSleep"),
		PalStomachDecreaceRate:              getConfigValueFloat(extractedConfig, "PalStomachDecreaceRate"),
		PalStaminaDecreaceRate:              getConfigValueFloat(extractedConfig, "PalStaminaDecreaceRate"),
		PalAutoHPRegeneRate:                 getConfigValueFloat(extractedConfig, "PalAutoHPRegeneRate"),
		PalAutoHpRegeneRateInSleep:          getConfigValueFloat(extractedConfig, "PalAutoHpRegeneRateInSleep"),
		BuildObjectDamageRate:               getConfigValueFloat(extractedConfig, "BuildObjectDamageRate"),
		BuildObjectDeteriorationDamageRate:  getConfigValueFloat(extractedConfig, "BuildObjectDeteriorationDamageRate"),
		CollectionDropRate:                  getConfigValueFloat(extractedConfig, "CollectionDropRate"),
		CollectionObjectHpRate:              getConfigValueFloat(extractedConfig, "CollectionObjectHpRate"),
		CollectionObjectRespawnSpeedRate:    getConfigValueFloat(extractedConfig, "CollectionObjectRespawnSpeedRate"),
		EnemyDropItemRate:                   getConfigValueFloat(extractedConfig, "EnemyDropItemRate"),
		DeathPenalty:                        getConfigValue(extractedConfig, "DeathPenalty"),
		EnablePlayerToPlayerDamage:          getConfigValueBool(extractedConfig, "bEnablePlayerToPlayerDamage"),
		EnableFriendlyFire:                  getConfigValueBool(extractedConfig, "bEnableFriendlyFire"),
		EnableInvaderEnemy:                  getConfigValueBool(extractedConfig, "bEnableInvaderEnemy"),
		ActiveUNKO:                          getConfigValueBool(extractedConfig, "bActiveUNKO"),
		EnableAimAssistPad:                  getConfigValueBool(extractedConfig, "bEnableAimAssistPad"),
		EnableAimAssistKeyboard:             getConfigValueBool(extractedConfig, "bEnableAimAssistKeyboard"),
		DropItemMaxNum:                      getConfigValueInt(extractedConfig, "DropItemMaxNum"),
		DropItemMaxNum_UNKO:                 getConfigValueInt(extractedConfig, "DropItemMaxNum_UNKO"),
		BaseCampMaxNum:                      getConfigValueInt(extractedConfig, "BaseCampMaxNum"),
		BaseCampWorkerMaxNum:                getConfigValueInt(extractedConfig, "BaseCampWorkerMaxNum"),
		DropItemAliveMaxHours:               getConfigValueFloat(extractedConfig, "DropItemAliveMaxHours"),
		AutoResetGuildNoOnlinePlayers:       getConfigValueBool(extractedConfig, "bAutoResetGuildNoOnlinePlayers"),
		AutoResetGuildTimeNoOnlinePlayers:   getConfigValueFloat(extractedConfig, "AutoResetGuildTimeNoOnlinePlayers"),
		GuildPlayerMaxNum:                   getConfigValueInt(extractedConfig, "GuildPlayerMaxNum"),
		PalEggDefaultHatchingTime:           getConfigValueFloat(extractedConfig, "PalEggDefaultHatchingTime"),
		WorkSpeedRate:                       getConfigValueFloat(extractedConfig, "WorkSpeedRate"),
		IsMultiplay:                         getConfigValueBool(extractedConfig, "bIsMultiplay"),
		IsPvP:                               getConfigValueBool(extractedConfig, "bIsPvP"),
		CanPickupOtherGuildDeathPenaltyDrop: getConfigValueBool(extractedConfig, "bCanPickupOtherGuildDeathPenaltyDrop"),
		EnableNonLoginPenalty:               getConfigValueBool(extractedConfig, "bEnableNonLoginPenalty"),
		EnableFastTravel:                    getConfigValueBool(extractedConfig, "bEnableFastTravel"),
		IsStartLocationSelectByMap:          getConfigValueBool(extractedConfig, "bIsStartLocationSelectByMap"),
		ExistPlayerAfterLogout:              getConfigValueBool(extractedConfig, "bExistPlayerAfterLogout"),
		EnableDefenseOtherGuildPlayer:       getConfigValueBool(extractedConfig, "bEnableDefenseOtherGuildPlayer"),
		CoopPlayerMaxNum:                    getConfigValueInt(extractedConfig, "CoopPlayerMaxNum"),
		ServerPlayerMaxNum:                  getConfigValueInt(extractedConfig, "ServerPlayerMaxNum"),
		ServerName:                          getConfigValue(extractedConfig, "ServerName"),
		ServerDescription:                   getConfigValue(extractedConfig, "ServerDescription"),
		AdminPassword:                       getConfigValue(extractedConfig, "AdminPassword"),
		ServerPassword:                      getConfigValue(extractedConfig, "ServerPassword"),
		PublicPort:                          getConfigValueInt(extractedConfig, "PublicPort"),
		PublicIP:                            getConfigValue(extractedConfig, "PublicIP"),
		RCONEnabled:                         getConfigValueBool(extractedConfig, "RCONEnabled"),
		RCONPort:                            getConfigValueInt(extractedConfig, "RCONPort"),
		Region:                              getConfigValue(extractedConfig, "Region"),
		UseAuth:                             getConfigValueBool(extractedConfig, "bUseAuth"),
		BanListURL:                          getConfigValue(extractedConfig, "BanListURL"),
	}
}

func getConfigValue(configString, configName string) string {
	configNameStartIndex := strings.Index(configString, configName)
	if configNameStartIndex == -1 {
		return ""
	}

	valueStart := strings.TrimPrefix(configString[strings.Index(configString, configName):], configName+"=")

	configLastIndex := strings.Index(valueStart, ",")
	if configLastIndex == -1 {
		return valueStart
	}
	return valueStart[:strings.Index(valueStart, ",")]
}

func getConfigValueFloat(configString, configName string) float64 {
	value, err := strconv.ParseFloat(getConfigValue(configString, configName), 64)
	if err != nil {
		return 1
	}
	return value
}

func getConfigValueInt(configString, configName string) int {
	value, err := strconv.Atoi(getConfigValue(configString, configName))
	if err != nil {
		return 1
	}
	return value
}

func getConfigValueBool(configString, configName string) bool {
	return getConfigValue(configString, configName) == "true"
}
