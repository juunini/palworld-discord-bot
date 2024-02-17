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

var Help string
var UnknownCommand string
