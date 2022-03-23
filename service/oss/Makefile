VERSION := $(shell echo $(shell cat cmd/main.go | grep "projectVersion =" | cut -d '=' -f2))
APP_NAME := oss
BUILD_DIR := build
APP := ${BUILD_DIR}/${APP_NAME}_v${VERSION}
PKG_NAME := ${APP_NAME}_v${VERSION}
PKG := ${PKG_NAME}.tar.gz

.PHONY: clean build pkg

main_path 	= "main"
go_version 	= $(shell go version | awk '{ print $3 }')
build_time 	= $(shell date "+%Y-%m-%d %H:%M:%S %Z")
git_commit 	= $(shell git rev-parse --short=10 HEAD)
flags		:= -ldflags "-X '${main_path}.goVersion=${go_version}' \
-X '${main_path}.buildTime=${build_time}' \
-X '${main_path}.gitCommit=${git_commit}' \
-X 'google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn'"

clean: ## Remove previous build
	@rm -rf ${BUILD_DIR}
	@go clean

swag:
	@echo '┌ start 	gen swag'
	@swag init -g server/http/http.go
	@echo '└ end 	gen swag'

build: swag
	@echo '┌ start	build $(APP)'
	@GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY=https://goproxy.cn,direct GOSUMDB="sum.golang.google.cn" go build -v $(flags) -o $(APP) cmd/main.go
	@echo '└ end	build $(APP)'

pkg: build ## Package
	mkdir -p ${PKG_NAME}/bin
	mkdir -p ${PKG_NAME}/etc
	cp ${APP} ${PKG_NAME}/bin/
	cp etc/*  ${PKG_NAME}/etc/
	tar zvcf ${PKG} ${PKG_NAME}
	rm -rf ${PKG_NAME}

run: swag
	@echo '┌ start	run $(APP)'
	@go run cmd/main.go -conf config/oss.toml
	@echo '└ end	run $(APP)'

REMOTE_BIN_PATH := /opt/dtalk/srv/app/bin
upload: build
	@echo '┌ start	upload $(APP)'
	@rsync -r $(APP) 107:$(REMOTE_BIN_PATH)
	@ssh 107 "cd $(REMOTE_BIN_PATH) && chmod 777 $(PKG_NAME) && ln -sf $(PKG_NAME) $(APP_NAME) && supervisorctl restart $(APP_NAME)"
	@echo '┌ start	upload $(APP)'