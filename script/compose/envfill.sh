#!/bin/bash
work_dir=$(
  cd "$(dirname "$0")" || exit
  pwd
)/../..

function randomPassword() {
  MYSQL_ROOT_PASSWORD=$(openssl rand -base64 16)
  MINIO_ROOT_PASSWORD=$(openssl rand -base64 16)
}

source key 2>/dev/null || randomPassword

eval "cat <<EOF
$(<env_tmpl)
EOF
" 1 >.env 2> /dev/null