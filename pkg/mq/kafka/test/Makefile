# golang1.17 or latest
GO_OS_ARCH=$(shell go version | awk '{ print $$4 }')
HOST_ARCH=$(shell echo "${GO_OS_ARCH}" | awk -F '/' '{ print $$2 }')
HOST_OS=$(shell echo "${GO_OS_ARCH}" | awk -F '/' '{ print $$1 }')
GO_ENV_BASE=GO111MODULE=on GOPROXY=https://goproxy.cn,direct GOSUMDB="sum.golang.google.cn"
GO_ENV=GOOS=$(HOST_OS) GOARCH=$(HOST_ARCH) ${GO_ENV_BASE}

buildDir := target
Consumer := "$(buildDir)"/consumer
ConsumerMain := consumer/main.go
Producer := "$(buildDir)"/producer
ProducerMain := producer/main.go

MockConsumer := "$(buildDir)"/mockconsumer
MockConsumerMain := mock/cmd/main.go

help: ## Display this help screen
	@printf "Help doc:\nUsage: make [command]\n"
	@printf "[command]\n"
	@grep -h -E '^([a-zA-Z_-]|\%)+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## 编译本机系统和指令集的可执行文件
	$(GO_ENV) go build -v -o $(Consumer) $(ConsumerMain)
	$(GO_ENV) go build -v -o $(Producer) $(ProducerMain)
	$(GO_ENV) go build -v -o $(MockConsumer) $(MockConsumerMain)

build_%: ## 编译目标机器的可执行文件（例如: make build_linux_amd64）
	TAR_OS=$(shell echo $* | awk -F'_' '{print $$1}'); \
	TAR_ARCH=$(shell echo $* | awk -F'_' '{print $$2}'); \
	GOOS=$${TAR_OS} GOARCH=$${TAR_ARCH} $(GO_ENV_BASE) go build -v -o $(Consumer)_"$${TAR_OS}"_"$${TAR_ARCH}" $(ConsumerMain);\
	GOOS=$${TAR_OS} GOARCH=$${TAR_ARCH} $(GO_ENV_BASE) go build -v -o $(Producer)_"$${TAR_OS}"_"$${TAR_ARCH}" $(ProducerMain)
