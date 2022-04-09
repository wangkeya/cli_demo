开发命令行工具指导
===============

## 1.新建cli_demo项目
```text
# 安装cobra-cli工具
go install github.com/spf13/cobra-cli@latest

# 初始化
go mod init cli_demo && cobra-cli init
```


## 2.添加COMMAND

```text
# 添加新的COMMAND
cobra-cli add version
```

## 3.定制FLAG

```text
# 添加新的COMMAND
cobra-cli add hello

```

## 4.子命令

```text
# 添加新的COMMAND
cobra-cli add get
# 为hello命令添加子命令
cobra-cli add pods -p getCmd

```