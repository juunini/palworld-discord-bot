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

	paramString, _ := strings.CutPrefix(command, "shutdown ")
	params := strings.Split(paramString, " ")

	if len(params) < 2 {
		console_decoration.PrintError(i18n.WrongParameters)
		return i18n.WrongParameters
	}

	seconds, err := utils.ToUint(params[0])
	if err != nil {
		console_decoration.PrintError(i18n.WrongParameters + ": " + err.Error())
		return i18n.WrongParameters
	}

	message := params[1]
	response, err := client.Shutdown(seconds, message)
	if err != nil {
		console_decoration.PrintError(i18n.FailedToShutdownCommand + ": " + err.Error())
		return i18n.FailedToShutdownCommand
	}

	return response
}
