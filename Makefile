.PHONY: all build run gotool clean help

BINARY="main"

all: gotool build run

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./${BINARY}

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ]; then rm ${BINARY}; fi

help:
	@echo "make - 格式化go代码，并编译成二进制文件，运行二进制文件"
	@echo "make build - 编译go代码成二进制文件"
	@echo "make run - 直接运行go代码"