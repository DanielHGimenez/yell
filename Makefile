
build-dev:
	go build -o ./build/debug/yell ./src/main.go

build-release:
	# Windows
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./build/release/bin/amd64-win/yell.exe ./src/main.go # 64-bit
	GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o ./build/release/bin/386-win/yell.exe ./src/main.go # 32-bit

	# MacOS
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./build/release/bin/amd64-darwin/yell ./src/main.go # 64-bit
	#GOOS=darwin GOARCH=386 go build -ldflags "-s -w" -o ./build/release/bin/386-darwin/yell ./src/main.go # 32-bit
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ./build/release/bin/arm64-darwin/yell ./src/main.go # Apple Silicon

	# Linux
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./build/release/bin/amd64-linux/yell ./src/main.go # 64-bit
	#GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o ./build/release/bin/386-linux/yell ./src/main.go # 32-bit

assembly-release-packages:
	mkdir -p ./build/release/package
	zip -j ./build/release/package/yell-win-x64 ./build/release/bin/amd64-win/yell.exe notification.mp3 example-config.yaml
	zip -j ./build/release/package/yell-win-x86 ./build/release/bin/386-win/yell.exe notification.mp3 example-config.yaml
	zip -j ./build/release/package/yell-macos-intel-x64 ./build/release/bin/amd64-darwin/yell notification.mp3 example-config.yaml
	zip -j ./build/release/package/yell-macos-silicon-arm64 ./build/release/bin/arm64-darwin/yell notification.mp3 example-config.yaml
	zip -j ./build/release/package/yell-linux-x64 ./build/release/bin/amd64-linux/yell notification.mp3 example-config.yaml
