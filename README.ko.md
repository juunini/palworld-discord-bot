# <div align="center">Palworld 디스코드 봇</div>

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

## 설치 방법

1. [Release](https://github.com/juunini/palworld-discord-bot/releases) 페이지에 가서 운영체제에 맞는 파일을 다운로드 받으세요.
2. 봇의 기능을 온전하게 쓰려면 팰월드 서버를 실행한 곳에서 같이 실행해야 합니다.
3. 다운로드 받은 파일을 실행하세요.
    - 실행하면 아래와 같은 화면이 나옵니다. 언어를 선택하세요.  
        ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/5950ee92-6e7b-4491-86fe-490ea4039cb0)  
    - 언어를 선택하면 아래와 같은 화면이 나옵니다. 녹색으로 표시된 주소를 브라우저에 입력하세요.  
        ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/33a776f2-b95a-4d3d-9814-2023705796c5)  
4. 브라우저에 주소를 입력하면 아래와 같은 화면이 나옵니다. 설정하신 후 재실행 하세요.  
    ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/d17c1e13-01d4-4d76-babc-1fee74f2f5be)  

## 세팅 방법

1. `PalworldSettings.ini` 파일의 `RCONEnabled` 값을 `True`로 변경하세요.  
    - 봇을 실행 후 디스코드에서 `!palbot serverSettings RCONEnabled true` 를 입력하셔도 됩니다.
2. `PalworldSettings.ini` 파일의 `AdminPassword` 값을 변경하세요.  
    - 봇을 실행 후 디스코드에서 `!palbot serverSettings AdminPassword <변경할 비밀번호>` 를 입력하셔도 됩니다.
3. 팰월드 서버를 재시작 하세요.
4. `http://localhost:60000` 에서 `팰월드 RCON 호출 기능을 사용합니다.` 부분을 체크하거나, `.env` 파일의 `PALWORLD_RCON_ENABLED` 값을 `true`로 변경하세요.
5. `http://localhost:60000` 에서 `팰월드 관리자 비밀번호` 부분을 변경하거나, `.env` 파일의 `PALWORLD_ADMIN_PASSWORD` 값을 변경하세요.


## 기능

| 기능 | 설명 | 이미지 |
|------|-------------|---------|
| 국제화 | 한국어(ko), 영어(en) 를 지원합니다. | |
| 접속중인 유저 대시보드 | 서버에 접속한 유저의 리스트를 보여줍니다. 유저의 변동이 생기면 자동으로 업데이트 됩니다. | ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/8a003361-28bc-41fd-9332-6939eea44053) |
| 유저 접속 알림 | 서버에 유저가 접속하거나, 떠나면 로그 채널에 메시지를 보냅니다. | ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/80b31996-0770-4276-8d5c-259feb43c70f) |
| 커맨드 | 관리자는 커맨드를 이용하여 서버를 관리할 수 있습니다. `!palbot` 또는 `!palbot help` 라고 쓰면 사용법을 출력합니다. | ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/5b5887e4-20f6-4d61-b623-441368305b34) |
| kick(관리자 전용) | 유저를 추방합니다. 다시 접속이 가능합니다. | ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/ab25b40c-3bd9-4a98-973b-5f0b0553786f) |
| ban(관리자 전용) | 유저를 차단합니다. 다시 접속이 불가능합니다. | |
| broadcast(관리자 전용) | 서버에 메시지를 전송합니다. 현재 영어만 사용이 가능합니다. [참고](https://github.com/juunini/palworld-rcon/issues/1) | ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/6912d32b-0458-4336-8afc-6b1bb85f417c) |
| shutdown(관리자 전용) | 서버를 종료합니다. | ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/29a0caca-803e-4ba0-ad6c-e2b296736b7e) |
| doExit(관리자 전용) | 서버를 강제 종료합니다. | ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/299174d4-a749-4f14-a7c5-ceb5bc6cdccb) |
| save(관리자 전용) | 서버를 저장합니다. | |
| startServer(관리자 전용) | 서버를 실행합니다. (윈도우: `C:\Program Files (x86)\Steam\steamapps\common\PalServer\PalServer.exe`, 리눅스: `~/Steam/steamapps/common/PalServer/PalServer.sh`) | ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/1a7ebb68-eb78-41ac-86a0-1cc89c404ce1) |
| serverSettings(관리자 전용) | 서버 설정을 변경합니다. | ![image](https://github.com/juunini/palworld-discord-bot/assets/41536271/a383b0cb-98d1-4945-b4cb-7a559601e073) |

## 주의사항

- 현재 Palworld는 RCON으로 메시지를 전송 시 Non-ASCII 문자를 지원하지 않습니다. 영어 말고 다른 언어를 사용할 수 없다고 보시면 됩니다. [참고](https://github.com/juunini/palworld-rcon/issues/1)
- Palworld의 RCON은 메시지를 보낼 때 띄어쓰기가 존재하면 메시지를 잘라버리는 현상이 발생합니다. 따라서, 띄어쓰기 부분은 `_`로 대체되어 전송됩니다. (broadcast, shutdown 커맨드)
