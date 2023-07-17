# goctl通用生成模板

## 模板特性
* 无需定义返回结构体api
* 返回json响应
* logic获取到请求头
* logic 通过success和error方法输出响应

## 快速开始

### http开发
```shell
# 创建工作空间并进入该目录
$ mkdir -p service/youapp/api && cd service/youapp/api
goctl api -o test.api
goctl api go -api test.api -dir .  --home /Users/return/workspace/mygo/zerocmf/template
```

### rpc开发
```shell
# 创建工作空间并进入该目录
$ mkdir -p service/youapp/rpc && cd service/youapp/rpc
goctl api -o test.proto
goctl goctl rpc protoc test.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --home /Users/return/workspace/mygo/zerocmf/template
```