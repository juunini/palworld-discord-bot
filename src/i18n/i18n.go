package i18n

func SetLanguage(language string) TranslateMap {
	switch language {
	case "ko":
		return ko

	default:
		return en
	}
}

type TranslateMap struct {
	Help           string
	UnknownCommand string
}
