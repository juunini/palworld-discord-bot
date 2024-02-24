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
	FailedToShowPlayerCommand = "ShowPlayer 커맨드를 실패했습니다."
	FailedToStartServerCommand = "서버를 시작하는데 실패했습니다."
	SuccessToStartServerCommand = "서버를 실행했습니다."
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
		message += fmt.Sprintf("`%s startServer` - 서버를 실행합니다.", commandPrefix)
		return message
	}
	UnknownCommand = "알 수 없는 명령어입니다."

	WebServerOpeningMessage = "웹 서버에 접속하실 수 없는 환경이라면, .env 파일을 수정하신 후 다시 실행해주세요."

	WebConfig = "웹 설정"
	EnableWebServer = "웹 서버 사용"
	EnableWebServerDisablePrompt = "웹 서버를 사용하지 않으시겠습니까? 이후 다시 설정하시려면 .env 파일을 수정하셔야 합니다."
	WebServerPort = "웹 서버 포트"
	WebServerPortTooltip = "웹 서버의 포트를 설정합니다. 되도록 60000 이상의 숫자를 사용하시길 권장합니다."

	Language = "언어"
	SelectLanguage = "언어 선택"

	DiscordBotConfig = "디스코드 봇 설정"
	EnableDiscordBot = "디스코드 봇 사용"
	DiscordBotToken = "디스코드 봇 토큰"
	DiscordBotTokenTooltip = "디스코드 봇의 토큰을 입력하세요. 토큰을 얻는 방법은 현재 마우스를 올리고 있는 ? 버튼을 클릭하시면 확인하실 수 있습니다."
	DiscordBotTokenDescription = "디스코드 봇의 토큰을 입력하세요. 토큰을 얻는 방법은 https://github.com/juunini/palworld-discord-bot/wiki 를 참고하세요."
	DiscordAdminUsernames = "디스코드 관리자 닉네임"
	DiscordAdminUsernamesTooltip = "관리자의 '디스코드 닉네임'을 입력하세요. 여러명일 경우 쉼표로 구분합니다."
	DiscordCommandCaseSensitive = "디스코드 커맨드 대소문자 구분"
	DiscordCommandCaseSensitiveTooltip = "디스코드 커맨드의 대소문자를 구분할지 설정합니다."
	DiscordCommandPrefix = "디스코드 봇 호출 명령어"
	DiscordCommandPrefixTooltip = "디스코드 봇 호출 명령어를 설정합니다."

	PalworldConfig = "팰월드 설정"
	EnablePalworldRcon = "팰월드 RCON 호출 기능을 사용합니다."
	EnablePalworldRconTooltip = "PalWorldSettings.ini의 RCONEnable 값을 미리 true로 설정해주세요."
	PalworldRconHost = "팰월드 RCON 호스트"
	PalworldRconHostTooltip = "팰월드 서버가 실행된 서버의 주소를 입력하세요."
	PalworldRconPort = "팰월드 RCON 포트"
	PalworldRconPortTooltip = "PalWorldSettings.ini 파일에 설정된 RCONPort 값을 입력하세요. (포트포워딩이나 리버스 프록시를 이용해 변경하셨으면 해당 포트를 입력하세요.)"
	PalworldAdminPassword = "팰월드 관리자 비밀번호"
	PalworldAdminPasswordTooltip = "PalWorldSettings.ini 파일에 설정된 AdminPassword 값을 입력하세요."
	PalworldServerFilePath = "팰월드 서버 파일 경로"
	PalworldServerFilePathTooltip = "팰월드 서버가 실행된 파일의 경로를 입력하세요."
	PalworldServerExecuteFlags = "팰월드 서버 실행 플래그"
	PalworldServerExecuteFlagsTooltip = "팰월드 서버가 실행될 때 사용할 플래그를 입력하세요."

	DiscordChannelConfig = "디스코드 채널 설정"
	DiscordDashboardChannelID = "대시보드 채널 ID"
	DiscordDashboardChannelIDTooltip = "대시보드 기능을 이용하기 위한 채널 ID를 입력하세요. 입력하지 않으면 해당 기능을 사용하지 않습니다. 채널 ID를 확인하는 방법은 현재 마우스를 올리고 있는 ? 버튼을 클릭하시면 확인하실 수 있습니다."
	DiscordDashboardChannelIDDescription = "대시보드 기능을 이용하기 위한 채널 ID를 입력하세요. 입력하지 않으면 해당 기능을 사용하지 않습니다. 채널 ID를 확인하는 방법은 https://github.com/juunini/palworld-discord-bot/wiki 를 참고하세요."
	DiscordLogChannelID = "로그 채널 ID"
	DiscordLogChannelIDTooltip = "유저의 접속/종료 기록을 확인하기 위한 채널 ID를 입력하세요. 입력하지 않으면 해당 기능을 사용하지 않습니다. 채널 ID를 확인하는 방법은 현재 마우스를 올리고 있는 ? 버튼을 클릭하시면 확인하실 수 있습니다."
	DiscordLogChannelIDDescription = "유저의 접속/종료 기록을 확인하기 위한 채널 ID를 입력하세요. 입력하지 않으면 해당 기능을 사용하지 않습니다. 채널 ID를 확인하는 방법은 https://github.com/juunini/palworld-discord-bot/wiki 를 참고하세요."
	DiscordDashboardOnlinePlayersMessageID = "온라인 상태 대시보드 Message ID"
	DiscordDashboardOnlinePlayersMessageIDTooltip = "자동으로 설정되는 값입니다. 필요한 경우가 아니라면 수정하지 마세요."
	DiscordDashboardPalworldConfigMessageID = "팰월드 설정 대시보드 Message ID"
	DiscordDashboardPalworldConfigMessageIDTooltip = "자동으로 설정되는 값입니다. 필요한 경우가 아니라면 수정하지 마세요."
	DiscordDashboardBotConfigMessageID = "봇 설정 대시보드 Message ID"
	DiscordDashboardBotConfigMessageIDTooltip = "자동으로 설정되는 값입니다. 필요한 경우가 아니라면 수정하지 마세요."

	DiscordCommandAliases = "디스코드 커맨드 커스터마이징"
	DiscordCommandAliasKickTooltip = "유저를 1회 강퇴하는 커맨드를 커스텀 합니다."
	DiscordCommandAliasBanTooltip = "유저를 차단하는 커맨드를 커스텀 합니다."
	DiscordCommandAliasBroadcastTooltip = "모든 유저에게 메시지를 전송하는 커맨드를 커스텀 합니다. (영어만 가능)"
	DiscordCommandAliasShutdownTooltip = "서버를 일정 시간 후 종료하는 커맨드를 커스텀 합니다."
	DiscordCommandAliasDoExitTooltip = "서버를 강제 종료하는 커맨드를 커스텀 합니다."
	DiscordCommandAliasSaveTooltip = "서버를 저장하는 커맨드를 커스텀 합니다."
	DiscordCommandAliasStartServerTooltip = "서버를 실행하는 커맨드를 커스텀 합니다."

	Close = "닫기"
	Confirm = "확인"
	Save = "저장"
}
