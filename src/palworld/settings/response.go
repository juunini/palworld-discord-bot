package settings

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/i18n"
)

func Response(s *discordgo.Session, channelID, command string) {
	after, found := strings.CutPrefix(command, config.DISCORD_COMMAND_ALIAS_SERVER_SETTINGS+" ")
	if !found || strings.TrimSpace(after) == "" {
		messages := i18n.SettingsHelp(config.DISCORD_COMMAND_PREFIX, config.DISCORD_COMMAND_ALIAS_SERVER_SETTINGS)

		fields := []*discordgo.MessageEmbedField{}

		for _, message := range messages {
			fields = append(fields, &discordgo.MessageEmbedField{
				Value:  message,
				Inline: false,
			})
		}

		embed := &discordgo.MessageEmbed{
			Fields: fields,
		}

		s.ChannelMessageSendEmbed(channelID, embed)

		return
	}

	if strings.HasPrefix(after, "DayTimeSpeedRate ") {
		value, _ := strings.CutPrefix(after, "DayTimeSpeedRate ")
		s.ChannelMessageSend(channelID, setDayTimeSpeedRate(s, value))
		return
	} else if strings.HasPrefix(after, "NightTimeSpeedRate ") {
		value, _ := strings.CutPrefix(after, "NightTimeSpeedRate ")
		s.ChannelMessageSend(channelID, setNightTimeSpeedRate(s, value))
		return
	} else if strings.HasPrefix(after, "ExpRate ") {
		value, _ := strings.CutPrefix(after, "ExpRate ")
		s.ChannelMessageSend(channelID, setExpRate(s, value))
		return
	} else if strings.HasPrefix(after, "PalCaptureRate ") {
		value, _ := strings.CutPrefix(after, "PalCaptureRate ")
		s.ChannelMessageSend(channelID, setPalCaptureRate(s, value))
		return
	} else if strings.HasPrefix(after, "PalSpawnNumRate ") {
		value, _ := strings.CutPrefix(after, "PalSpawnNumRate ")
		s.ChannelMessageSend(channelID, setPalSpawnNumRate(s, value))
		return
	} else if strings.HasPrefix(after, "PalDamageRateAttack ") {
		value, _ := strings.CutPrefix(after, "PalDamageRateAttack ")
		s.ChannelMessageSend(channelID, setPalDamageRateAttack(s, value))
		return
	} else if strings.HasPrefix(after, "PalDamageRateDefense ") {
		value, _ := strings.CutPrefix(after, "PalDamageRateDefense ")
		s.ChannelMessageSend(channelID, setPalDamageRateDefense(s, value))
		return
	} else if strings.HasPrefix(after, "PlayerDamageRateAttack ") {
		value, _ := strings.CutPrefix(after, "PlayerDamageRateAttack ")
		s.ChannelMessageSend(channelID, setPlayerDamageRateAttack(s, value))
		return
	} else if strings.HasPrefix(after, "PlayerDamageRateDefense ") {
		value, _ := strings.CutPrefix(after, "PlayerDamageRateDefense ")
		s.ChannelMessageSend(channelID, setPlayerDamageRateDefense(s, value))
		return
	} else if strings.HasPrefix(after, "PlayerStomachDecreaceRate ") {
		value, _ := strings.CutPrefix(after, "PlayerStomachDecreaceRate ")
		s.ChannelMessageSend(channelID, setPlayerStomachDecreaceRate(s, value))
		return
	} else if strings.HasPrefix(after, "PlayerStaminaDecreaceRate ") {
		value, _ := strings.CutPrefix(after, "PlayerStaminaDecreaceRate ")
		s.ChannelMessageSend(channelID, setPlayerStaminaDecreaceRate(s, value))
		return
	} else if strings.HasPrefix(after, "PlayerAutoHPRegeneRate ") {
		value, _ := strings.CutPrefix(after, "PlayerAutoHPRegeneRate ")
		s.ChannelMessageSend(channelID, setPlayerAutoHPRegeneRate(s, value))
		return
	} else if strings.HasPrefix(after, "PlayerAutoHpRegeneRateInSleep ") {
		value, _ := strings.CutPrefix(after, "PlayerAutoHpRegeneRateInSleep ")
		s.ChannelMessageSend(channelID, setPlayerAutoHpRegeneRateInSleep(s, value))
		return
	} else if strings.HasPrefix(after, "PalStomachDecreaceRate ") {
		value, _ := strings.CutPrefix(after, "PalStomachDecreaceRate ")
		s.ChannelMessageSend(channelID, setPalStomachDecreaceRate(s, value))
		return
	} else if strings.HasPrefix(after, "PalStaminaDecreaceRate ") {
		value, _ := strings.CutPrefix(after, "PalStaminaDecreaceRate ")
		s.ChannelMessageSend(channelID, setPalStaminaDecreaceRate(s, value))
		return
	} else if strings.HasPrefix(after, "PalAutoHPRegeneRate ") {
		value, _ := strings.CutPrefix(after, "PalAutoHPRegeneRate ")
		s.ChannelMessageSend(channelID, setPalAutoHPRegeneRate(s, value))
		return
	} else if strings.HasPrefix(after, "PalAutoHpRegeneRateInSleep ") {
		value, _ := strings.CutPrefix(after, "PalAutoHpRegeneRateInSleep ")
		s.ChannelMessageSend(channelID, setPalAutoHpRegeneRateInSleep(s, value))
		return
	} else if strings.HasPrefix(after, "BuildObjectDamageRate ") {
		value, _ := strings.CutPrefix(after, "BuildObjectDamageRate ")
		s.ChannelMessageSend(channelID, setBuildObjectDamageRate(s, value))
		return
	} else if strings.HasPrefix(after, "BuildObjectDeteriorationDamageRate ") {
		value, _ := strings.CutPrefix(after, "BuildObjectDeteriorationDamageRate ")
		s.ChannelMessageSend(channelID, setBuildObjectDeteriorationDamageRate(s, value))
		return
	} else if strings.HasPrefix(after, "CollectionDropRate ") {
		value, _ := strings.CutPrefix(after, "CollectionDropRate ")
		s.ChannelMessageSend(channelID, setCollectionDropRate(s, value))
		return
	} else if strings.HasPrefix(after, "CollectionObjectHpRate ") {
		value, _ := strings.CutPrefix(after, "CollectionObjectHpRate ")
		s.ChannelMessageSend(channelID, setCollectionObjectHpRate(s, value))
		return
	} else if strings.HasPrefix(after, "CollectionObjectRespawnSpeedRate ") {
		value, _ := strings.CutPrefix(after, "CollectionObjectRespawnSpeedRate ")
		s.ChannelMessageSend(channelID, setCollectionObjectRespawnSpeedRate(s, value))
		return
	} else if strings.HasPrefix(after, "EnemyDropItemRate ") {
		value, _ := strings.CutPrefix(after, "EnemyDropItemRate ")
		s.ChannelMessageSend(channelID, setEnemyDropItemRate(s, value))
		return
	} else if strings.HasPrefix(after, "DeathPenalty ") {
		value, _ := strings.CutPrefix(after, "DeathPenalty ")
		s.ChannelMessageSend(channelID, setDeathPenalty(s, value))
		return
	} else if strings.HasPrefix(after, "EnablePlayerToPlayerDamage ") {
		value, _ := strings.CutPrefix(after, "EnablePlayerToPlayerDamage ")
		s.ChannelMessageSend(channelID, setEnablePlayerToPlayerDamage(s, value))
		return
	} else if strings.HasPrefix(after, "EnableFriendlyFire ") {
		value, _ := strings.CutPrefix(after, "EnableFriendlyFire ")
		s.ChannelMessageSend(channelID, setEnableFriendlyFire(s, value))
		return
	} else if strings.HasPrefix(after, "EnableInvaderEnemy ") {
		value, _ := strings.CutPrefix(after, "EnableInvaderEnemy ")
		s.ChannelMessageSend(channelID, setEnableInvaderEnemy(s, value))
		return
	} else if strings.HasPrefix(after, "ActiveUNKO ") {
		value, _ := strings.CutPrefix(after, "ActiveUNKO ")
		s.ChannelMessageSend(channelID, setActiveUNKO(s, value))
		return
	} else if strings.HasPrefix(after, "EnableAimAssistPad ") {
		value, _ := strings.CutPrefix(after, "EnableAimAssistPad ")
		s.ChannelMessageSend(channelID, setEnableAimAssistPad(s, value))
		return
	} else if strings.HasPrefix(after, "EnableAimAssistKeyboard ") {
		value, _ := strings.CutPrefix(after, "EnableAimAssistKeyboard ")
		s.ChannelMessageSend(channelID, setEnableAimAssistKeyboard(s, value))
		return
	} else if strings.HasPrefix(after, "DropItemMaxNum ") {
		value, _ := strings.CutPrefix(after, "DropItemMaxNum ")
		s.ChannelMessageSend(channelID, setDropItemMaxNum(s, value))
		return
	} else if strings.HasPrefix(after, "DropItemMaxNum_UNKO ") {
		value, _ := strings.CutPrefix(after, "DropItemMaxNum_UNKO ")
		s.ChannelMessageSend(channelID, setDropItemMaxNum_UNKO(s, value))
		return
	} else if strings.HasPrefix(after, "BaseCampMaxNum ") {
		value, _ := strings.CutPrefix(after, "BaseCampMaxNum ")
		s.ChannelMessageSend(channelID, setBaseCampMaxNum(s, value))
		return
	} else if strings.HasPrefix(after, "BaseCampWorkerMaxNum ") {
		value, _ := strings.CutPrefix(after, "BaseCampWorkerMaxNum ")
		s.ChannelMessageSend(channelID, setBaseCampWorkerMaxNum(s, value))
		return
	} else if strings.HasPrefix(after, "DropItemAliveMaxHours ") {
		value, _ := strings.CutPrefix(after, "DropItemAliveMaxHours ")
		s.ChannelMessageSend(channelID, setDropItemAliveMaxHours(s, value))
		return
	} else if strings.HasPrefix(after, "AutoResetGuildNoOnlinePlayers ") {
		value, _ := strings.CutPrefix(after, "AutoResetGuildNoOnlinePlayers ")
		s.ChannelMessageSend(channelID, setAutoResetGuildNoOnlinePlayers(s, value))
		return
	} else if strings.HasPrefix(after, "AutoResetGuildTimeNoOnlinePlayers ") {
		value, _ := strings.CutPrefix(after, "AutoResetGuildTimeNoOnlinePlayers ")
		s.ChannelMessageSend(channelID, setAutoResetGuildTimeNoOnlinePlayers(s, value))
		return
	} else if strings.HasPrefix(after, "GuildPlayerMaxNum ") {
		value, _ := strings.CutPrefix(after, "GuildPlayerMaxNum ")
		s.ChannelMessageSend(channelID, setGuildPlayerMaxNum(s, value))
		return
	} else if strings.HasPrefix(after, "PalEggDefaultHatchingTime ") {
		value, _ := strings.CutPrefix(after, "PalEggDefaultHatchingTime ")
		s.ChannelMessageSend(channelID, setPalEggDefaultHatchingTime(s, value))
		return
	} else if strings.HasPrefix(after, "WorkSpeedRate ") {
		value, _ := strings.CutPrefix(after, "WorkSpeedRate ")
		s.ChannelMessageSend(channelID, setWorkSpeedRate(s, value))
		return
	} else if strings.HasPrefix(after, "IsMultiplay ") {
		value, _ := strings.CutPrefix(after, "IsMultiplay ")
		s.ChannelMessageSend(channelID, setIsMultiplay(s, value))
		return
	} else if strings.HasPrefix(after, "IsPvP ") {
		value, _ := strings.CutPrefix(after, "IsPvP ")
		s.ChannelMessageSend(channelID, setIsPvP(s, value))
		return
	} else if strings.HasPrefix(after, "CanPickupOtherGuildDeathPenaltyDrop ") {
		value, _ := strings.CutPrefix(after, "CanPickupOtherGuildDeathPenaltyDrop ")
		s.ChannelMessageSend(channelID, setCanPickupOtherGuildDeathPenaltyDrop(s, value))
		return
	} else if strings.HasPrefix(after, "EnableNonLoginPenalty ") {
		value, _ := strings.CutPrefix(after, "EnableNonLoginPenalty ")
		s.ChannelMessageSend(channelID, setEnableNonLoginPenalty(s, value))
		return
	} else if strings.HasPrefix(after, "EnableFastTravel ") {
		value, _ := strings.CutPrefix(after, "EnableFastTravel ")
		s.ChannelMessageSend(channelID, setEnableFastTravel(s, value))
		return
	} else if strings.HasPrefix(after, "IsStartLocationSelectByMap ") {
		value, _ := strings.CutPrefix(after, "IsStartLocationSelectByMap ")
		s.ChannelMessageSend(channelID, setIsStartLocationSelectByMap(s, value))
		return
	} else if strings.HasPrefix(after, "ExistPlayerAfterLogout ") {
		value, _ := strings.CutPrefix(after, "ExistPlayerAfterLogout ")
		s.ChannelMessageSend(channelID, setExistPlayerAfterLogout(s, value))
		return
	} else if strings.HasPrefix(after, "EnableDefenseOtherGuildPlayer ") {
		value, _ := strings.CutPrefix(after, "EnableDefenseOtherGuildPlayer ")
		s.ChannelMessageSend(channelID, setEnableDefenseOtherGuildPlayer(s, value))
		return
	} else if strings.HasPrefix(after, "CoopPlayerMaxNum ") {
		value, _ := strings.CutPrefix(after, "CoopPlayerMaxNum ")
		s.ChannelMessageSend(channelID, setCoopPlayerMaxNum(s, value))
		return
	} else if strings.HasPrefix(after, "ServerPlayerMaxNum ") {
		value, _ := strings.CutPrefix(after, "ServerPlayerMaxNum ")
		s.ChannelMessageSend(channelID, setServerPlayerMaxNum(s, value))
		return
	} else if strings.HasPrefix(after, "ServerName ") {
		value, _ := strings.CutPrefix(after, "ServerName ")
		s.ChannelMessageSend(channelID, setServerName(s, value))
		return
	} else if strings.HasPrefix(after, "ServerDescription ") {
		value, _ := strings.CutPrefix(after, "ServerDescription ")
		s.ChannelMessageSend(channelID, setServerDescription(s, value))
		return
	} else if strings.HasPrefix(after, "AdminPassword ") {
		value, _ := strings.CutPrefix(after, "AdminPassword ")
		s.ChannelMessageSend(channelID, setAdminPassword(s, value))
		return
	} else if strings.HasPrefix(after, "ServerPassword ") {
		value, _ := strings.CutPrefix(after, "ServerPassword ")
		s.ChannelMessageSend(channelID, setServerPassword(s, value))
		return
	} else if strings.HasPrefix(after, "PublicPort ") {
		value, _ := strings.CutPrefix(after, "PublicPort ")
		s.ChannelMessageSend(channelID, setPublicPort(s, value))
		return
	} else if strings.HasPrefix(after, "PublicIP ") {
		value, _ := strings.CutPrefix(after, "PublicIP ")
		s.ChannelMessageSend(channelID, setPublicIP(s, value))
		return
	} else if strings.HasPrefix(after, "RCONEnabled ") {
		value, _ := strings.CutPrefix(after, "RCONEnabled ")
		s.ChannelMessageSend(channelID, setRCONEnabled(s, value))
		return
	} else if strings.HasPrefix(after, "RCONPort ") {
		value, _ := strings.CutPrefix(after, "RCONPort ")
		s.ChannelMessageSend(channelID, setRCONPort(s, value))
		return
	} else if strings.HasPrefix(after, "Region ") {
		value, _ := strings.CutPrefix(after, "Region ")
		s.ChannelMessageSend(channelID, setRegion(s, value))
		return
	} else if strings.HasPrefix(after, "UseAuth ") {
		value, _ := strings.CutPrefix(after, "UseAuth ")
		s.ChannelMessageSend(channelID, setUseAuth(s, value))
		return
	} else if strings.HasPrefix(after, "BanListURL ") {
		value, _ := strings.CutPrefix(after, "BanListURL ")
		s.ChannelMessageSend(channelID, setBanListURL(s, value))
		return
	}

	s.ChannelMessageSend(channelID, "Unknown settings command")
}
