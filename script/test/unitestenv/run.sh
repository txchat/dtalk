#!/bin/bash
work_dir=$(
    cd "$(dirname "$0")" || exit
    pwd
)/../../../

cp -R "${work_dir}/script/test/unitestenv/" "${work_dir}/test_uni/"
cp -R "${work_dir}/script/mysql/." "${work_dir}/test_uni/"
cp -R "${work_dir}/script/nginx/." "${work_dir}/test_uni/"

chmod +x "${work_dir}/test_uni/wait-for-it.sh"

if [ ! -d "${work_dir}/test_uni/" ]; then
    echo "runtime file not exists"
    exit 1
fi

envOsArch=$(go version | awk '{ print $4 }')
targetOS=$(echo "${envOsArch}" | awk -F '/' '{ print $1 }')
targetARCH=$(echo "${envOsArch}" | awk -F '/' '{ print $2 }')

exportGOEnv() {
    export GOOS=${targetOS}
    export GOARCH=${targetARCH}
    export CGO_ENABLED=0
    export GO111MODULE=on
    export GOPROXY=https://goproxy.cn,direct
    export GOSUMDB='sum.golang.google.cn'
}

source .env

exportComponentEnv() {
    export MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
}

exportGOEnv
exportComponentEnv

#docker compose -f components.compose.yaml up && $(GOENV) go test -v "${work_dir}"/...
#docker compose -f components.compose.yaml down

if cd "${work_dir}/test_uni/" && docker compose -f components.compose.yaml up -d && ./wait-for-it.sh; then
    go test -v "${work_dir}"/app/services/storage/internal/dao/...
fi

#go test -v "${work_dir}"/...

# shutdown
#docker compose -f components.compose.yaml down
