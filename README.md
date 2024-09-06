# PCB-华为云token管理服务

> Powered by [Function graph](https://www.huaweicloud.com/product/functiongraph.html)  of Huawei cloud

## Pre

prepare the `.env` file in the root dictionary:

```bash
CONFIG_APPMODE=release     # app mode: release or debug
CONFIG_PORT=:8000          # server port
CONFIG_MYSQL_ADDRESS=###   # Mysql url
CONFIG_MYSQL_PORT=###      # Mysql port 
CONFIG_MYSQL_USER=###      # Mysql username
CONFIG_MYSQL_PASS=###      # Mysql password
CONFIG_MYSQL_DBNAME=###    # Mysql database name
CONFIG_HUAWEI_USER=###     # huawei cloud username
CONFIG_HUAWEI_IAMUSER=###  # IAM username
CONFIG_HUAWEI_IAMPASS=###  # IAM password
```

run the code

```
go run main.go
```

## Build

> [Tutorial written by huawei](https://support.huaweicloud.com/intl/zh-cn/bestpractice-functiongraph/functiongraph_05_1168.html)

configure the linux:amd64 environment

```bash
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux
```

build
```bash
# go build -o [目标可执行程序] [源程序]
# 
go build -o go-http-demo main.go
```

restore the environment configuration
```bash
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=windows
go env -w GOOS=windows
```