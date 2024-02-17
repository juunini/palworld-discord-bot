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
	FailedToBroadcastCommand = "Broadcast 커맨드를 실패했습니다."
	FailedToKickCommand = "Kick 커맨드를 실패했습니다."
	FailedToBanCommand = "Ban 커맨드를 실패했습니다."
	WrongParameters = "잘못된 파라미터를 입력하셨습니다. 다시 확인해주세요."
	Help = func(commandPrefix string, isAdmin bool) string {
		message := fmt.Sprintf("`%s help` - 도움말을 표시합니다.\n", commandPrefix)

		if !isAdmin {
			return message
		}

		message += fmt.Sprintf("`%s kick <SteamID>` - <SteamID> 를 1회 추방합니다. 다시 접속이 가능합니다.\n", commandPrefix)
		message += fmt.Sprintf("`%s ban <SteamID>` - <SteamID> 를 다시 접속할 수 없게 블락시킵니다.\n", commandPrefix)
		message += fmt.Sprintf("`%s broadcast <메시지>` - SYSTEM 메시지로 모든 유저에게 <메시지>를 전송합니다.\n", commandPrefix)
		message += fmt.Sprintf("`%s shutdown <초> <메시지>` - <메시지>를 전송한 후 <초> 후에 서버를 종료합니다.\n", commandPrefix)
		message += fmt.Sprintf("`%s doExit` - 서버를 강제 종료합니다.\n", commandPrefix)
		message += fmt.Sprintf("`%s save` - 서버를 저장합니다.", commandPrefix)
		return message
	}
	UnknownCommand = "알 수 없는 명령어입니다."
}
