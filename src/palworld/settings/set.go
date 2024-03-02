package settings

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
	palworld_settings "github.com/juunini/palworld-settings"
)

func setDayTimeSpeedRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.DayTimeSpeedRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.DayTimeSpeedRate, val)
}

func setNightTimeSpeedRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.NightTimeSpeedRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.NightTimeSpeedRate, val)
}

func setExpRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.ExpRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.ExpRate, val)
}

func setPalCaptureRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalCaptureRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PalCaptureRate, val)
}

func setPalSpawnNumRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalSpawnNumRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PalSpawnNumRate, val)
}

func setPalDamageRateAttack(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalDamageRateAttack = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PalDamageRateAttack, val)
}

func setPalDamageRateDefense(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalDamageRateDefense = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PalDamageRateDefense, val)
}

func setPlayerDamageRateAttack(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerDamageRateAttack = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PlayerDamageRateAttack, val)
}

func setPlayerDamageRateDefense(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerDamageRateDefense = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PlayerDamageRateDefense, val)
}

func setPlayerStomachDecreaceRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerStomachDecreaceRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PlayerStomachDecreaceRate, val)
}

func setPlayerStaminaDecreaceRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerStaminaDecreaceRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PlayerStaminaDecreaceRate, val)
}

func setPlayerAutoHPRegeneRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerAutoHPRegeneRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PlayerAutoHPRegeneRate, val)
}

func setPlayerAutoHpRegeneRateInSleep(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PlayerAutoHpRegeneRateInSleep = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PlayerAutoHpRegeneRateInSleep, val)
}

func setPalStomachDecreaceRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalStomachDecreaceRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PalStomachDecreaceRate, val)
}

func setPalStaminaDecreaceRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalStaminaDecreaceRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PalStaminaDecreaceRate, val)
}

func setPalAutoHPRegeneRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalAutoHPRegeneRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PalAutoHPRegeneRate, val)
}

func setPalAutoHpRegeneRateInSleep(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalAutoHpRegeneRateInSleep = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PalAutoHpRegeneRateInSleep, val)
}

func setBuildObjectDamageRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.BuildObjectDamageRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.BuildObjectDamageRate, val)
}

func setBuildObjectDeteriorationDamageRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.BuildObjectDeteriorationDamageRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.BuildObjectDeteriorationDamageRate, val)
}

func setCollectionDropRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.CollectionDropRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.CollectionDropRate, val)
}

func setCollectionObjectHpRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.CollectionObjectHpRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.CollectionObjectHpRate, val)
}

func setCollectionObjectRespawnSpeedRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.CollectionObjectRespawnSpeedRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.CollectionObjectRespawnSpeedRate, val)
}

func setEnemyDropItemRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnemyDropItemRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.EnemyDropItemRate, val)
}

func setDeathPenalty(s *discordgo.Session, value string) string {
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
	RenewDashboard(s)
	return fmt.Sprintf("%s: %s", i18n.DeathPenalty, val)
}

func setEnablePlayerToPlayerDamage(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnablePlayerToPlayerDamage = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.EnablePlayerToPlayerDamage, val)
}

func setEnableFriendlyFire(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableFriendlyFire = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.EnableFriendlyFire, val)
}

func setEnableInvaderEnemy(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableInvaderEnemy = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.EnableInvaderEnemy, val)
}

func setActiveUNKO(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.ActiveUNKO = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.ActiveUNKO, val)
}

func setEnableAimAssistPad(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableAimAssistPad = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.EnableAimAssistPad, val)
}

func setEnableAimAssistKeyboard(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableAimAssistKeyboard = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.EnableAimAssistKeyboard, val)
}

func setDropItemMaxNum(s *discordgo.Session, value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.DropItemMaxNum = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %d", i18n.DropItemMaxNum, val)
}

func setDropItemMaxNum_UNKO(s *discordgo.Session, value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.DropItemMaxNum_UNKO = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %d", i18n.DropItemMaxNum_UNKO, val)
}

func setBaseCampMaxNum(s *discordgo.Session, value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.BaseCampMaxNum = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %d", i18n.BaseCampMaxNum, val)
}

