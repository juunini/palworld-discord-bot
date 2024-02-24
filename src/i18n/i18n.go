package i18n

var Languages = []map[string]string{
	{"CODE": "ko", "NAME": "한국어"},
	{"CODE": "en", "NAME": "English"},
}

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
var FailedToShowPlayerCommand string
var FailedToStartServerCommand string
var SuccessToStartServerCommand string
var WrongParameters string
var Help func(commandPrefix string, isAdmin bool) string
var UnknownCommand string
