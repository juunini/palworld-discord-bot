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

var (
	BotRunningStart             string
	FailedToConnectRconServer   string
	FailedToSaveCommand         string
	FailedToDoExitCommand       string
	FailedToShutdownCommand     string
	FailedToBroadcastCommand    string
	FailedToKickCommand         string
	FailedToBanCommand          string
	FailedToShowPlayerCommand   string
	FailedToStartServerCommand  string
	SuccessToStartServerCommand string
	WrongParameters             string
	UnknownCommand              string
	Help                        func(commandPrefix string, isAdmin bool) string
)

var (
	WebConfig                    string
	EnableWebServer              string
	EnableWebServerDisablePrompt string
	WebServerPort                string
	WebServerPortTooltip         string
)