func setBaseCampWorkerMaxNum(s *discordgo.Session, value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.BaseCampWorkerMaxNum = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %d", i18n.BaseCampWorkerMaxNum, val)
}

func setDropItemAliveMaxHours(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.DropItemAliveMaxHours = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.DropItemAliveMaxHours, val)
}

func setAutoResetGuildNoOnlinePlayers(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.AutoResetGuildNoOnlinePlayers = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.AutoResetGuildNoOnlinePlayers, val)
}

func setAutoResetGuildTimeNoOnlinePlayers(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.AutoResetGuildTimeNoOnlinePlayers = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.AutoResetGuildTimeNoOnlinePlayers, val)
}

func setGuildPlayerMaxNum(s *discordgo.Session, value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.GuildPlayerMaxNum = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %d", i18n.GuildPlayerMaxNum, val)
}

func setPalEggDefaultHatchingTime(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PalEggDefaultHatchingTime = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.PalEggDefaultHatchingTime, val)
}

func setWorkSpeedRate(s *discordgo.Session, value string) string {
	val, err := utils.ToFloat64(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.WorkSpeedRate = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %f", i18n.WorkSpeedRate, val)
}

func setIsMultiplay(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.IsMultiplay = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.IsMultiplay, val)
}

func setIsPvP(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.IsPvP = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.IsPvP, val)
}

func setCanPickupOtherGuildDeathPenaltyDrop(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.CanPickupOtherGuildDeathPenaltyDrop = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.CanPickupOtherGuildDeathPenaltyDrop, val)
}

func setEnableNonLoginPenalty(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableNonLoginPenalty = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.EnableNonLoginPenalty, val)
}

func setEnableFastTravel(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableFastTravel = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.EnableFastTravel, val)
}

func setIsStartLocationSelectByMap(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.IsStartLocationSelectByMap = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.IsStartLocationSelectByMap, val)
}

func setExistPlayerAfterLogout(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.ExistPlayerAfterLogout = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.ExistPlayerAfterLogout, val)
}

func setEnableDefenseOtherGuildPlayer(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.EnableDefenseOtherGuildPlayer = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.EnableDefenseOtherGuildPlayer, val)
}

func setCoopPlayerMaxNum(s *discordgo.Session, value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.CoopPlayerMaxNum = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %d", i18n.CoopPlayerMaxNum, val)
}

func setServerPlayerMaxNum(s *discordgo.Session, value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.ServerPlayerMaxNum = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %d", i18n.ServerPlayerMaxNum, val)
}

func setServerName(s *discordgo.Session, value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.ServerName = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %s", i18n.ServerName, val)
}

func setServerDescription(s *discordgo.Session, value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.ServerDescription = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %s", i18n.ServerDescription, val)
}

func setAdminPassword(s *discordgo.Session, value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.AdminPassword = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %s", i18n.AdminPassword, val)
}

func setServerPassword(s *discordgo.Session, value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.ServerPassword = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %s", i18n.ServerPassword, val)
}

func setPublicPort(s *discordgo.Session, value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.PublicPort = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %d", i18n.PublicPort, val)
}

func setPublicIP(s *discordgo.Session, value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.PublicIP = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %s", i18n.PublicIP, val)
}

func setRCONEnabled(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.RCONEnabled = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.RCONEnabled, val)
}

