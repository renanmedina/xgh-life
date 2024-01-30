all: build-linux build-windows build-macos
build-linux:
	env GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o ./builds/linux/xgh-bot-32
	env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./builds/linux/xgh-bot-64
build-windows:
	env GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o ./builds/windows/xhg-bot-win-32.exe
	env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./builds/windows/xhg-bot-win-64.exe
build-macos:
	env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./builds/macos/xhg-bot-macos-64