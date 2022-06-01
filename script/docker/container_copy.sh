#!/bin/bash
work_dir=$(
  cd "$(dirname "$0")" || exit
  pwd
)/../..

volumeName=$1
files=$2

# example: ./container_copy.sh txchat-answer-config "../../target/*"
# or: ./container_copy.sh txchat-answer-config "../../target/answer.toml"
docker container create --name dummy -v "${volumeName}":/root hello-world
echo "copy:"
for fname in ${files} ; do \
  echo "${fname##*/}"
  docker cp "${fname}" dummy:/root/
done
docker rm dummy

