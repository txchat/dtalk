#!/bin/bash
# shellcheck disable=SC2034
work_dir=$(
    cd "$(dirname "$0")" || exit
    pwd
)/../..

created_volume=()
created_network=()

serviceList=$1
projectVersion=$2

# platform adaptation
HOST_OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case ${HOST_OS} in
    "darwin")
        function dosed() {
            sed -r -i '' "$@"
        }
        ;;
    *)
        function dosed() {
            sed -r -i "$@"
        }
        ;;
esac

function network_create() {
    networkName=$1
    filterName=$(docker network ls | awk 'NR>1{ print $2 }' | grep -w "${networkName}")
    if [ "$filterName" == "" ]; then
        #不存在就创建
        created_network+=("$networkName")
        docker network create "$networkName"
        echo "$networkName network created"
    fi
}

function volume_create() {
    volumeName=$1
    filterName=$(docker volume ls | awk 'NR>1{ print $2 }' | grep -w "${volumeName}")
    if [ "$filterName" == "" ]; then
        #不存在就创建
        created_volume+=("$volumeName")
        docker volume create "$volumeName"
        echo "$volumeName volume created"
    fi
}

function initRedis() {
    for vname in ${created_volume[*]}; do
        if [ "${vname}" = "txchat-redis-config" ]; then
            docker container create --name dummy -v "txchat-redis-config":/root hello-world
            docker cp redis.conf dummy:/root/redis.conf
            docker rm dummy
        fi
    done
}

function initMySQL() {
    # shellcheck disable=SC2048
    for vname in ${created_volume[*]}; do
        if [ "${vname}" = "txchat-mysql-init" ]; then
            echo "starting init MySQL"
            # 将初始化sql文件传入mysql初始化卷中
            docker container create --name dummy -v "txchat-mysql-init":/root hello-world
            docker cp dtalk_biz.sql dummy:/root/dtalk_biz.sql
            docker cp dtalk_record.sql dummy:/root/dtalk_record.sql
            docker rm dummy
        fi
    done
}

function initNginx() {
    # shellcheck disable=SC2048
    for vname in ${created_volume[*]}; do
        if [ "${vname}" = "txchat-nginx-config" ]; then
            echo "starting init Nginx"
            docker container create --name dummy -v "txchat-nginx-config":/root hello-world
            docker cp conf.d/dtalk.conf dummy:/root/dtalk.conf
            docker cp conf.d/dtalk_pprof.conf dummy:/root/dtalk_pprof.conf
            docker rm dummy
        fi
    done
}

volumes=("txchat-zookeeper-data" "txchat-kafka-data" "txchat-redis-data" "txchat-redis-config" "txchat-redis-log" "txchat-etcd-data" "txchat-mysql-data" "txchat-mysql-config" "txchat-mysql-log" "txchat-mysql-init" "txchat-minio-data" "txchat-nginx-config" "txchat-nginx-log" "txchat-prometheus-config" "txchat-grafana-data" "txchat-grafana-provisioning")
networks=("txchat-components" "txchat-service")

for service in ${serviceList}; do
    sDir=$(dirname "${service}")
    sName=$(basename "${service}")
    volumes+=("txchat-${sName}-config")
    # 将「-」转为「_」并将小写转大写
    upperSName=$(echo "${sName//[-]/_}" | tr '[:lower:]' '[:upper:]')
    # 修改.env文件服务镜像版本号
    dosed "s/(${upperSName}_IMAGE=)\s*(.+)/\1${projectVersion}/" .env
done

# shellcheck disable=SC2048
for vname in ${volumes[*]}; do
    volume_create "${vname}"
done

# shellcheck disable=SC2048
for name in ${networks[*]}; do
    network_create "${name}"
done

initRedis
initMySQL
initNginx
