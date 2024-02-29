package palworld_config

import "fmt"

func (p *Config) ToString() string {
	configString := "[/Script/Pal.PalGameWorldSettings]\nOptionSettings=("
	configString += fmt.Sprintf("Difficulty=%s,", p.Difficulty)
	configString += fmt.Sprintf("DayTimeSpeedRate=%f,", p.DayTimeSpeedRate)
	configString += fmt.Sprintf("NightTimeSpeedRate=%f,", p.NightTimeSpeedRate)
	configString += fmt.Sprintf("ExpRate=%f,", p.ExpRate)
	configString += fmt.Sprintf("PalCaptureRate=%f,", p.PalCaptureRate)
	configString += fmt.Sprintf("PalSpawnNumRate=%f,", p.PalSpawnNumRate)
	configString += fmt.Sprintf("PalDamageRateAttack=%f,", p.PalDamageRateAttack)
	configString += fmt.Sprintf("PalDamageRateDefense=%f,", p.PalDamageRateDefense)
	configString += fmt.Sprintf("PlayerDamageRateAttack=%f,", p.PlayerDamageRateAttack)
	configString += fmt.Sprintf("PlayerDamageRateDefense=%f,", p.PlayerDamageRateDefense)
	configString += fmt.Sprintf("PlayerStomachDecreaceRate=%f,", p.PlayerStomachDecreaceRate)
	configString += fmt.Sprintf("PlayerStaminaDecreaceRate=%f,", p.PlayerStaminaDecreaceRate)
	configString += fmt.Sprintf("PlayerAutoHPRegeneRate=%f,", p.PlayerAutoHPRegeneRate)
	configString += fmt.Sprintf("PlayerAutoHpRegeneRateInSleep=%f,", p.PlayerAutoHpRegeneRateInSleep)
	configString += fmt.Sprintf("PalStomachDecreaceRate=%f,", p.PalStomachDecreaceRate)
	configString += fmt.Sprintf("PalStaminaDecreaceRate=%f,", p.PalStaminaDecreaceRate)
	configString += fmt.Sprintf("PalAutoHPRegeneRate=%f,", p.PalAutoHPRegeneRate)
	configString += fmt.Sprintf("PalAutoHpRegeneRateInSleep=%f,", p.PalAutoHpRegeneRateInSleep)
	configString += fmt.Sprintf("BuildObjectDamageRate=%f,", p.BuildObjectDamageRate)
	configString += fmt.Sprintf("BuildObjectDeteriorationDamageRate=%f,", p.BuildObjectDeteriorationDamageRate)
	configString += fmt.Sprintf("CollectionDropRate=%f,", p.CollectionDropRate)
	configString += fmt.Sprintf("CollectionObjectHpRate=%f,", p.CollectionObjectHpRate)
	configString += fmt.Sprintf("CollectionObjectRespawnSpeedRate=%f,", p.CollectionObjectRespawnSpeedRate)
	configString += fmt.Sprintf("EnemyDropItemRate=%f,", p.EnemyDropItemRate)
	configString += fmt.Sprintf("DeathPenalty=%s,", p.DeathPenalty)
	configString += fmt.Sprintf("bEnablePlayerToPlayerDamage=%s,", boolToString(p.EnablePlayerToPlayerDamage))
	configString += fmt.Sprintf("bEnableFriendlyFire=%s,", boolToString(p.EnableFriendlyFire))
	configString += fmt.Sprintf("bEnableInvaderEnemy=%s,", boolToString(p.EnableInvaderEnemy))
	configString += fmt.Sprintf("bActiveUNKO=%s,", boolToString(p.ActiveUNKO))
	configString += fmt.Sprintf("bEnableAimAssistPad=%s,", boolToString(p.EnableAimAssistPad))
	configString += fmt.Sprintf("bEnableAimAssistKeyboard=%s,", boolToString(p.EnableAimAssistKeyboard))
	configString += fmt.Sprintf("DropItemMaxNum=%d,", p.DropItemMaxNum)
	configString += fmt.Sprintf("DropItemMaxNum_UNKO=%d,", p.DropItemMaxNum_UNKO)
	configString += fmt.Sprintf("BaseCampMaxNum=%d,", p.BaseCampMaxNum)
	configString += fmt.Sprintf("BaseCampWorkerMaxNum=%d,", p.BaseCampWorkerMaxNum)
	configString += fmt.Sprintf("DropItemAliveMaxHours=%f,", p.DropItemAliveMaxHours)
	configString += fmt.Sprintf("bAutoResetGuildNoOnlinePlayers=%s,", boolToString(p.AutoResetGuildNoOnlinePlayers))
	configString += fmt.Sprintf("AutoResetGuildTimeNoOnlinePlayers=%f,", p.AutoResetGuildTimeNoOnlinePlayers)
	configString += fmt.Sprintf("GuildPlayerMaxNum=%d,", p.GuildPlayerMaxNum)
	configString += fmt.Sprintf("PalEggDefaultHatchingTime=%f,", p.PalEggDefaultHatchingTime)
	configString += fmt.Sprintf("WorkSpeedRate=%f,", p.WorkSpeedRate)
	configString += fmt.Sprintf("bIsMultiplay=%s,", boolToString(p.IsMultiplay))
	configString += fmt.Sprintf("bIsPvP=%s,", boolToString(p.IsPvP))
	configString += fmt.Sprintf("bCanPickupOtherGuildDeathPenaltyDrop=%s,", boolToString(p.CanPickupOtherGuildDeathPenaltyDrop))
	configString += fmt.Sprintf("bEnableNonLoginPenalty=%s,", boolToString(p.EnableNonLoginPenalty))
	configString += fmt.Sprintf("bEnableFastTravel=%s,", boolToString(p.EnableFastTravel))
	configString += fmt.Sprintf("bIsStartLocationSelectByMap=%s,", boolToString(p.IsStartLocationSelectByMap))
	configString += fmt.Sprintf("bExistPlayerAfterLogout=%s,", boolToString(p.ExistPlayerAfterLogout))
	configString += fmt.Sprintf("bEnableDefenseOtherGuildPlayer=%s,", boolToString(p.EnableDefenseOtherGuildPlayer))
	configString += fmt.Sprintf("CoopPlayerMaxNum=%d,", p.CoopPlayerMaxNum)
	configString += fmt.Sprintf("ServerPlayerMaxNum=%d,", p.ServerPlayerMaxNum)
	configString += fmt.Sprintf(`ServerName="%s",`, p.ServerName)
	configString += fmt.Sprintf(`ServerDescription="%s",`, p.ServerDescription)
	configString += fmt.Sprintf(`AdminPassword="%s",`, p.AdminPassword)
	configString += fmt.Sprintf(`ServerPassword="%s",`, p.ServerPassword)
	configString += fmt.Sprintf("PublicPort=%d,", p.PublicPort)
	configString += fmt.Sprintf(`PublicIP="%s",`, p.PublicIP)
	configString += fmt.Sprintf("RCONEnabled=%s,", boolToString(p.RCONEnabled))
	configString += fmt.Sprintf("RCONPort=%d,", p.RCONPort)
	configString += fmt.Sprintf(`Region="%s",`, p.Region)
	configString += fmt.Sprintf("bUseAuth=%s,", boolToString(p.UseAuth))
	configString += fmt.Sprintf(`BanListURL="%s"`, p.BanListURL)
	configString += ")"
	return configString
}

func boolToString(b bool) string {
	if b {
		return "True"
	}
	return "False"
}
