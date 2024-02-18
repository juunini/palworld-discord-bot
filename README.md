# <div align="center">Palworld Discord BOT</div>

<div align="center">
    <img src="https://github.com/juunini/palworld-rcon/assets/41536271/8414cd69-68f4-45bc-a052-9c4afa652582" alt="Palworld Icon" />
</div>

<div align="center">
    <img src="https://img.shields.io/badge/Palworld-f5d601?style=for-the-badge&logoColor=white&logo=data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAAAAAAD/2wBDAAMCAgICAgMCAgIDAwMDBAYEBAQEBAgGBgUGCQgKCgkICQkKDA8MCgsOCwkJDRENDg8QEBEQCgwSExIQEw8QEBD/2wBDAQMDAwQDBAgEBAgQCwkLEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBD/wAARCAAgACADAREAAhEBAxEB/8QAGAAAAwEBAAAAAAAAAAAAAAAABQYIBwT/xAAuEAABAwIDBQgCAwAAAAAAAAACAwQFBhIAEyIHFCMyUggVJDNCYnKCNJJjovL/xAAZAQEBAQADAAAAAAAAAAAAAAAHCAUEBgn/xAArEQAABQMDAwIHAQAAAAAAAAAAAQIDBAUGEQchURITQTEyCBUiM2FxkbH/2gAMAwEAAhEDEQA/AKtqKoYqmIV1PzrnJZskjVVK24/iI+ssee1NprlZkogQC9TCc7IJho1ubBIiXVeVqn3w9mF6Xj1fxY9kk3VdWfzqqiQXe1IBt6y58V9bGhlKYYblVTdRkBq4tQFQ3TbYMOyezmVSp/vuO2qTBOgSzct2g1XQI+khyhP9THGvUNHLdmsKjMMdJ+DyY4cO+104u/Ic6j4wB9N1G8eOnUDPN0m0s1SBXgFcg5SLTmpX+7nH0n1CQGczah6eO2ZMQl77Z5wFS27ki3GgnUe4Z/2jpdGLiYU5JbKjQlGhui9AcUbCL72409Do7D1xttyC5/wcm5UqehKU2HCPScykPBrQkqKCIZSquUkCu9JZXlezWQlcPRi9kHG7TZHuRZErSEZmqS4WTHZIS7KLcvGLwFxdMvPbCgRKj9cZr9Xhpynuen4GxEtGtSy76GPp/YA0nNs63q5rPU8i8KPimT1q4cuWqrbxCqqXAEVREzIMgr+m0MTR8QNx0ie21FhK6nCzkKthUebAe65KekgF7UUc2e7I5xZ1b4dqqqHzDih/dIcDGlsv5fcjDSN8nuEqctR091KvJDs7JUS5pSjqdZzDzei3zPtvuFASMSAB+N13749A6y2pmP0xfJb/AME1dKY1SI1cjc9upOomFZz7Gn4tyo4et45w/epEW4gqZCCpCJCao5pJDbePm3dWAeuVWVToSpLZZXvkNFKmHHw2R+oSqYp9tT0b3e3WVXVNU13C6vOuqZ3GRWaPqOgeXEj3BUJNRmHOlcjuaEm2nuEJd2pbRqk7R7MqS2UQLkYVoqBPXLshSN0fMACN2gdN+v8A3WOmWjp22+iq1LfgH9fu5DbSW0+RtWxulZWhdnsPTc48Fd80AyNVI7g1mRAIl7AtH6YpGUpKo6+3wA2bKOTJU4nwZCqZmEY1pSLyBmAPdptgbVe3SQiqFpW+7VgHq8c3VKaMvfsQaKe+TiUucEQnig6rczYOoKYbKtaggRSSlEiELCPUOakQEXCI0lbPXp1DiW79sibaj5nM9izyQ79AqKZae3wP/9k=" alt="Palworld" />
    <img src="https://img.shields.io/badge/Discord-5865F2?style=for-the-badge&logo=discord&logoColor=white" alt="Discord" />
</div>

---

[한국어](./README.ko.md) | [English](./README.md)

---

## Install

1. Go to the [Release](https://github.com/juunini/palworld-discord-bot/releases) page and download the file that matches your operating system.
2. Run the downloaded file.
    - When you run it, you will be prompted to enter the required values in the terminal.
    - If you prefer to set the values directly instead of entering them in the terminal, copy the `.env.sample` file and create a `.env` file, then fill in the necessary values.
    - Refer to the [Environment Variables](#environment-variables) section below for the required values.

## Features

- Displays a list of users connected to the server. It automatically updates when there are changes in the user list.
    - ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/8a003361-28bc-41fd-9332-6939eea44053)
- Sends a message to the log channel when a user connects or disconnects from the server.
    - ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/80b31996-0770-4276-8d5c-259feb43c70f)
- Administrators can manage the server using commands. Typing `!palbot` or `!palbot help` will display the usage. [Reference](https://tech.palworldgame.com/settings-and-operation/commands)
    - ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/93c77216-cd17-43f9-a44a-3b4b4d553914)
    - ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/ab25b40c-3bd9-4a98-973b-5f0b0553786f)

## Environment Variables

| Name | Description | Default |
|------|-------------|---------|
| DISCORD_BOT_TOKEN | | |
| DISCORD_ADMIN_USERNAMES | seperated by comma, example: "admin1, admin2, admin3" | |
| DISCORD_COMMAND_PREFIX | | !palbot |
| PALWORLD_RCON_HOST | | 127.0.0.1 |
| PALWORLD_RCON_PORT | | 25575 |
| PALWORLD_RCON_PASSWORD | | |
| DISCORD_DASHBOARD_CHANNEL_ID | Bot displays the status of the server, the connection user, and so on in this channel. | |
| DISCORD_LOG_CHANNEL_ID | Bot sends a log to this channel, e.g. user connection, user connection termination, etc... | |
| LANGUAGE | Currently supports Korean(ko), English(en) | en |
