package config

import palworld_settings "github.com/juunini/palworld-settings"

const (
	DEFAULT_PALWORLD_WINDOWS_SERVER_PATH = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\PalServer\\PalServer.exe"
	DEFAULT_PALWORLD_LINUX_SERVER_PATH   = "~/Steam/steamapps/common/PalServer/PalServer.sh"

	DEFAULT_PALWORLD_WINDOWS_SETTINGS_PATH = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\PalServer\\Pal\\Saved\\Config\\WindowsServer\\PalWorldSettings.ini"
	DEFAULT_PALWORLD_LINUX_SETTINGS_PATH   = "~/steamapps/common/PalServer/Pal/Saved/Config/LinuxServer/PalWorldSettings.ini"
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
	PALWORLD_RCON_PORT                 = 25575
	PALWORLD_ADMIN_PASSWORD            = ""
	PALWORLD_SERVER_FILE_PATH          = palworldServerPathFromEnv()
	PALWORLD_SERVER_EXECUTE_FLAGS      = []string{"-useperfthreads", "-NoAsyncLoadingThread", "-UseMultithreadForDS"}
	PALWORLD_SERVER_SETTINGS_FILE_PATH = palworldServerSettingsFromEnv()

	DISCORD_DASHBOARD_CHANNEL_ID                   = ""
	DISCORD_LOG_CHANNEL_ID                         = ""
	DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID    = ""
	DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID   = ""
	DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID        = ""
	DISCORD_DASHBOARD_PALWORLD_SETTINGS_MESSAGE_ID = ""

	DISCORD_COMMAND_ALIAS_HELP            = "help"
	DISCORD_COMMAND_ALIAS_KICK            = "kick"
	DISCORD_COMMAND_ALIAS_BAN             = "ban"
	DISCORD_COMMAND_ALIAS_BROADCAST       = "broadcast"
	DISCORD_COMMAND_ALIAS_SHUTDOWN        = "shutdown"
	DISCORD_COMMAND_ALIAS_DO_EXIT         = "doExit"
	DISCORD_COMMAND_ALIAS_SAVE            = "save"
	DISCORD_COMMAND_ALIAS_START_SERVER    = "startServer"
	DISCORD_COMMAND_ALIAS_SERVER_SETTINGS = "serverSettings"

	LANGUAGE = ""

	PALWORLD_SERVER_SETTINGS = palworld_settings.DEFAULT
)
