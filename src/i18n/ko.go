package i18n

import (
	"fmt"

	"github.com/juunini/palworld-discord-bot/src/config"
)

func ko() {
	Help = fmt.Sprintf(`%s help - 도움말을 표시합니다.`, config.DISCORD_COMMAND_PREFIX)
	UnknownCommand = "알 수 없는 명령어입니다."
}
