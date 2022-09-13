# kafka 消费者和生产者测试

## 编译
在当前目录执行：
```shell
make build
```


运行消费者：
```shell
./consumer --topic test-topic --group test-group --broker "127.0.0.1:9092" --cons 8 --pros 8
```

运行生产者：
```shell
./producer --topic test-topic --broker "127.0.0.1:9092" --number 16
```