#!/bin/bash
work_dir=$(
    cd "$(dirname "$0")" || exit
    pwd
)/../../../

cp -R ./ "${work_dir}/test_compose/"
cp -R "${work_dir}/script/mysql/." "${work_dir}/test_compose/"
cp -R "${work_dir}/script/nginx/." "${work_dir}/test_compose/"
cp -R "${work_dir}/script/redis/." "${work_dir}/test_compose/"
cp -R "${work_dir}/script/prometheus/." "${work_dir}/test_compose/"

chmod +x "${work_dir}/test_compose/startup.sh"
chmod +x "${work_dir}/test_compose/shutdown.sh"
