#!/bin/bash
work_dir=$(
    cd "$(dirname "$0")" || exit
    pwd
)/../..
# import .sh files

# $1:build target dir; $2:build target architecture(like linux_amd64); $3:service list
targetDir=$1
targetOsArch=$2
serviceList=$3
gatewayList=$4

if [ "${targetDir}" = "" ]; then
    echo "target directory is empty"
    exit 1
fi

golangOsArch() {
    if [ "$1" = "" ]; then
        envOsArch=$(go version | awk '{ print $4 }')
        targetOS=$(echo "${envOsArch}" | awk -F '/' '{ print $1 }')
        targetARCH=$(echo "${envOsArch}" | awk -F '/' '{ print $2 }')
    else
        # set target os and arch
        targetOS=$(echo "$1" | awk -F '_' '{ print $1 }')
        targetARCH=$(echo "$1" | awk -F '_' '{ print $2 }')
    fi

    if [ "${targetOS}" = "" ] || [ "${targetARCH}" = "" ]; then
        echo "error: got empty os or arch type!"
        exit 1
    fi
}

ldflagsOfInject() {
    versionFilePath="main"
    #projectVersion=$(git describe --abbrev=0 --tags)
    projectVersion=$(git describe --abbrev=8 --tags)
    goVersion=$(go version | awk '{ print $3 }')
    #goVersion=$(shell go version | awk '{ print $$3 }')
    gitCommit=$(git rev-parse --short=8 HEAD)
    #buildTime=`date +%Y%m%d`
    buildTime=$(date "+%Y-%m-%d %H:%M:%S %Z")
    osArch="${targetOS}/${targetARCH}"

    ldflags="\
  -X '${versionFilePath}.projectVersion=${projectVersion}' \
  -X '${versionFilePath}.goVersion=${goVersion}' \
  -X '${versionFilePath}.gitCommit=${gitCommit}' \
  -X '${versionFilePath}.buildTime=${buildTime}' \
  -X '${versionFilePath}.osArch=${osArch}' \
  -X 'google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn'"
}

exportGOEnv() {
    export GOOS=${targetOS}
    export GOARCH=${targetARCH}
    export CGO_ENABLED=0
    export GO111MODULE=on
    export GOPROXY=https://goproxy.cn,direct
    export GOSUMDB='sum.golang.google.cn'
}

golangOsArch "${targetOsArch}"
ldflagsOfInject
exportGOEnv

buildTargetDir="${work_dir}/${targetDir}/"

# $1:service name; $2: config root path; $3: service main filename
buildService() {
    serviceName="$1"
    configPath="$2"
    configName="$3"
    serviceMain="$4"

    echo "┌ start building ${serviceName} service"
    cp "${configPath}/${configName}" "${buildTargetDir}/${configName}"
    go build -ldflags "${ldflags}" -o "${buildTargetDir}/${serviceName}" "${serviceMain}" || exit
    echo "└ building ${serviceName} service success"
}

mkdir "${buildTargetDir}"
for service in ${serviceList}; do
    sDir=$(dirname "${service}")
    sName=$(basename "${service}")
    buildService "${sName}" "${work_dir}/app/services/${sDir}/${sName}/etc" "${sName}.yaml" "${work_dir}/app/services/${sDir}/${sName}/${sName}.go"
    #    echo "${sName}" "${work_dir}/app/services/${sDir}/${sName}/etc" "${sName}.yaml" "${work_dir}/app/services/${sDir}/${sName}/${sName}.go" "$(printf '\r')"
done

for gateway in ${gatewayList}; do
    sDir=$(dirname "${gateway}")
    sName=$(basename "${gateway}")
    buildService "${sName}" "${work_dir}/app/gateway/${sDir}/${sName}/etc" "${sName}-api.yaml" "${work_dir}/app/gateway/${sDir}/${sName}/${sName}.go"
    #    echo "${sName}" "${work_dir}/app/gateway/${sDir}/${sName}/etc" "${sName}-api.yaml" "${work_dir}/app/gateway/${sDir}/${sName}/${sName}.go" "$(printf '\r')"
done
