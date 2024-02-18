GOOS=windows GOARCH=amd64 go build -o out/palworld-discord-bot.exe
GOOS=linux GOARCH=amd64 go build -o out/palworld-discord-bot-linux
GOOS=linux GOARCH=386 go build -o out/palworld-discord-bot-linux-386
GOOS=linux GOARCH=arm64 go build -o out/palworld-discord-bot-linux-arm
GOOS=darwin GOARCH=amd64 go build -o out/palworld-discord-bot-macos
GOOS=darwin GOARCH=arm64 go build -o out/palworld-discord-bot-macos-apple-silicon
