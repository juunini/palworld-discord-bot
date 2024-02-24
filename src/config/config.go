package config

const (
	DEFAULT_PALWORLD_WINDOWS_SERVER_PATH = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\PalServer\\PalServer.exe"
	DEFAULT_PALWORLD_LINUX_SERVER_PATH   = "~/Steam/steamapps/common/PalServer/PalServer.sh"
)

var (
	WEB_SERVER_ENABLED = true
	WEB_SERVER_PORT    = 60000

	DISCORD_BOT_ENABLED            = false
	DISCORD_BOT_TOKEN              = ""
	DISCORD_ADMIN_USERNAMES        = []string{}
	DISCORD_COMMAND_CASE_SENSITIVE = false
	DISCORD_COMMAND_PREFIX         = "!palbot"

	PALWORLD_RCON_ENABLED              = false
	PALWORLD_RCON_HOST                 = "127.0.0.1"
	PALWORLD_RCON_PORT            uint = 25575
	PALWORLD_ADMIN_PASSWORD            = ""
	PALWORLD_SERVER_FILE_PATH          = ""
	PALWORLD_SERVER_EXECUTE_FLAGS      = []string{"-useperfthreads", "-NoAsyncLoadingThread", "-UseMultithreadForDS"}

	DISCORD_DASHBOARD_CHANNEL_ID                 = ""
	DISCORD_LOG_CHANNEL_ID                       = ""
	DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID  = ""
	DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID = ""
	DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID      = ""

	LANGUAGE = ""
)
