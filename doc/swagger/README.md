# Swagger API 文档使用说明

本目录包含了校园二手交易平台所有微服务的Swagger API文档。

## 文档结构

```
swagger/
├── user/                # 用户服务Swagger文档
│   └── swagger.json
├── im/                  # 即时通讯服务Swagger文档
│   ├── swagger.json
│   └── websocket-swagger.json
├── lostfound/           # 失物招领服务Swagger文档
│   └── swagger.json
├── product/             # 商品服务Swagger文档
│   └── swagger.json
├── article/             # 文章服务Swagger文档
│   └── swagger.json
├── search/              # 搜索服务Swagger文档
│   └── swagger.json
├── ai/                  # AI辅助服务Swagger文档
│   └── swagger.json
├── map/                 # 地图服务Swagger文档
│   └── swagger.json
├── upload/              # 上传服务Swagger文档
│   └── swagger.json
├── gateway/             # 网关服务Swagger文档
│   └── swagger.json
└── aggregated-swagger.json # 聚合所有服务的API文档
```

## 使用Docker查看Swagger文档

### 启动Swagger UI

在项目根目录下执行以下命令启动所有服务的Swagger UI：

```bash
cd deploy/docker
docker-compose -f swagger-ui.yml up -d
```

### 访问Swagger UI

服务启动后，可以通过以下URL访问各个服务的Swagger UI：

- 所有服务API聚合: http://localhost:8900
- 用户服务: http://localhost:8901
- 即时通讯服务: http://localhost:8902
- 失物招领服务: http://localhost:8903
- 商品服务: http://localhost:8904
- 文章服务: http://localhost:8905
- 搜索服务: http://localhost:8906
- AI辅助服务: http://localhost:8907
- 地图服务: http://localhost:8908
- 上传服务: http://localhost:8909

## 手动查看Swagger文档

也可以通过以下命令单独启动某个服务的Swagger UI：

```bash
# 例如，单独启动用户服务的Swagger UI
docker run --rm -p 8080:8080 -e SWAGGER_JSON=/foo/swagger.json -v $(pwd)/user:/foo swaggerapi/swagger-ui
```

## 修改和更新Swagger文档

如果API定义有更新，需要重新生成Swagger文档。可以使用以下命令：

```bash
# 安装goctl-swagger插件
go install github.com/zeromicro/goctl-swagger@latest

# 生成Swagger文档
goctl apis plugin -plugin goctl-swagger="swagger -filename swagger.json" -apis <apis-file-path> -dir <output-directory>
```

## 使用Swagger UI测试API

Swagger UI不仅可以查看API文档，还可以直接测试API。在Swagger UI界面中：

1. 点击要测试的API端点
2. 点击"Try it out"按钮
3. 填写请求参数
4. 点击"Execute"执行请求

## 注意事项

1. 需要登录的API需要在Swagger UI界面顶部点击"Authorize"按钮，输入Bearer Token进行身份验证
2. 上传文件的API需要使用multipart/form-data格式
3. WebSocket API不能直接在Swagger UI中测试，请参考WebSocket文档使用专门的WebSocket客户端进行测试