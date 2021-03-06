BASEDIR = $(shell pwd)


# build with verison infos
versionDir = "github.com/puti-projects/puti/internal/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-extldflags -static -w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: build
build:
	@echo "Building binary file."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags ${ldflags} -o ./puti
clean:
	@echo "Cleaning."
	go clean
gotool:
	@echo "Running go tool."
	go vet .
test:
	@echo "Testing."
	go test -v ./...
ca:
	@echo "Generating ca files."
	openssl req -new -nodes -x509 -out configs/server.crt -keyout configs/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

help:
	@echo "make - run gotool and build"
	@echo "make build - compile the source code"
	@echo "make clean - run go clean"
	@echo "make gotool - run go tool 'vet'"
	@echo "make test - run go test"
	@echo "make ca - generate ca files"

.PHONY: build clean test gotool ca help
