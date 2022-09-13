# kafka 消费者和生产者测试

## 编译
在当前目录执行：
```shell
make build
```

##### 批量生产和消费
```shell
# 运行消费者
./consumer --topic test-topic --group test-group --broker "127.0.0.1:9092" --cons 8 --pros 8

# 运行生产者
./producer --topic test-topic --broker "127.0.0.1:9092" --number 16
```

##### 带缓冲情况下生产速度大于消费速度
```shell
# 启动消费者
./consumer --topic test-topic --group test-group --broker "127.0.0.1:9092" --cons 1 --pros 1

# 启动生产者
./producer --topic test-topic --broker "127.0.0.1:9092" --number 4

# 等待一定时间后，command+c 关闭生产者。消费者从缓冲队列消费数据

# 片刻后，command+c 关闭消费者，消费者进程等待缓冲队列全部消费后优雅退出
```

##### 无缓冲情况下生产速度大于消费速度
```shell
# 启动消费者
./consumer --topic test-topic --group test-group --broker "127.0.0.1:9092" --cons 1 --pros 1 --bcc 1 --ccc 1

# 启动生产者
./producer --topic test-topic --broker "127.0.0.1:9092" --number 4

# 等待一定时间后，command+c 关闭生产者。消费者从kafka broker消费数据

# 片刻后，command+c 关闭消费者，消费者进程立刻优雅退出
```

##### 消费者连接超时
```shell
./consumer --topic test-topic --group test-group --broker "127.0.0.1:9092" --cons 1 --pros 1 --bcc 1 --ccc 1 --cto 1
```