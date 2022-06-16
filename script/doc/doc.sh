#!/bin/bash
work_dir=$(
  cd "$(dirname "$0")" || exit
  pwd
)/../..

api_version="$1"
swagger_name="swagger"
gateway_name="gateway"
cd "${work_dir}/${gateway_name}/api/${api_version}" || exit
swag init -g "${gateway_name}".go -o "${work_dir}/${swagger_name}"/docs || exit

# TODO 先默认已经安装依赖
swag2md_name="swag2md"
#cd "${work_dir}/tools/${swag2md_name}" || exit
#echo "start building ${swag2md_name} tool"
#go build -o "${work_dir}/${swag2md_name}" || exit
#echo "building ${swag2md_name} tool success"

echo "start generating api.md"
"${swag2md_name}" \
  -t "即时通讯TxChat${api_version}接口文档" \
  -s "${work_dir}/${swagger_name}"/docs/swagger.json \
  -o "${work_dir}/api-${api_version}.md"
echo "generating api-${api_version}.md success"
