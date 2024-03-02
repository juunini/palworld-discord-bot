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

	return []string{"Not implemented yet"}
}