func setRCONPort(s *discordgo.Session, value string) string {
	val, err := utils.ToInt(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.RCONPort = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %d", i18n.RCONPort, val)
}

func setRegion(s *discordgo.Session, value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.Region = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %s", i18n.Region, val)
}

func setUseAuth(s *discordgo.Session, value string) string {
	val, err := utils.ToBool(value)
	if err != nil {
		return err.Error()
	}

	config.PALWORLD_SERVER_SETTINGS.UseAuth = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %t", i18n.UseAuth, val)
}

func setBanListURL(s *discordgo.Session, value string) string {
	val := strings.TrimSpace(value)

	config.PALWORLD_SERVER_SETTINGS.BanListURL = val
	saveSettings()
	RenewDashboard(s)
	return fmt.Sprintf("%s: %s", i18n.BanListURL, val)
}

func saveSettings() {
	settingsString := config.PALWORLD_SERVER_SETTINGS.ToString()
	os.WriteFile(config.PALWORLD_SERVER_SETTINGS_FILE_PATH, []byte(settingsString), 0644)
}

func RenewDashboard(s *discordgo.Session) {
	if config.DISCORD_DASHBOARD_CHANNEL_ID == "" {
		return
	}

	messageID := config.DISCORD_DASHBOARD_PALWORLD_SETTINGS_MESSAGE_ID
	if messageID == "" {
		message, err := s.ChannelMessageSendEmbed(config.DISCORD_DASHBOARD_CHANNEL_ID, DashboardContent())
		if err != nil {
			console_decoration.PrintError(err.Error())
			s.ChannelMessageSend(config.DISCORD_LOG_CHANNEL_ID, err.Error())
			return
		}

		config.DISCORD_DASHBOARD_PALWORLD_SETTINGS_MESSAGE_ID = message.ID
		config.Save()

		return
	}

	_, err := s.ChannelMessageEditEmbed(config.DISCORD_DASHBOARD_CHANNEL_ID, messageID, DashboardContent())
	if err != nil {
		console_decoration.PrintError(err.Error())
		s.ChannelMessageSend(config.DISCORD_LOG_CHANNEL_ID, err.Error())
		return
	}
}

func DashboardContent() *discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{}

	value := fmt.Sprintf("- %s: %v\n", i18n.DayTimeSpeedRate, config.PALWORLD_SERVER_SETTINGS.DayTimeSpeedRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.NightTimeSpeedRate, config.PALWORLD_SERVER_SETTINGS.NightTimeSpeedRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.ExpRate, config.PALWORLD_SERVER_SETTINGS.ExpRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.PalCaptureRate, config.PALWORLD_SERVER_SETTINGS.PalCaptureRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.PalSpawnNumRate, config.PALWORLD_SERVER_SETTINGS.PalSpawnNumRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.PalDamageRateAttack, config.PALWORLD_SERVER_SETTINGS.PalDamageRateAttack)
	value += fmt.Sprintf("- %s: %v\n", i18n.PalDamageRateDefense, config.PALWORLD_SERVER_SETTINGS.PalDamageRateDefense)
	value += fmt.Sprintf("- %s: %v\n", i18n.PlayerDamageRateAttack, config.PALWORLD_SERVER_SETTINGS.PlayerDamageRateAttack)
	value += fmt.Sprintf("- %s: %v\n", i18n.PlayerDamageRateDefense, config.PALWORLD_SERVER_SETTINGS.PlayerDamageRateDefense)
	value += fmt.Sprintf("- %s: %v", i18n.PlayerStomachDecreaceRate, config.PALWORLD_SERVER_SETTINGS.PlayerStomachDecreaceRate)

	fields = append(fields, &discordgo.MessageEmbedField{
		Inline: false,
		Value:  value,
	})

	value = fmt.Sprintf("- %s: %v\n", i18n.PlayerStaminaDecreaceRate, config.PALWORLD_SERVER_SETTINGS.PlayerStaminaDecreaceRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.PlayerAutoHPRegeneRate, config.PALWORLD_SERVER_SETTINGS.PlayerAutoHPRegeneRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.PlayerAutoHpRegeneRateInSleep, config.PALWORLD_SERVER_SETTINGS.PlayerAutoHpRegeneRateInSleep)
	value += fmt.Sprintf("- %s: %v\n", i18n.PalStomachDecreaceRate, config.PALWORLD_SERVER_SETTINGS.PalStomachDecreaceRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.PalStaminaDecreaceRate, config.PALWORLD_SERVER_SETTINGS.PalStaminaDecreaceRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.PalAutoHPRegeneRate, config.PALWORLD_SERVER_SETTINGS.PalAutoHPRegeneRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.PalAutoHpRegeneRateInSleep, config.PALWORLD_SERVER_SETTINGS.PalAutoHpRegeneRateInSleep)
	value += fmt.Sprintf("- %s: %v\n", i18n.BuildObjectDamageRate, config.PALWORLD_SERVER_SETTINGS.BuildObjectDamageRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.BuildObjectDeteriorationDamageRate, config.PALWORLD_SERVER_SETTINGS.BuildObjectDeteriorationDamageRate)
	value += fmt.Sprintf("- %s: %v", i18n.CollectionDropRate, config.PALWORLD_SERVER_SETTINGS.CollectionDropRate)

	fields = append(fields, &discordgo.MessageEmbedField{
		Inline: false,
		Value:  value,
	})

	value = fmt.Sprintf("- %s: %v\n", i18n.CollectionObjectHpRate, config.PALWORLD_SERVER_SETTINGS.CollectionObjectHpRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.CollectionObjectRespawnSpeedRate, config.PALWORLD_SERVER_SETTINGS.CollectionObjectRespawnSpeedRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.EnemyDropItemRate, config.PALWORLD_SERVER_SETTINGS.EnemyDropItemRate)
	value += fmt.Sprintf("- %s: %v\n", i18n.DeathPenalty, config.PALWORLD_SERVER_SETTINGS.DeathPenalty)
	value += fmt.Sprintf("- %s: %v\n", i18n.EnablePlayerToPlayerDamage, config.PALWORLD_SERVER_SETTINGS.EnablePlayerToPlayerDamage)
	value += fmt.Sprintf("- %s: %v\n", i18n.EnableFriendlyFire, config.PALWORLD_SERVER_SETTINGS.EnableFriendlyFire)
	value += fmt.Sprintf("- %s: %v\n", i18n.EnableInvaderEnemy, config.PALWORLD_SERVER_SETTINGS.EnableInvaderEnemy)
	value += fmt.Sprintf("- %s: %v\n", i18n.ActiveUNKO, config.PALWORLD_SERVER_SETTINGS.ActiveUNKO)
	value += fmt.Sprintf("- %s: %v\n", i18n.EnableAimAssistPad, config.PALWORLD_SERVER_SETTINGS.EnableAimAssistPad)
	value += fmt.Sprintf("- %s: %v", i18n.EnableAimAssistKeyboard, config.PALWORLD_SERVER_SETTINGS.EnableAimAssistKeyboard)

	fields = append(fields, &discordgo.MessageEmbedField{
		Inline: false,
		Value:  value,
	})

	value = fmt.Sprintf("- %s: %v\n", i18n.DropItemMaxNum, config.PALWORLD_SERVER_SETTINGS.DropItemMaxNum)
	value += fmt.Sprintf("- %s: %v\n", i18n.DropItemMaxNum_UNKO, config.PALWORLD_SERVER_SETTINGS.DropItemMaxNum_UNKO)
	value += fmt.Sprintf("- %s: %v\n", i18n.BaseCampMaxNum, config.PALWORLD_SERVER_SETTINGS.BaseCampMaxNum)
	value += fmt.Sprintf("- %s: %v\n", i18n.BaseCampWorkerMaxNum, config.PALWORLD_SERVER_SETTINGS.BaseCampWorkerMaxNum)
	value += fmt.Sprintf("- %s: %v\n", i18n.DropItemAliveMaxHours, config.PALWORLD_SERVER_SETTINGS.DropItemAliveMaxHours)
	value += fmt.Sprintf("- %s: %v\n", i18n.AutoResetGuildNoOnlinePlayers, config.PALWORLD_SERVER_SETTINGS.AutoResetGuildNoOnlinePlayers)
	value += fmt.Sprintf("- %s: %v\n", i18n.AutoResetGuildTimeNoOnlinePlayers, config.PALWORLD_SERVER_SETTINGS.AutoResetGuildTimeNoOnlinePlayers)
	value += fmt.Sprintf("- %s: %v\n", i18n.GuildPlayerMaxNum, config.PALWORLD_SERVER_SETTINGS.GuildPlayerMaxNum)
	value += fmt.Sprintf("- %s: %v\n", i18n.PalEggDefaultHatchingTime, config.PALWORLD_SERVER_SETTINGS.PalEggDefaultHatchingTime)
	value += fmt.Sprintf("- %s: %v", i18n.WorkSpeedRate, config.PALWORLD_SERVER_SETTINGS.WorkSpeedRate)

	fields = append(fields, &discordgo.MessageEmbedField{
		Inline: false,
		Value:  value,
	})

	value = fmt.Sprintf("- %s: %v\n", i18n.IsMultiplay, config.PALWORLD_SERVER_SETTINGS.IsMultiplay)
	value += fmt.Sprintf("- %s: %v\n", i18n.IsPvP, config.PALWORLD_SERVER_SETTINGS.IsPvP)
	value += fmt.Sprintf("- %s: %v\n", i18n.CanPickupOtherGuildDeathPenaltyDrop, config.PALWORLD_SERVER_SETTINGS.CanPickupOtherGuildDeathPenaltyDrop)
	value += fmt.Sprintf("- %s: %v\n", i18n.EnableNonLoginPenalty, config.PALWORLD_SERVER_SETTINGS.EnableNonLoginPenalty)
	value += fmt.Sprintf("- %s: %v\n", i18n.EnableFastTravel, config.PALWORLD_SERVER_SETTINGS.EnableFastTravel)
	value += fmt.Sprintf("- %s: %v\n", i18n.IsStartLocationSelectByMap, config.PALWORLD_SERVER_SETTINGS.IsStartLocationSelectByMap)
	value += fmt.Sprintf("- %s: %v\n", i18n.ExistPlayerAfterLogout, config.PALWORLD_SERVER_SETTINGS.ExistPlayerAfterLogout)
	value += fmt.Sprintf("- %s: %v\n", i18n.EnableDefenseOtherGuildPlayer, config.PALWORLD_SERVER_SETTINGS.EnableDefenseOtherGuildPlayer)
	value += fmt.Sprintf("- %s: %v\n", i18n.CoopPlayerMaxNum, config.PALWORLD_SERVER_SETTINGS.CoopPlayerMaxNum)
	value += fmt.Sprintf("- %s: %v", i18n.ServerPlayerMaxNum, config.PALWORLD_SERVER_SETTINGS.ServerPlayerMaxNum)

	fields = append(fields, &discordgo.MessageEmbedField{
		Inline: false,
		Value:  value,
	})

	value = fmt.Sprintf("- %s: %v\n", i18n.ServerName, config.PALWORLD_SERVER_SETTINGS.ServerName)
	value += fmt.Sprintf("- %s: %v\n", i18n.ServerDescription, config.PALWORLD_SERVER_SETTINGS.ServerDescription)
	value += fmt.Sprintf("- %s: %v\n", i18n.PublicPort, config.PALWORLD_SERVER_SETTINGS.PublicPort)
	value += fmt.Sprintf("- %s: %v\n", i18n.PublicIP, config.PALWORLD_SERVER_SETTINGS.PublicIP)
	value += fmt.Sprintf("- %s: %v\n", i18n.RCONEnabled, config.PALWORLD_SERVER_SETTINGS.RCONEnabled)
	value += fmt.Sprintf("- %s: %v\n", i18n.Region, config.PALWORLD_SERVER_SETTINGS.Region)
	value += fmt.Sprintf("- %s: %v\n", i18n.UseAuth, config.PALWORLD_SERVER_SETTINGS.UseAuth)
	value += fmt.Sprintf("- %s: %s", i18n.BanListURL, config.PALWORLD_SERVER_SETTINGS.BanListURL)
	// value += fmt.Sprintf("- %s: %v\n", i18n.AdminPassword, config.PALWORLD_SERVER_SETTINGS.AdminPassword) // Do not show AdminPassword
	// value += fmt.Sprintf("- %s: %v\n", i18n.ServerPassword, config.PALWORLD_SERVER_SETTINGS.ServerPassword) // Do not show ServerPassword
	// value += fmt.Sprintf("- %s: %v\n", i18n.RCONPort, config.PALWORLD_SERVER_SETTINGS.RCONPort) // Do not show RCONPort

	fields = append(fields, &discordgo.MessageEmbedField{
		Inline: false,
		Value:  value,
	})

	return &discordgo.MessageEmbed{
		Title:  "PalWorldSettings.ini",
		Fields: fields,
	}
}
