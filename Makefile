
build-dev:
	go build -o ./build/debug/yell ./src/main.go

build-release:
	# Windows
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./build/release/bin/yell-amd64.exe ./src/main.go # 64-bit
	GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o ./build/release/bin/yell-386.exe ./src/main.go # 32-bit

	# MacOS
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./build/release/bin/yell-amd64-darwin ./src/main.go # 64-bit
	#GOOS=darwin GOARCH=386 go build -ldflags "-s -w" -o ./build/release/bin/yell-386-darwin ./src/main.go # 32-bit
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ./build/release/bin/yell-arm64-darwin ./src/main.go # Apple Silicon

	# Linux
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./build/release/bin/yell-amd64-linux ./src/main.go # 64-bit
	#GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o ./build/release/bin/yell-386-linux ./src/main.go # 32-bit

assembly-release-packages:
	zip -j yell-win-x64 ./build/release/bin/yell-amd64.exe notification.mp3 example-config.yaml
	zip -j yell-win-x86 ./build/release/bin/yell-386.exe notification.mp3 example-config.yaml
	zip -j yell-darwin-x64 ./build/release/bin/yell-amd64-darwin notification.mp3 example-config.yaml
	zip -j yell-darwin-arm64 ./build/release/bin/yell-arm64-darwin notification.mp3 example-config.yaml
	zip -j yell-linux-x64 ./build/release/bin/yell-amd64-linux notification.mp3 example-config.yaml
