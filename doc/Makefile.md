# Makefile 使用说明

## 简介

本项目使用 `Makefile` 来管理各种常用的开发任务，包括构建、运行、测试和清理等操作。使用 Makefile 可以简化开发流程，提高开发效率。

## 主要功能

Makefile 提供了以下几个主要功能：

### 构建服务

```bash
make build
```
此命令将构建所有微服务，并将可执行文件输出到 `bin` 目录。

### 运行服务

运行单个服务：
```bash
# 运行用户服务
make run-user

# 运行社交服务
make run-social
```

运行所有服务：
```bash
make run-all
```

### 生成代码

```bash
make generate
```
此命令使用 `goctl` 工具从 proto 文件生成 gRPC 相关代码。

### 运行测试

```bash
make test
```
运行项目中的所有单元测试。

### 清理生成的文件

```bash
make clean
```
清理构建过程中生成的临时文件和可执行文件。

### 生成文档

```bash
make doc
```
生成项目的文档文件，输出到 `doc` 目录。

### 查看帮助

```bash
make help
```
显示所有可用的 Makefile 命令及其说明。

## 工作原理

Makefile 中的每个目标（target）都定义了一系列要执行的命令。例如，当运行 `make build` 时，系统会执行以下命令：

```makefile
build:
	@echo "构建所有服务..."
	go build -o bin/user app/user/rpc/user.go
	go build -o bin/social app/social/rpc/social.go
```

## 注意事项

1. 运行 Makefile 命令前，请确保已安装 Go 编程语言及相关工具。
2. 对于 `generate` 命令，需要先安装 `goctl` 工具。
3. 在 Windows 系统上，部分命令可能需要调整或使用 Git Bash 等工具来执行。

## 扩展与自定义

如需添加新的服务或功能，可以按照现有模式扩展 Makefile。例如，添加新服务的构建和运行命令：

```makefile
# 构建新服务
build-newservice:
	go build -o bin/newservice app/newservice/rpc/newservice.go

# 运行新服务
run-newservice:
	cd app/newservice/rpc && go run newservice.go
```

然后更新 `build` 和 `run-all` 目标以包含新服务。