package commands

import (
	"os/exec"

	"github.com/juunini/palworld-discord-bot/src/config"
	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
)

func startServer() string {
	if err := exec.Command(
		config.PALWORLD_SERVER_FILE_PATH,
		config.PALWORLD_SERVER_EXECUTE_FLAGS...,
	).Start(); err != nil {
		console_decoration.PrintError(i18n.FailedToStartServerCommand + ": " + err.Error())
		return i18n.FailedToStartServerCommand
	}

	return i18n.SuccessToStartServerCommand
}
