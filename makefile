BINARY_NAME=tg_whi_exporter
EXPORTER_VERSION = ${EXPORTER_VERSION}

fmt:
		go fmt 
.PHONY:fmt

init:
	    go mod tidy
		go get golang.org/x/lint/golint
.PHONY:init

lint: fmt
		golint 
.PHONY:lint

vet: fmt
		go vet 
.PHONY:vet

build: vet
		GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-${EXPORTER_VERSION}-darwin-arm64 main.go
		GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-${EXPORTER_VERSION}-linux-amd64 main.go
.PHONY:build

clean:
		go clean
		rm -rf bin/
.PHONY:clean