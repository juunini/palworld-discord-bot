package i18n

func SetLanguage(language string) {
	en()

	switch language {
	case "ko":
		ko()
		return
	}
}

var BotRunningStart string
var FailedToConnectRconServer string
var FailedToSaveCommand string
var FailedToDoExitCommand string
var FailedToShutdownCommand string
var WrongParameters string
var Help string
var UnknownCommand string
