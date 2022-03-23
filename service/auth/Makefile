# golang1.9 or latest
# 1. make help
# 2. make dep
# 3. make build
# ...

VERSION := $(shell echo $(shell cat version.go | grep "Version =" | cut -d '=' -f2))
APP_NAME := auth
BUILD_DIR := build
APP := ${BUILD_DIR}/${APP_NAME}
PKG_NAME := ${APP_NAME}_v${VERSION}
PKG := ${PKG_NAME}.tar.gz

LDFLAGS := -ldflags "-w -s -X github.com/txchat/dtalk/service/$(APP_NAME).GitCommit=`git rev-parse --short=8 HEAD`"
LDGRPC := -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"

.PHONY: clean build pkg

clean: ## Remove previous build
	@rm -rf ${BUILD_DIR}
	@go clean

build: #checkgofmt ## Build the binary file
	swag init -g server/http/http.go
	GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY=https://goproxy.cn,direct GOSUMDB="sum.golang.google.cn" go build -v $(LDGRPC) $(LDFLAGS) -o $(APP) cmd/main.go

pkg: build ## Package
	mkdir -p ${PKG_NAME}/bin
	mkdir -p ${PKG_NAME}/etc
	cp ${APP} ${PKG_NAME}/bin/
	cp config/$(APP_NAME).toml ${PKG_NAME}/etc/
	tar zvcf ${PKG} ${PKG_NAME}
	rm -rf ${PKG_NAME}

run:
	swag init -g server/http/http.go
	go run $(LDGRPC) $(LDFLAGS) cmd/main.go -conf config/$(APP_NAME).toml