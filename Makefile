.PHONY: all build-win build-linux run gotool clean help

BINARY="bluebell"

all: gotool build

build-win:
	go build -o ${BINARY}

build-linux:
	SET CGO_ENABLED=0
	SET GOOS=linux
	SET GOARCH=amd64
	go build -o ${BINARY}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - Formatting Go code, and compiling it into binaries"
	@echo "make build - Compile Go code, generate binary files"
	@echo "make run - Run Go code"
	@echo "make clean - Remove binaries and vim swap files"
	@echo "make gotool - Run the Go tools 'fmt' and 'vet'"