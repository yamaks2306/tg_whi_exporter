BINARY_NAME=tg_whi_exporter

fmt:
		go fmt 
.PHONY:fmt

lint: fmt
		golint 
.PHONY:lint

vet: fmt
		go vet 
.PHONY:vet

build: vet
		GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin-arm64 main.go
		GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-amd64 main.go
.PHONY:build

clean:
		go clean
		rm bin/${BINARY_NAME}-darwin-arm64
		rm bin/${BINARY_NAME}-linux-amd64
.PHONY:clean