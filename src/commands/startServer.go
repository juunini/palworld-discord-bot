package commands

import (
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/juunini/palworld-discord-bot/src/console_decoration"
	"github.com/juunini/palworld-discord-bot/src/i18n"
)

func startServer() string {
	serverPath := "~/Steam/steamapps/common/PalServer/PalServer.sh"

	if runtime.GOOS == "windows" {
		serverPath = filepath.Join("C:", "/", "Program Files (x86)", "Steam", "steamapps", "common", "PalServer", "PalServer.exe")
	}

	if err := exec.Command(serverPath, "-useperfthreads", "-NoAsyncLoadingThread", "-UseMultithreadForDS").Start(); err != nil {
		console_decoration.PrintError(i18n.FailedToStartServerCommand + ": " + err.Error())
		return i18n.FailedToStartServerCommand
	}

	return i18n.SuccessToStartServerCommand
}
