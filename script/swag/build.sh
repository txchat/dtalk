#!/bin/bash
file_path=$(
    cd "$(dirname "$0")" || exit
    pwd
)/../..

echo $file_path

#projectPath = $(dirname $(dirname "$PWD"))

if [ -z $1 ]; then
    echo 'ERROR: undefined version
    please input gateway version, example: ./build.sh v1'
else
    swag init -d $file_path/gateway/api/$1/ -g internal/handler/routes.go -o $file_path/gateway/api/$1/docs/
fi
