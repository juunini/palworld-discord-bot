package commands

import (
	"strings"

	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
	"github.com/juunini/palworld-discord-bot/src/utils"
)

func shutdown(command string) string {
	client, err := utils.RconClient()
	if err != nil {
		return i18n.FailedToConnectRconServer
	}

	defer client.Disconnect()

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
	paramString, _ := strings.CutPrefix(command, "shutdown ")
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
