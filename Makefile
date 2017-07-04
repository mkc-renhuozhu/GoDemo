BASE_DIR        ?= `pwd`
BIN_PATH        ?= "$(BASE_DIR)/bin"
LINUX_BIN_PATH  ?= "$(BASE_DIR)/bin/linux_amd64"
WINDOWS_BIN_PATH ?="$(BASE_DIR)/bin/windows_amd64"
CMD_PKG         ?= helloGo/cmd

all:build

help:
	@echo "available targets:"
	@echo "  * linux      - build linux binary"
	@echo "  * build      - build"
	@echo "  * windows    - build windows binary"

linux:
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_BIN_PATH)/helloGo $(CMD_PKG)

build:
	go build -o $(BIN_PATH)/helloGo $(CMD_PKG)

windows:
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_BIN_PATH)/helloGo.exe $(CMD_PKG)