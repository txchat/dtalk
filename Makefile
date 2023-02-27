# golang1.17 or latest
TARGETDIR=target

projectVersion=$(shell git describe --abbrev=8 --tags)
gitCommit=$(shell git rev-parse --short=8 HEAD)

pkgCommitName=${projectVersion}_${gitCommit}
servers=backup call device generator group offline oss pusher storage transfer version
gateways=center chat

help: ## Display this help screen
	@printf "Help doc:\nUsage: make [command]\n"
	@printf "[command]\n"
	@grep -h -E '^([a-zA-Z_-]|\%)+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: clean ## 编译本机系统和指令集的可执行文件
	./script/builder/builder.sh ${TARGETDIR} "" "${servers}" "${gateways}"

build_%: clean ## 编译目标机器的可执行文件（例如: make build_linux_amd64）
	./script/builder/builder.sh ${TARGETDIR} $* "${servers}" "${gateways}"

pkg: build ## 编译并打包本机系统和指令集的可执行文件
	tar -zcvf ${TARGETDIR}_'host'_${pkgCommitName}.tar.gz ${TARGETDIR}/

pkg_%: build_% ## 编译并打包目标机器的可执行文件（例如: make pkg_linux_amd64）
	tar -zcvf ${TARGETDIR}_$*_${pkgCommitName}.tar.gz ${TARGETDIR}/

images: build_linux_amd64 ## 打包docker镜像
	cp script/docker/*Dockerfile ${TARGETDIR}
	cd ${TARGETDIR} && for i in $(servers) ; do \
		docker build --build-arg server_name=$$i . -f server.Dockerfile -t txchat-$$i:${projectVersion}; \
	done && for i in $(gateways) ; do \
		docker build --build-arg server_name=$$i . -f gateway.Dockerfile -t txchat-$$i:${projectVersion}; \
	done

init-compose: images ## 使用docker compose启动
	cp -R script/compose/. run_compose/
	cp -R script/mysql/. run_compose/
	cp -R script/nginx/. run_compose/
	cd run_compose && \
	./envfill.sh;\
	./initwork.sh "${servers} ${gateways}" "${projectVersion}"

docker-compose-up: ## 使用docker compose启动
	@if [ ! -d "run_compose/" ]; then \
		exit -1;\
     fi; \
	cd run_compose && \
	docker compose -f components.compose.yaml -f service.compose.yaml up -d

docker-compose-%: ## 使用docker compose 命令(服务列表：make docker-compose-ls；停止服务：make docker-compose-stop；卸载服务：make docker-compose-down)
	@if [ ! -d "run_compose/" ]; then \
       cp -R script/compose/. run_compose/; \
     fi; \
    cd run_compose && \
    docker compose -f components.compose.yaml -f service.compose.yaml $*

test-init:
	cp -R script/test/components/. test_compose/
	cp -R script/mysql/. test_compose/
	cp -R script/nginx/. test_compose/

test-up:
	@if [ ! -d "test_compose/" ]; then \
		exit -1;\
	 fi; \
	cd test_compose && \
	docker compose -f components.compose.yaml up -d

test-%:
	@if [ ! -d "test_compose/" ]; then \
       cp -R script/compose/. test_compose/; \
     fi; \
    cd test_compose && \
    docker compose -f components.compose.yaml $*

test:
	./script/test/unitestenv/run.sh

clean:
	rm -rf ${TARGETDIR}

.PHONY: doc swagger
doc:
	# ./script/doc/doc.sh v1
	goctl -v || GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest \
&& goctl api doc --dir app/gateway/center --o docs/api/center && goctl api doc --dir app/gateway/chat --o docs/api/chat

swagger:
	goctl -v || GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest \
&& goctl-swagger -v || GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest \
&& goctl api plugin -plugin goctl-swagger="swagger -filename center.json" -api app/gateway/center/center.api -dir . \
$$ goctl api plugin -plugin goctl-swagger="swagger -filename chat.json" -api app/gateway/chat/chat.api -dir .

.PHONY: fmt_proto fmt_shell fmt_go

fmt: fmt_proto fmt_shell fmt_go ## 文件格式化

fmt_proto: ## protobuf文件格式化
	@find . -name '*.proto' -not -path "./vendor/*" | xargs clang-format -i

fmt_shell: ## shell文件格式化
	@find . -name '*.sh' -not -path "./vendor/*" | xargs shfmt -w -s -i 4 -ci -bn

fmt_go: ## go源码格式化
	@find . -name '*.go' -not -path "./vendor/*" | xargs gofmt -s -w
	@find . -name '*.go' -not -path "./vendor/*" | xargs goimports -l -w

.PHONY: checkgofmt linter linter_test

check: checkgofmt linter ## check format and linter

checkgofmt: ## get all go files and run go fmt on them
	@files=$$(find . -name '*.go' -not -path "./vendor/*" | xargs gofmt -l -s); if [ -n "$$files" ]; then \
		  echo "Error: 'make fmt' needs to be run on:"; \
		  find . -name '*.go' -not -path "./vendor/*" | xargs gofmt -l -s ;\
		  exit 1; \
		  fi;
	@files=$$(find . -name '*.go' -not -path "./vendor/*" | xargs goimports -l -w); if [ -n "$$files" ]; then \
		  echo "Error: 'make fmt' needs to be run on:"; \
		  find . -name '*.go' -not -path "./vendor/*" | xargs goimports -l -w ;\
		  exit 1; \
		  fi;

linter: ## Use gometalinter check code, ignore some unserious warning
	@golangci-lint run ./... && find . -name '*.sh' -not -path "./vendor/*" | xargs shellcheck

linter_test: ## Use gometalinter check code, for local test
	@chmod +x ./script/golinter.sh
	@./script/golinter.sh "test" "${p}"
	@find . -name '*.sh' -not -path "./vendor/*" | xargs shellcheck
