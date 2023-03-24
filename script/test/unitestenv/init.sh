#!/bin/bash
work_dir=$(
    cd "$(dirname "$0")" || exit
    pwd
)/../../../

cp -R "${work_dir}/script/test/unitestenv/" "${work_dir}/test_uni/"
cp -R "${work_dir}/script/mysql/." "${work_dir}/test_uni/"
cp -R "${work_dir}/script/nginx/." "${work_dir}/test_uni/"

chmod +x "${work_dir}/test_uni/wait-for-it.sh"
exit 0
