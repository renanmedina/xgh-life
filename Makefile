build-linux:
	env GOOS=linux GOARCH=386 go build -o ./xgh-bot-32
	env GOOS=linux GOARCH=amd64 go build -o ./xgh-bot-64
build-windows:
	env GOOS=windows GOARCH=386 go build -o ./xhg-bot-win-32.exe
	env GOOS=windows GOARCH=amd64 go build -o ./xhg-bot-win-64.exe