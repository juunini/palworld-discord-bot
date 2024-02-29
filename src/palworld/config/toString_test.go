package palworld_config_test

import (
	"testing"

	palworld_config "github.com/juunini/palworld-discord-bot/src/palworld/config"
)

func TestConfigToString(t *testing.T) {
	p := &palworld_config.Config{
		Difficulty:                          "easy",
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
		DeathPenalty:                        "drop",
		EnablePlayerToPlayerDamage:          false,
		EnableFriendlyFire:                  true,
		EnableInvaderEnemy:                  true,
		ActiveUNKO:                          true,
		EnableAimAssistPad:                  true,
		EnableAimAssistKeyboard:             true,
		DropItemMaxNum:                      10,
		DropItemMaxNum_UNKO:                 10,
		BaseCampMaxNum:                      10,
		BaseCampWorkerMaxNum:                10,
		DropItemAliveMaxHours:               1.0,
		AutoResetGuildNoOnlinePlayers:       true,
		AutoResetGuildTimeNoOnlinePlayers:   1.0,
		GuildPlayerMaxNum:                   10,
		PalEggDefaultHatchingTime:           1.0,
		WorkSpeedRate:                       1.0,
		IsMultiplay:                         true,
		IsPvP:                               true,
		CanPickupOtherGuildDeathPenaltyDrop: true,
		EnableNonLoginPenalty:               true,
		EnableFastTravel:                    true,
		IsStartLocationSelectByMap:          true,
		ExistPlayerAfterLogout:              true,
		EnableDefenseOtherGuildPlayer:       true,
		CoopPlayerMaxNum:                    10,
		ServerPlayerMaxNum:                  10,
		ServerName:                          "MyServer",
		ServerDescription:                   "MyServerDescription",
		AdminPassword:                       "admin123",
		ServerPassword:                      "password123",
		PublicPort:                          8080,
		PublicIP:                            "192.168.0.1",
		RCONEnabled:                         true,
		RCONPort:                            1234,
		Region:                              "us-west",
		UseAuth:                             true,
		BanListURL:                          "https://example.com/banlist",
	}

	expected := `[` +
		`/Script/Pal.PalGameWorldSettings]` + "\n" +
		`OptionSettings=(` +
		`Difficulty=easy,` +
		`DayTimeSpeedRate=1.000000,` +
		`NightTimeSpeedRate=1.000000,` +
		`ExpRate=1.000000,` +
		`PalCaptureRate=1.000000,` +
		`PalSpawnNumRate=1.000000,` +
		`PalDamageRateAttack=1.000000,` +
		`PalDamageRateDefense=1.000000,` +
		`PlayerDamageRateAttack=1.000000,` +
		`PlayerDamageRateDefense=1.000000,` +
		`PlayerStomachDecreaceRate=1.000000,` +
		`PlayerStaminaDecreaceRate=1.000000,` +
		`PlayerAutoHPRegeneRate=1.000000,` +
		`PlayerAutoHpRegeneRateInSleep=1.000000,` +
		`PalStomachDecreaceRate=1.000000,` +
		`PalStaminaDecreaceRate=1.000000,` +
		`PalAutoHPRegeneRate=1.000000,` +
		`PalAutoHpRegeneRateInSleep=1.000000,` +
		`BuildObjectDamageRate=1.000000,` +
		`BuildObjectDeteriorationDamageRate=1.000000,` +
		`CollectionDropRate=1.000000,` +
		`CollectionObjectHpRate=1.000000,` +
		`CollectionObjectRespawnSpeedRate=1.000000,` +
		`EnemyDropItemRate=1.000000,` +
		`DeathPenalty=drop,` +
		`bEnablePlayerToPlayerDamage=False,` +
		`bEnableFriendlyFire=True,` +
		`bEnableInvaderEnemy=True,` +
		`bActiveUNKO=True,` +
		`bEnableAimAssistPad=True,` +
		`bEnableAimAssistKeyboard=True,` +
		`DropItemMaxNum=10,` +
		`DropItemMaxNum_UNKO=10,` +
		`BaseCampMaxNum=10,` +
		`BaseCampWorkerMaxNum=10,` +
		`DropItemAliveMaxHours=1.000000,` +
		`bAutoResetGuildNoOnlinePlayers=True,` +
		`AutoResetGuildTimeNoOnlinePlayers=1.000000,` +
		`GuildPlayerMaxNum=10,` +
		`PalEggDefaultHatchingTime=1.000000,` +
		`WorkSpeedRate=1.000000,` +
		`bIsMultiplay=True,` +
		`bIsPvP=True,` +
		`bCanPickupOtherGuildDeathPenaltyDrop=True,` +
		`bEnableNonLoginPenalty=True,` +
		`bEnableFastTravel=True,` +
		`bIsStartLocationSelectByMap=True,` +
		`bExistPlayerAfterLogout=True,` +
		`bEnableDefenseOtherGuildPlayer=True,` +
		`CoopPlayerMaxNum=10,` +
		`ServerPlayerMaxNum=10,` +
		`ServerName="MyServer",` +
		`ServerDescription="MyServerDescription",` +
		`AdminPassword="admin123",` +
		`ServerPassword="password123",` +
		`PublicPort=8080,` +
		`PublicIP="192.168.0.1",` +
		`RCONEnabled=True,` +
		`RCONPort=1234,` +
		`Region="us-west",` +
		`bUseAuth=True,` +
		`BanListURL="https://example.com/banlist"` +
		`)`

	result := p.ToString()

	if result != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, result)
	}
}
