package utils

import "github.com/juunini/palworld-discord-bot/src/config"

func IsAdmin(username string) bool {
	for _, adminUsername := range config.DISCORD_ADMIN_USERNAMES {
		if username == adminUsername {
			return true
		}
	}

	return false
}
