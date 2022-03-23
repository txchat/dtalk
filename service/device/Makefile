# golang1.9 or latest
# 1. make help
# 2. make dep
# 3. make build
# ...

VERSION := $(shell echo $(shell cat cmd/main.go | grep "projectVersion =" | cut -d '=' -f2))
APP_NAME := device
BUILD_DIR := build
APP := ${BUILD_DIR}/${APP_NAME}_v${VERSION}
PKG_NAME := ${APP_NAME}_v${VERSION}
PKG := ${PKG_NAME}.tar.gz

main_path 	= "main"
go_version 	= $(shell go version | awk '{ print $3 }')
build_time 	= $(shell date "+%Y-%m-%d %H:%M:%S %Z")
git_commit 	= $(shell git rev-parse --short=10 HEAD)
flags		:= -ldflags "-X '${main_path}.goVersion=${go_version}' \
-X '${main_path}.buildTime=${build_time}' \
-X '${main_path}.gitCommit=${git_commit}' \
-X 'google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn'"

.PHONY: clean build pkg

clean: ## Remove previous build
	@rm -rf ${BUILD_DIR}
	@go clean

build: #checkgofmt ## Build the binary file
	GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY=https://goproxy.cn,direct GOSUMDB="sum.golang.google.cn" go build -v $(flags) -o $(APP) cmd/main.go

pkg: build ## Package
	mkdir -p ${PKG_NAME}/bin
	mkdir -p ${PKG_NAME}/etc
	cp ${APP} ${PKG_NAME}/bin/
	cp etc/*  ${PKG_NAME}/etc/
	tar zvcf ${PKG} ${PKG_NAME}
	rm -rf ${PKG_NAME}

REMOTE_BIN_PATH := /opt/dtalk/srv/app/bin
upload: build
	rsync -r $(APP) 107:$(REMOTE_BIN_PATH)
	ssh 107 "cd $(REMOTE_BIN_PATH) && chmod 777 $(PKG_NAME) && ln -sf $(PKG_NAME) $(APP_NAME) && supervisorctl restart $(APP_NAME)"