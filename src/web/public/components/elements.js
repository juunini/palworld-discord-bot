export const elements = {
  webServerEnabled: document.getElementById("webServerEnabled"),
  webServerPort: document.getElementById("webServerPort"),

  language: document.getElementById("language"),

  discordBotEnabled: document.getElementById("discordBotEnabled"),
  discordBotToken: document.getElementById("discordBotToken"),
  discordAdminUsernames: document.getElementById("discordAdminUsernames"),
  discordCommandCaseSensitive: document.getElementById("discordCommandCaseSensitive"),
  discordCommandPrefix: document.getElementById("discordCommandPrefix"),

  palworldRconEnabled: document.getElementById("palworldRconEnabled"),
  palworldRconHost: document.getElementById("palworldRconHost"),
  palworldRconPort: document.getElementById("palworldRconPort"),
  palworldAdminPassword: document.getElementById("palworldAdminPassword"),
  palworldServerFilePath: document.getElementById("palworldServerFilePath"),
  palworldServerExecuteFlags: document.getElementById("palworldServerExecuteFlags"),
  palworldServerSettingsFilePath: document.getElementById("palworldServerSettingsFilePath"),

  discordDashboardChannelId: document.getElementById("discordDashboardChannelId"),
  discordLogChannelId: document.getElementById("discordLogChannelId"),
  discordDashboardOnlinePlayersMessageId: document.getElementById("discordDashboardOnlinePlayersMessageId"),
  discordDashboardPalworldConfigMessageId: document.getElementById("discordDashboardPalworldConfigMessageId"),
  discordDashboardBotConfigMessageId: document.getElementById("discordDashboardBotConfigMessageId"),

  discordCommandAliasHelp: document.getElementById("discordCommandAliasHelp"),
  discordCommandAliasKick: document.getElementById("discordCommandAliasKick"),
  discordCommandAliasBan: document.getElementById("discordCommandAliasBan"),
  discordCommandAliasBroadcast: document.getElementById("discordCommandAliasBroadcast"),
  discordCommandAliasShutdown: document.getElementById("discordCommandAliasShutdown"),
  discordCommandAliasDoExit: document.getElementById("discordCommandAliasDoExit"),
  discordCommandAliasSave: document.getElementById("discordCommandAliasSave"),
  discordCommandAliasStartServer: document.getElementById("discordCommandAliasStartServer"),
  discordCommandAliasServerSettings: document.getElementById("discordCommandAliasServerSettings"),
};

export function setValueInElements(config) {
  elements.webServerEnabled.setAttribute("checked", config.WEB_SERVER_ENABLED);
  elements.webServerPort.setAttribute("value", config.WEB_SERVER_PORT);

  elements.language.value = config.LANGUAGE;

  elements.discordBotEnabled.setAttribute("checked", config.DISCORD_BOT_ENABLED);
  elements.discordBotToken.setAttribute("value", config.DISCORD_BOT_TOKEN);
  elements.discordAdminUsernames.setAttribute("value", config.DISCORD_ADMIN_USERNAMES);
  elements.discordCommandCaseSensitive.setAttribute("checked", config.DISCORD_COMMAND_CASE_SENSITIVE);
  elements.discordCommandPrefix.setAttribute("value", config.DISCORD_COMMAND_PREFIX);

  elements.palworldRconEnabled.setAttribute("checked", config.PALWORLD_RCON_ENABLED);
  elements.palworldRconHost.setAttribute("value", config.PALWORLD_RCON_HOST);
  elements.palworldRconPort.setAttribute("value", config.PALWORLD_RCON_PORT);
  elements.palworldAdminPassword.setAttribute("value", config.PALWORLD_ADMIN_PASSWORD);
  elements.palworldServerFilePath.setAttribute("value", config.PALWORLD_SERVER_FILE_PATH);
  elements.palworldServerExecuteFlags.setAttribute("value", config.PALWORLD_SERVER_EXECUTE_FLAGS);
  elements.palworldServerSettingsFilePath.setAttribute("value", config.PALWORLD_SERVER_SETTINGS_FILE_PATH);

  elements.discordDashboardChannelId.setAttribute("value", config.DISCORD_DASHBOARD_CHANNEL_ID);
  elements.discordLogChannelId.setAttribute("value", config.DISCORD_LOG_CHANNEL_ID);
  elements.discordDashboardOnlinePlayersMessageId.setAttribute("value", config.DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID);
  elements.discordDashboardPalworldConfigMessageId.setAttribute("value", config.DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID);
  elements.discordDashboardBotConfigMessageId.setAttribute("value", config.DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID);

  elements.discordCommandAliasHelp.setAttribute("value", config.DISCORD_COMMAND_ALIAS_HELP);
  elements.discordCommandAliasKick.setAttribute("value", config.DISCORD_COMMAND_ALIAS_KICK);
  elements.discordCommandAliasBan.setAttribute("value", config.DISCORD_COMMAND_ALIAS_BAN);
  elements.discordCommandAliasBroadcast.setAttribute("value", config.DISCORD_COMMAND_ALIAS_BROADCAST);
  elements.discordCommandAliasShutdown.setAttribute("value", config.DISCORD_COMMAND_ALIAS_SHUTDOWN);
  elements.discordCommandAliasDoExit.setAttribute("value", config.DISCORD_COMMAND_ALIAS_DO_EXIT);
  elements.discordCommandAliasSave.setAttribute("value", config.DISCORD_COMMAND_ALIAS_SAVE);
  elements.discordCommandAliasStartServer.setAttribute("value", config.DISCORD_COMMAND_ALIAS_START_SERVER);
  elements.discordCommandAliasServerSettings.setAttribute("value", config.DISCORD_COMMAND_ALIAS_SERVER_SETTINGS);
}

