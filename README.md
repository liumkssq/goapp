# 校园二手交易平台 API 和 Protocol Buffers 文件

本项目包含基于 go-zero 框架的校园二手交易平台的 API 定义和 Protocol Buffers 接口定义。

## 项目结构

```
goapp/
├── user/                   # 用户服务
│   ├── user.api            # HTTP API 接口定义
│   └── user.proto          # gRPC 接口定义
├── im/                     # 即时通讯服务
│   ├── im.api              # HTTP API 接口定义
│   ├── im.proto            # gRPC 接口定义
│   └── websocket.api       # WebSocket 接口定义
├── lostfound/              # 失物招领服务
│   ├── lostfound.api       # HTTP API 接口定义
│   └── lostfound.proto     # gRPC 接口定义
├── product/                # 商品服务
│   ├── product.api         # HTTP API 接口定义
│   └── product.proto       # gRPC 接口定义
├── article/                # 文章服务
│   ├── article.api         # HTTP API 接口定义
│   └── article.proto       # gRPC 接口定义
├── search/                 # 搜索服务
│   ├── search.api          # HTTP API 接口定义
│   └── search.proto        # gRPC 接口定义
├── ai/                     # AI 辅助服务
│   ├── ai.api              # HTTP API 接口定义
│   └── ai.proto            # gRPC 接口定义
├── map/                    # 地图服务
│   ├── map.api             # HTTP API 接口定义
│   └── map.proto           # gRPC 接口定义
├── upload/                 # 上传服务
│   ├── upload.api          # HTTP API 接口定义
│   └── upload.proto        # gRPC 接口定义
└── gateway/                # API 网关服务
    └── gateway.api         # 网关配置定义
```

## 服务说明

### 用户服务 (User)

用户服务提供用户认证、信息管理等功能，包括：
- 用户注册、登录、登出
- 用户信息管理
- 密码重置
- 验证码发送
- 用户通知
- 用户关注

### 即时通讯服务 (IM)

即时通讯服务提供实时消息功能，包括：
- WebSocket 连接管理
- 用户好友管理
- 消息发送与接收
- 群组管理
- 会话管理
- 在线状态管理

### 失物招领服务 (LostFound)

失物招领服务提供校园失物招领功能，包括：
- 失物/招领信息发布
- 信息浏览和搜索
- 信息状态管理
- 评论和点赞
- 用户失物招领管理

### 商品服务 (Product)

商品服务提供二手商品交易功能，包括：
- 商品发布、编辑、删除
- 商品分类管理
- 商品搜索和浏览
- 商品收藏
- 评论和举报

### 文章服务 (Article)

文章服务提供校园社区文章功能，包括：
- 文章发布、编辑、删除
- 文章分类管理
- 文章搜索和浏览
- 文章点赞和收藏
- 评论管理

### 搜索服务 (Search)

搜索服务提供全局搜索和专项搜索功能，包括：
- 全局搜索
- 商品搜索
- 文章搜索
- 失物招领搜索
- 用户搜索
- 搜索历史和热门搜索

### AI 辅助服务 (AI)

AI 辅助服务提供智能功能增强，包括：
- 商品图片分析
- 内容生成
- 标签生成
- 价格估算
- 智能对话
- 文本生成
- 图像生成
- 图像识别
- 文本翻译

### 地图服务 (Map)

地图服务提供位置相关功能，包括：
- 地址解析和逆地址解析
- 位置搜索
- 周边地点检索
- 距离计算
- 用户地址管理

### 上传服务 (Upload)

上传服务提供文件上传功能，包括：
- 图片上传
- 文件上传
- 头像上传
- URL 图片上传

### API 网关 (Gateway)

API 网关服务作为所有微服务的统一入口，提供：
- 路由分发
- 认证授权
- 请求转发

## 使用说明

### 生成 API 服务代码

使用 goctl 工具生成 API 服务代码：

```bash
goctl api go -api <path-to-api-file> -dir <output-directory>
```

例如：

```bash
goctl api go -api user/user.api -dir ./gen/user
```

### 生成 RPC 服务代码

使用 goctl 工具生成 RPC 服务代码：

```bash
goctl rpc protoc <path-to-proto-file> --go_out=<output-directory> --go-grpc_out=<output-directory> --zrpc_out=<output-directory>
```

例如：

```bash
goctl rpc protoc user/user.proto --go_out=./gen/user --go-grpc_out=./gen/user --zrpc_out=./gen/user
```

### 实现业务逻辑

1. 在生成的服务代码中实现业务逻辑
2. 启动服务
3. 通过网关统一访问各服务

## 部署建议

- 使用 Docker 容器化每个微服务
- 使用 Kubernetes 编排服务
- 使用 Etcd 进行服务发现
- 使用 Redis 进行缓存
- 使用 MySQL 或 PostgreSQL 作为主数据库
- 使用 MongoDB 存储非结构化数据

## 开发规范

1. API 命名遵循 RESTful 设计风格
2. 函数和变量命名采用驼峰式
3. 注释遵循 Golang 标准
4. 统一的错误处理
5. 代码提交前进行测试

## 重要说明

1. API 接口需要进行适当的权限控制
2. 敏感操作需要记录日志
3. 接口应实现限流防止恶意请求
4. 文件上传需限制大小和类型
5. 用户密码需要加密存储