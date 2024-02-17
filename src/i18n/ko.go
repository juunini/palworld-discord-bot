package i18n

import (
	"fmt"
)

func ko() {
	BotRunningStart = "봇이 실행되었습니다. 종료하려면 CTRL-C를 누르세요."
	FailedToConnectRconServer = "RCON 서버에 연결하는데 실패했습니다."
	FailedToSaveCommand = "Save 커맨드를 실패했습니다."
	FailedToDoExitCommand = "DoExit 커맨드를 실패했습니다."
	FailedToShutdownCommand = "Shutdown 커맨드를 실패했습니다."
	WrongParameters = "잘못된 파라미터를 입력하셨습니다. 다시 확인해주세요."
	Help = func(commandPrefix string) string {
		return fmt.Sprintf(`%s help - 도움말을 표시합니다.`, commandPrefix)
	}
	UnknownCommand = "알 수 없는 명령어입니다."
}
