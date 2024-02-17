package i18n

func SetLanguage(language string) {
	switch language {
	case "ko":
		ko()
		return

	default:
		en()
	}
}

var BotRunningStart string
var FailedToConnectRconServer string
var FailedToSaveCommand string
var FailedToDoExitCommand string
var Help string
var UnknownCommand string
