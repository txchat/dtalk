#!/bin/bash
work_dir=$(
    cd "$(dirname "$0")" || exit
    pwd
)/../..

REPOSITORY=$1
TAG=$2

# check image
targetLen=$(docker images | grep -w "${REPOSITORY}" | grep -w "${TAG}" | awk '{ print length($0) }')

if [ "${targetLen}" != "" ]; then
    exit 0
else
    docker pull "${REPOSITORY}:${TAG}"
    exit $?
fi