export function configFromElements() {
  return {
    WEB_SERVER_ENABLED: elements.webServerEnabled.checked,
    WEB_SERVER_PORT: Number(elements.webServerPort.value),

    LANGUAGE: elements.language.value,

    DISCORD_BOT_ENABLED: elements.discordBotEnabled.checked,
    DISCORD_BOT_TOKEN: elements.discordBotToken.value,
    DISCORD_ADMIN_USERNAMES: elements.discordAdminUsernames.value,
    DISCORD_COMMAND_CASE_SENSITIVE: elements.discordCommandCaseSensitive.checked,
    DISCORD_COMMAND_PREFIX: elements.discordCommandPrefix.value,

    PALWORLD_RCON_ENABLED: elements.palworldRconEnabled.checked,
    PALWORLD_RCON_HOST: elements.palworldRconHost.value,
    PALWORLD_RCON_PORT: Number(elements.palworldRconPort.value),
    PALWORLD_ADMIN_PASSWORD: elements.palworldAdminPassword.value,
    PALWORLD_SERVER_FILE_PATH: elements.palworldServerFilePath.value,
    PALWORLD_SERVER_EXECUTE_FLAGS: elements.palworldServerExecuteFlags.value,
    PALWORLD_SERVER_SETTINGS_FILE_PATH: elements.palworldServerSettingsFilePath.value,

    DISCORD_DASHBOARD_CHANNEL_ID: elements.discordDashboardChannelId.value,
    DISCORD_LOG_CHANNEL_ID: elements.discordLogChannelId.value,
    DISCORD_DASHBOARD_ONLINE_PLAYERS_MESSAGE_ID: elements.discordDashboardOnlinePlayersMessageId.value,
    DISCORD_DASHBOARD_PALWORLD_CONFIG_MESSAGE_ID: elements.discordDashboardPalworldConfigMessageId.value,
    DISCORD_DASHBOARD_BOT_CONFIG_MESSAGE_ID: elements.discordDashboardBotConfigMessageId.value,

    DISCORD_COMMAND_ALIAS_HELP: elements.discordCommandAliasHelp.value,
    DISCORD_COMMAND_ALIAS_KICK: elements.discordCommandAliasKick.value,
    DISCORD_COMMAND_ALIAS_BAN: elements.discordCommandAliasBan.value,
    DISCORD_COMMAND_ALIAS_BROADCAST: elements.discordCommandAliasBroadcast.value,
    DISCORD_COMMAND_ALIAS_SHUTDOWN: elements.discordCommandAliasShutdown.value,
    DISCORD_COMMAND_ALIAS_DO_EXIT: elements.discordCommandAliasDoExit.value,
    DISCORD_COMMAND_ALIAS_SAVE: elements.discordCommandAliasSave.value,
    DISCORD_COMMAND_ALIAS_START_SERVER: elements.discordCommandAliasStartServer.value,
    DISCORD_COMMAND_ALIAS_SERVER_SETTINGS: elements.discordCommandAliasServerSettings.value,
  }
}
