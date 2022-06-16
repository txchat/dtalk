# golang1.17 or latest
TARGETDIR=target

projectVersion=$(shell git describe --abbrev=8 --tags)
gitCommit=$(shell git rev-parse --short=8 HEAD)

pkgCommitName=${projectVersion}_${gitCommit}
servers=answer backend backup call device discovery gateway generator group offline-push oss pusher store

help: ## Display this help screen
	@printf "Help doc:\nUsage: make [command]\n"
	@printf "[command]\n"
	@grep -h -E '^([a-zA-Z_-]|\%)+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: clean ## 编译本机系统和指令集的可执行文件
	./script/builder/builder.sh ${TARGETDIR} "" "${servers}"

build_%: clean ## 编译目标机器的可执行文件（例如: make build_linux_amd64）
	./script/builder/builder.sh ${TARGETDIR} $* "${servers}"

pkg: build ## 编译并打包本机系统和指令集的可执行文件
	tar -zcvf ${TARGETDIR}_'host'_${pkgCommitName}.tar.gz ${TARGETDIR}/

pkg_%: build_% ## 编译并打包目标机器的可执行文件（例如: make pkg_linux_amd64）
	tar -zcvf ${TARGETDIR}_$*_${pkgCommitName}.tar.gz ${TARGETDIR}/

images: build_linux_amd64 ## 打包docker镜像
	cp script/docker/*Dockerfile ${TARGETDIR}
	cd ${TARGETDIR} && for i in $(servers) ; do \
		docker build --build-arg server_name=$$i . -f server.Dockerfile -t txchat-$$i:${projectVersion}; \
	done

init-compose: images ## 使用docker compose启动
	cp -R script/compose/. run_compose/
	cp -R script/mysql/. run_compose/
	cp -R script/nginx/. run_compose/
	cd run_compose && \
	./envfill.sh;\
	./initwork.sh "${servers}" "${projectVersion}"

docker-compose-up: ## 使用docker compose启动
	@if [ ! -d "run_compose/" ]; then \
		exit -1;\
     fi; \
	cd run_compose && \
	docker-compose -f components.compose.yaml -f service.compose.yaml up -d

docker-compose-%: ## 使用docker compose 命令(服务列表：make docker-compose-ls；停止服务：make docker-compose-stop；卸载服务：make docker-compose-down)
	@if [ ! -d "run_compose/" ]; then \
       cp -R script/compose/. run_compose/; \
     fi; \
    cd run_compose && \
    docker-compose -f components.compose.yaml -f service.compose.yaml $*

.PHONY: doc
doc:
	./script/doc/doc.sh v1

test:
	$(GOENV) go test -v ./...

clean:
	rm -rf ${TARGETDIR}