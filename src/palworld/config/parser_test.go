package palworld_config_test

import (
	"testing"

	palworld_config "github.com/juunini/palworld-discord-bot/src/palworld/config"
)

func TestParse(t *testing.T) {
	configString := `Difficulty=Normal, DayTimeSpeedRate=1.0, NightTimeSpeedRate=1.0, ExpRate=1.0, PalCaptureRate=1.0, PalSpawnNumRate=1.0, PalDamageRateAttack=1.0, PalDamageRateDefense=1.0, PlayerDamageRateAttack=1.0, PlayerDamageRateDefense=1.0, PlayerStomachDecreaceRate=1.0, PlayerStaminaDecreaceRate=1.0, PlayerAutoHPRegeneRate=1.0, PlayerAutoHpRegeneRateInSleep=1.0, PalStomachDecreaceRate=1.0, PalStaminaDecreaceRate=1.0, PalAutoHPRegeneRate=1.0, PalAutoHpRegeneRateInSleep=1.0, BuildObjectDamageRate=1.0, BuildObjectDeteriorationDamageRate=1.0, CollectionDropRate=1.0, CollectionObjectHpRate=1.0, CollectionObjectRespawnSpeedRate=1.0, EnemyDropItemRate=1.0, DeathPenalty=DropAll, bEnablePlayerToPlayerDamage=true, bEnableFriendlyFire=true, bEnableInvaderEnemy=true, bActiveUNKO=true, bEnableAimAssistPad=true, bEnableAimAssistKeyboard=true, DropItemMaxNum=100, DropItemMaxNum_UNKO=100, BaseCampMaxNum=100, BaseCampWorkerMaxNum=100, DropItemAliveMaxHours=24.0, bAutoResetGuildNoOnlinePlayers=true, AutoResetGuildTimeNoOnlinePlayers=24.0, GuildPlayerMaxNum=100, PalEggDefaultHatchingTime=24.0, WorkSpeedRate=1.0, bIsMultiplay=true, bIsPvP=true, bCanPickupOtherGuildDeathPenaltyDrop=true, bEnableNonLoginPenalty=true, bEnableFastTravel=true, bIsStartLocationSelectByMap=true, bExistPlayerAfterLogout=true, bEnableDefenseOtherGuildPlayer=true, CoopPlayerMaxNum=100, ServerPlayerMaxNum=100, ServerName=MyServer, ServerDescription=MyServerDescription, AdminPassword=MyAdminPassword, ServerPassword=MyServerPassword, PublicPort=8080, PublicIP=127.0.0.1, RCONEnabled=true, RCONPort=9000, Region=US, bUseAuth=true, BanListURL=http://example.com/banlist)`

	expectedConfig := &palworld_config.Config{
		Difficulty:                          "Normal",
		DayTimeSpeedRate:                    1.0,
		NightTimeSpeedRate:                  1.0,
		ExpRate:                             1.0,
		PalCaptureRate:                      1.0,
		PalSpawnNumRate:                     1.0,
		PalDamageRateAttack:                 1.0,
		PalDamageRateDefense:                1.0,
		PlayerDamageRateAttack:              1.0,
		PlayerDamageRateDefense:             1.0,
		PlayerStomachDecreaceRate:           1.0,
		PlayerStaminaDecreaceRate:           1.0,
		PlayerAutoHPRegeneRate:              1.0,
		PlayerAutoHpRegeneRateInSleep:       1.0,
		PalStomachDecreaceRate:              1.0,
		PalStaminaDecreaceRate:              1.0,
		PalAutoHPRegeneRate:                 1.0,
		PalAutoHpRegeneRateInSleep:          1.0,
		BuildObjectDamageRate:               1.0,
		BuildObjectDeteriorationDamageRate:  1.0,
		CollectionDropRate:                  1.0,
		CollectionObjectHpRate:              1.0,
		CollectionObjectRespawnSpeedRate:    1.0,
		EnemyDropItemRate:                   1.0,
		DeathPenalty:                        "DropAll",
		EnablePlayerToPlayerDamage:          true,
		EnableFriendlyFire:                  true,
		EnableInvaderEnemy:                  true,
		ActiveUNKO:                          true,
		EnableAimAssistPad:                  true,
		EnableAimAssistKeyboard:             true,
		DropItemMaxNum:                      100,
		DropItemMaxNum_UNKO:                 100,
		BaseCampMaxNum:                      100,
		BaseCampWorkerMaxNum:                100,
		DropItemAliveMaxHours:               24.0,
		AutoResetGuildNoOnlinePlayers:       true,
		AutoResetGuildTimeNoOnlinePlayers:   24.0,
		GuildPlayerMaxNum:                   100,
		PalEggDefaultHatchingTime:           24.0,
		WorkSpeedRate:                       1.0,
		IsMultiplay:                         true,
		IsPvP:                               true,
		CanPickupOtherGuildDeathPenaltyDrop: true,
		EnableNonLoginPenalty:               true,
		EnableFastTravel:                    true,
		IsStartLocationSelectByMap:          true,
		ExistPlayerAfterLogout:              true,
		EnableDefenseOtherGuildPlayer:       true,
		CoopPlayerMaxNum:                    100,
		ServerPlayerMaxNum:                  100,
		ServerName:                          "MyServer",
		ServerDescription:                   "MyServerDescription",
		AdminPassword:                       "MyAdminPassword",
		ServerPassword:                      "MyServerPassword",
		PublicPort:                          8080,
		PublicIP:                            "127.0.0.1",
		RCONEnabled:                         true,
		RCONPort:                            9000,
		Region:                              "US",
		UseAuth:                             true,
		BanListURL:                          "http://example.com/banlist",
	}

	config := palworld_config.Parse(configString)

	if config.Difficulty != expectedConfig.Difficulty {
		t.Errorf("Expected Difficulty to be %s, but got %s", expectedConfig.Difficulty, config.Difficulty)
	}

	if config.DayTimeSpeedRate != expectedConfig.DayTimeSpeedRate {
		t.Errorf("Expected DayTimeSpeedRate to be %f, but got %f", expectedConfig.DayTimeSpeedRate, config.DayTimeSpeedRate)
	}

	if config.NightTimeSpeedRate != expectedConfig.NightTimeSpeedRate {
		t.Errorf("Expected NightTimeSpeedRate to be %f, but got %f", expectedConfig.NightTimeSpeedRate, config.NightTimeSpeedRate)
	}

	if config.ExpRate != expectedConfig.ExpRate {
		t.Errorf("Expected ExpRate to be %f, but got %f", expectedConfig.ExpRate, config.ExpRate)
	}

	if config.PalCaptureRate != expectedConfig.PalCaptureRate {
		t.Errorf("Expected PalCaptureRate to be %f, but got %f", expectedConfig.PalCaptureRate, config.PalCaptureRate)
	}

	if config.PalSpawnNumRate != expectedConfig.PalSpawnNumRate {
		t.Errorf("Expected PalSpawnNumRate to be %f, but got %f", expectedConfig.PalSpawnNumRate, config.PalSpawnNumRate)
	}

	if config.PalDamageRateAttack != expectedConfig.PalDamageRateAttack {
		t.Errorf("Expected PalDamageRateAttack to be %f, but got %f", expectedConfig.PalDamageRateAttack, config.PalDamageRateAttack)
	}

	if config.PalDamageRateDefense != expectedConfig.PalDamageRateDefense {
		t.Errorf("Expected PalDamageRateDefense to be %f, but got %f", expectedConfig.PalDamageRateDefense, config.PalDamageRateDefense)
	}

	if config.PlayerDamageRateAttack != expectedConfig.PlayerDamageRateAttack {
		t.Errorf("Expected PlayerDamageRateAttack to be %f, but got %f", expectedConfig.PlayerDamageRateAttack, config.PlayerDamageRateAttack)
	}

	if config.PlayerDamageRateDefense != expectedConfig.PlayerDamageRateDefense {
		t.Errorf("Expected PlayerDamageRateDefense to be %f, but got %f", expectedConfig.PlayerDamageRateDefense, config.PlayerDamageRateDefense)
	}

	if config.PlayerStomachDecreaceRate != expectedConfig.PlayerStomachDecreaceRate {
		t.Errorf("Expected PlayerStomachDecreaceRate to be %f, but got %f", expectedConfig.PlayerStomachDecreaceRate, config.PlayerStomachDecreaceRate)
	}

	if config.PlayerStaminaDecreaceRate != expectedConfig.PlayerStaminaDecreaceRate {
		t.Errorf("Expected PlayerStaminaDecreaceRate to be %f, but got %f", expectedConfig.PlayerStaminaDecreaceRate, config.PlayerStaminaDecreaceRate)
	}

	if config.PlayerAutoHPRegeneRate != expectedConfig.PlayerAutoHPRegeneRate {
		t.Errorf("Expected PlayerAutoHPRegeneRate to be %f, but got %f", expectedConfig.PlayerAutoHPRegeneRate, config.PlayerAutoHPRegeneRate)
	}

	if config.PlayerAutoHpRegeneRateInSleep != expectedConfig.PlayerAutoHpRegeneRateInSleep {
		t.Errorf("Expected PlayerAutoHpRegeneRateInSleep to be %f, but got %f", expectedConfig.PlayerAutoHpRegeneRateInSleep, config.PlayerAutoHpRegeneRateInSleep)
	}

	if config.PalStomachDecreaceRate != expectedConfig.PalStomachDecreaceRate {
		t.Errorf("Expected PalStomachDecreaceRate to be %f, but got %f", expectedConfig.PalStomachDecreaceRate, config.PalStomachDecreaceRate)
	}

	if config.PalStaminaDecreaceRate != expectedConfig.PalStaminaDecreaceRate {
		t.Errorf("Expected PalStaminaDecreaceRate to be %f, but got %f", expectedConfig.PalStaminaDecreaceRate, config.PalStaminaDecreaceRate)
	}

	if config.PalAutoHPRegeneRate != expectedConfig.PalAutoHPRegeneRate {
		t.Errorf("Expected PalAutoHPRegeneRate to be %f, but got %f", expectedConfig.PalAutoHPRegeneRate, config.PalAutoHPRegeneRate)
	}

	if config.PalAutoHpRegeneRateInSleep != expectedConfig.PalAutoHpRegeneRateInSleep {
		t.Errorf("Expected PalAutoHpRegeneRateInSleep to be %f, but got %f", expectedConfig.PalAutoHpRegeneRateInSleep, config.PalAutoHpRegeneRateInSleep)
	}

	if config.BuildObjectDamageRate != expectedConfig.BuildObjectDamageRate {
		t.Errorf("Expected BuildObjectDamageRate to be %f, but got %f", expectedConfig.BuildObjectDamageRate, config.BuildObjectDamageRate)
	}

	if config.BuildObjectDeteriorationDamageRate != expectedConfig.BuildObjectDeteriorationDamageRate {
		t.Errorf("Expected BuildObjectDeteriorationDamageRate to be %f, but got %f", expectedConfig.BuildObjectDeteriorationDamageRate, config.BuildObjectDeteriorationDamageRate)
	}

	if config.CollectionDropRate != expectedConfig.CollectionDropRate {
		t.Errorf("Expected CollectionDropRate to be %f, but got %f", expectedConfig.CollectionDropRate, config.CollectionDropRate)
	}

	if config.CollectionObjectHpRate != expectedConfig.CollectionObjectHpRate {
		t.Errorf("Expected CollectionObjectHpRate to be %f, but got %f", expectedConfig.CollectionObjectHpRate, config.CollectionObjectHpRate)
	}

	if config.CollectionObjectRespawnSpeedRate != expectedConfig.CollectionObjectRespawnSpeedRate {
		t.Errorf("Expected CollectionObjectRespawnSpeedRate to be %f, but got %f", expectedConfig.CollectionObjectRespawnSpeedRate, config.CollectionObjectRespawnSpeedRate)
	}

	if config.EnemyDropItemRate != expectedConfig.EnemyDropItemRate {
		t.Errorf("Expected EnemyDropItemRate to be %f, but got %f", expectedConfig.EnemyDropItemRate, config.EnemyDropItemRate)
	}

	if config.DeathPenalty != expectedConfig.DeathPenalty {
		t.Errorf("Expected DeathPenalty to be %s, but got %s", expectedConfig.DeathPenalty, config.DeathPenalty)
	}

	if config.EnablePlayerToPlayerDamage != expectedConfig.EnablePlayerToPlayerDamage {
		t.Errorf("Expected EnablePlayerToPlayerDamage to be %t, but got %t", expectedConfig.EnablePlayerToPlayerDamage, config.EnablePlayerToPlayerDamage)
	}

	if config.EnableFriendlyFire != expectedConfig.EnableFriendlyFire {
		t.Errorf("Expected EnableFriendlyFire to be %t, but got %t", expectedConfig.EnableFriendlyFire, config.EnableFriendlyFire)
	}

	if config.EnableInvaderEnemy != expectedConfig.EnableInvaderEnemy {
		t.Errorf("Expected EnableInvaderEnemy to be %t, but got %t", expectedConfig.EnableInvaderEnemy, config.EnableInvaderEnemy)
	}

	if config.ActiveUNKO != expectedConfig.ActiveUNKO {
		t.Errorf("Expected ActiveUNKO to be %t, but got %t", expectedConfig.ActiveUNKO, config.ActiveUNKO)
	}

	if config.EnableAimAssistPad != expectedConfig.EnableAimAssistPad {
		t.Errorf("Expected EnableAimAssistPad to be %t, but got %t", expectedConfig.EnableAimAssistPad, config.EnableAimAssistPad)
	}

	if config.EnableAimAssistKeyboard != expectedConfig.EnableAimAssistKeyboard {
		t.Errorf("Expected EnableAimAssistKeyboard to be %t, but got %t", expectedConfig.EnableAimAssistKeyboard, config.EnableAimAssistKeyboard)
	}

	if config.DropItemMaxNum != expectedConfig.DropItemMaxNum {
		t.Errorf("Expected DropItemMaxNum to be %d, but got %d", expectedConfig.DropItemMaxNum, config.DropItemMaxNum)
	}

	if config.DropItemMaxNum_UNKO != expectedConfig.DropItemMaxNum_UNKO {
		t.Errorf("Expected DropItemMaxNum_UNKO to be %d, but got %d", expectedConfig.DropItemMaxNum_UNKO, config.DropItemMaxNum_UNKO)
	}

	if config.BaseCampMaxNum != expectedConfig.BaseCampMaxNum {
		t.Errorf("Expected BaseCampMaxNum to be %d, but got %d", expectedConfig.BaseCampMaxNum, config.BaseCampMaxNum)
	}

	if config.BaseCampWorkerMaxNum != expectedConfig.BaseCampWorkerMaxNum {
		t.Errorf("Expected BaseCampWorkerMaxNum to be %d, but got %d", expectedConfig.BaseCampWorkerMaxNum, config.BaseCampWorkerMaxNum)
	}

	if config.DropItemAliveMaxHours != expectedConfig.DropItemAliveMaxHours {
		t.Errorf("Expected DropItemAliveMaxHours to be %f, but got %f", expectedConfig.DropItemAliveMaxHours, config.DropItemAliveMaxHours)
	}

	if config.AutoResetGuildNoOnlinePlayers != expectedConfig.AutoResetGuildNoOnlinePlayers {
		t.Errorf("Expected AutoResetGuildNoOnlinePlayers to be %t, but got %t", expectedConfig.AutoResetGuildNoOnlinePlayers, config.AutoResetGuildNoOnlinePlayers)
	}

	if config.AutoResetGuildTimeNoOnlinePlayers != expectedConfig.AutoResetGuildTimeNoOnlinePlayers {
		t.Errorf("Expected AutoResetGuildTimeNoOnlinePlayers to be %f, but got %f", expectedConfig.AutoResetGuildTimeNoOnlinePlayers, config.AutoResetGuildTimeNoOnlinePlayers)
	}

	if config.GuildPlayerMaxNum != expectedConfig.GuildPlayerMaxNum {
		t.Errorf("Expected GuildPlayerMaxNum to be %d, but got %d", expectedConfig.GuildPlayerMaxNum, config.GuildPlayerMaxNum)
	}

	if config.PalEggDefaultHatchingTime != expectedConfig.PalEggDefaultHatchingTime {
		t.Errorf("Expected PalEggDefaultHatchingTime to be %f, but got %f", expectedConfig.PalEggDefaultHatchingTime, config.PalEggDefaultHatchingTime)
	}

	if config.WorkSpeedRate != expectedConfig.WorkSpeedRate {
		t.Errorf("Expected WorkSpeedRate to be %f, but got %f", expectedConfig.WorkSpeedRate, config.WorkSpeedRate)
	}

	if config.IsMultiplay != expectedConfig.IsMultiplay {
		t.Errorf("Expected IsMultiplay to be %t, but got %t", expectedConfig.IsMultiplay, config.IsMultiplay)
	}

	if config.IsPvP != expectedConfig.IsPvP {
		t.Errorf("Expected IsPvP to be %t, but got %t", expectedConfig.IsPvP, config.IsPvP)
	}

	if config.CanPickupOtherGuildDeathPenaltyDrop != expectedConfig.CanPickupOtherGuildDeathPenaltyDrop {
		t.Errorf("Expected CanPickupOtherGuildDeathPenaltyDrop to be %t, but got %t", expectedConfig.CanPickupOtherGuildDeathPenaltyDrop, config.CanPickupOtherGuildDeathPenaltyDrop)
	}

	if config.EnableNonLoginPenalty != expectedConfig.EnableNonLoginPenalty {
		t.Errorf("Expected EnableNonLoginPenalty to be %t, but got %t", expectedConfig.EnableNonLoginPenalty, config.EnableNonLoginPenalty)
	}

	if config.EnableFastTravel != expectedConfig.EnableFastTravel {
		t.Errorf("Expected EnableFastTravel to be %t, but got %t", expectedConfig.EnableFastTravel, config.EnableFastTravel)
	}

	if config.IsStartLocationSelectByMap != expectedConfig.IsStartLocationSelectByMap {
		t.Errorf("Expected IsStartLocationSelectByMap to be %t, but got %t", expectedConfig.IsStartLocationSelectByMap, config.IsStartLocationSelectByMap)
	}

	if config.ExistPlayerAfterLogout != expectedConfig.ExistPlayerAfterLogout {
		t.Errorf("Expected ExistPlayerAfterLogout to be %t, but got %t", expectedConfig.ExistPlayerAfterLogout, config.ExistPlayerAfterLogout)
	}

	if config.EnableDefenseOtherGuildPlayer != expectedConfig.EnableDefenseOtherGuildPlayer {
		t.Errorf("Expected EnableDefenseOtherGuildPlayer to be %t, but got %t", expectedConfig.EnableDefenseOtherGuildPlayer, config.EnableDefenseOtherGuildPlayer)
	}

	if config.CoopPlayerMaxNum != expectedConfig.CoopPlayerMaxNum {
		t.Errorf("Expected CoopPlayerMaxNum to be %d, but got %d", expectedConfig.CoopPlayerMaxNum, config.CoopPlayerMaxNum)
	}

	if config.ServerPlayerMaxNum != expectedConfig.ServerPlayerMaxNum {
		t.Errorf("Expected ServerPlayerMaxNum to be %d, but got %d", expectedConfig.ServerPlayerMaxNum, config.ServerPlayerMaxNum)
	}

	if config.ServerName != expectedConfig.ServerName {
		t.Errorf("Expected ServerName to be %s, but got %s", expectedConfig.ServerName, config.ServerName)
	}

	if config.ServerDescription != expectedConfig.ServerDescription {
		t.Errorf("Expected ServerDescription to be %s, but got %s", expectedConfig.ServerDescription, config.ServerDescription)
	}

	if config.AdminPassword != expectedConfig.AdminPassword {
		t.Errorf("Expected AdminPassword to be %s, but got %s", expectedConfig.AdminPassword, config.AdminPassword)
	}

	if config.ServerPassword != expectedConfig.ServerPassword {
		t.Errorf("Expected ServerPassword to be %s, but got %s", expectedConfig.ServerPassword, config.ServerPassword)
	}

	if config.PublicPort != expectedConfig.PublicPort {
		t.Errorf("Expected PublicPort to be %d, but got %d", expectedConfig.PublicPort, config.PublicPort)
	}

	if config.PublicIP != expectedConfig.PublicIP {
		t.Errorf("Expected PublicIP to be %s, but got %s", expectedConfig.PublicIP, config.PublicIP)
	}

	if config.RCONEnabled != expectedConfig.RCONEnabled {
		t.Errorf("Expected RCONEnabled to be %t, but got %t", expectedConfig.RCONEnabled, config.RCONEnabled)
	}

	if config.RCONPort != expectedConfig.RCONPort {
		t.Errorf("Expected RCONPort to be %d, but got %d", expectedConfig.RCONPort, config.RCONPort)
	}

	if config.Region != expectedConfig.Region {
		t.Errorf("Expected Region to be %s, but got %s", expectedConfig.Region, config.Region)
	}

	if config.UseAuth != expectedConfig.UseAuth {
		t.Errorf("Expected UseAuth to be %t, but got %t", expectedConfig.UseAuth, config.UseAuth)
	}

	if config.BanListURL != expectedConfig.BanListURL {
		t.Errorf("Expected BanListURL to be %s, but got %s", expectedConfig.BanListURL, config.BanListURL)
	}
}
