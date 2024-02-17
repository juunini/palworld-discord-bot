package i18n

import (
	"fmt"

	"github.com/juunini/palworld-discord-bot/src/config"
)

func ko() {
	BotRunningStart = "봇이 실행되었습니다. 종료하려면 CTRL-C를 누르세요."
	Help = fmt.Sprintf(`%s help - 도움말을 표시합니다.`, config.DISCORD_COMMAND_PREFIX)
	UnknownCommand = "알 수 없는 명령어입니다."
}