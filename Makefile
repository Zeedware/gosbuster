PROJECT_NAME = gosbuster
BUILD_DIR = build/

install:
	go mod tidy

build:
	go build -o $(BUILD_DIR)$(PROJECT_NAME)

test:
	go vet
	golint
	go test

getpath:
	@echo $(BUILD_DIR)$(PROJECT_NAME)

build_windows:
	env GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)${PROJECT_NAME}-windows-x64.exe

build_linux:
	env GOOS=linux GOARCH=amd64 go build -o build/v${PROJECT_NAME}-linux-x64

build_mac:
	env GOOS=darwin GOARCH=amd64 go build -o build/v${PROJECT_NAME}-mac-x64
