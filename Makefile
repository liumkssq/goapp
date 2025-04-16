.PHONY: all build run-user run-social run-all generate test clean help doc list-services

# 默认目标
all: build

# 构建所有服务
build:
	@echo "构建所有服务..."
	go build -o bin/user app/user/rpc/user.go
	go build -o bin/social app/social/rpc/social.go


# 运行用户服务
run-user:
	@echo "启动用户服务..."
	cd app/user/rpc && go run user.go

# 运行社交服务
run-social:
	@echo "启动社交服务..."
	cd app/social/rpc && go run social.go

# 列出所有服务
list-services:
	@echo "可用的服务列表:"
	@powershell -Command "Get-ChildItem -Path app -Recurse -Filter '*.go' | Where-Object { Select-String -Path $_.FullName -Pattern 'func main' -Quiet } | ForEach-Object { $_.Directory.FullName.Replace('$(shell pwd)', '.') }"

# 运行所有服务 (适用于Windows PowerShell)
run-all:
	@powershell -Command "[Console]::OutputEncoding = [System.Text.Encoding]::UTF8; Write-Host '启动所有服务...'"
	@powershell -Command "$ErrorActionPreference = 'SilentlyContinue'; $services = @('app/user/rpc', 'app/social/rpc', 'app/im/rpc', 'app/product/rpc', 'app/article/rpc', 'app/search/rpc', 'app/lostfound/rpc', 'app/ai/rpc', 'app/map/rpc', 'app/upload/rpc', 'app/task/rpc', 'app/gateway'); foreach ($$dir in $$services) { if (Test-Path "$$dir/*.go") { Start-Process -NoNewWindow powershell -ArgumentList '-Command', \"cd $$dir; go run *.go\" } }"
	@echo "所有服务已启动"

# 生成gRPC代码
generate:
	@echo "生成gRPC代码..."
	goctl rpc protoc app/user/rpc/user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
	goctl rpc protoc app/social/rpc/social.proto --go_out=. --go-grpc_out=. --zrpc_out=.

# 运行测试
test:
	@echo "运行测试..."
	go test -v ./...

# 清理生成的文件
clean:
	@echo "清理生成的文件..."
	@powershell -Command "if (Test-Path bin) { Remove-Item -Recurse -Force bin }"
	@powershell -Command "Get-ChildItem -Path . -Filter '*.go.gen' -Recurse | Remove-Item -Force"

# 生成项目文档
doc:
	@echo "生成项目文档..."
	@powershell -Command "if (-not (Test-Path doc)) { New-Item -ItemType Directory -Path doc }"
	@powershell -Command "if (-not (Test-Path doc\modules)) { New-Item -ItemType Directory -Path doc\modules }"
	@echo "文档生成完成，请查看doc目录"

# 帮助信息
help:
	@echo "Makefile 命令帮助:"
	@echo "make               - 构建所有服务"
	@echo "make build         - 构建所有服务"
	@echo "make run-user      - 运行用户服务"
	@echo "make run-social    - 运行社交服务"
	@echo "make run-all       - 运行所有服务"
	@echo "make list-services - 列出所有可用服务"
	@echo "make generate      - 生成gRPC代码"
	@echo "make test          - 运行测试"
	@echo "make clean         - 清理生成的文件"
	@echo "make doc           - 生成项目文档"
	@echo "make help          - 显示帮助信息"