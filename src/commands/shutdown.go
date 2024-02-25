package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
	palworldrcon "github.com/juunini/palworld-rcon"
)

func shutdown(client *palworldrcon.Client, command string) string {
	seconds, message, errorMessage := shutdownParameters(command)
	if errorMessage != "" {
		return errorMessage
	}

	response, err := client.Shutdown(seconds, message)
	if err != nil {
		console_decoration.PrintError(i18n.FailedToShutdownCommand + ": " + err.Error())
		return i18n.FailedToShutdownCommand
	}

	return response
}

func shutdownParameters(command string) (int, string, string) {
	paramString, _ := strings.CutPrefix(command, config.DISCORD_COMMAND_ALIAS_SHUTDOWN+" ")
	params := strings.Split(paramString, " ")

	if len(params) < 2 {
		console_decoration.PrintError(i18n.WrongParameters)
		return 0, "", i18n.WrongParameters
	}

	seconds, err := utils.ToInt(params[0])
	if err != nil {
		console_decoration.PrintError(i18n.WrongParameters + ": " + err.Error())
		return 0, "", i18n.WrongParameters
	}

	message := strings.Join(params[1:], "_")
	return seconds, message, ""
}
