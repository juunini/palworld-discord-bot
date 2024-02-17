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
var FailedToBroadcastCommand string
var FailedToKickCommand string
var FailedToBanCommand string
var WrongParameters string
var Help func(string) string
var UnknownCommand string
