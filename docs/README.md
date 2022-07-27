## 配置文件

```toml
# Env 三种模式  
Env = "debug"      # 测试环境  
Env = "release"    # 线上环境  
Env = "benchmark"  # 压测环境
```

## 部署环境

环境需要增加以下环境变量  
GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn

## 快速部署

`make build_xxx` xxx 可选为 amd 或者 arm  
如果要选择平台则自行修改`dtalk/script/build/util.sh`中的`initOS()`编译设置

`make build` 包括编译二进制, 配置文件并打包

`make quick_build` 仅编译二进制