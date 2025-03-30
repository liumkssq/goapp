# 校园二手交易平台 API 和 Proto 文件

本项目包含了校园二手交易平台后端服务的 API 定义和 Protocol Buffers 文件。这些文件用于定义和实现基于 go-zero 框架的微服务架构。

## 项目结构

```
goapp/
├── user/                  # 用户服务
│    ├── user.api          # 用户服务 API 定义
│    └── user.proto        # 用户服务 Protocol Buffers 定义
├── im/                    # 即时通讯服务
│    ├── im.api            # 即时通讯服务 API 定义
│    ├── websocket.api     # WebSocket API 定义
│    └── im.proto          # 即时通讯服务 Protocol Buffers 定义
├── lostfound/             # 失物招领服务
│    ├── lostfound.api     # 失物招领服务 API 定义
│    └── lostfound.proto   # 失物招领服务 Protocol Buffers 定义
├── product/               # 商品服务
│    ├── product.api       # 商品服务 API 定义
│    └── product.proto     # 商品服务 Protocol Buffers 定义
├── article/               # 文章服务
│    ├── article.api       # 文章服务 API 定义
│    └── article.proto     # 文章服务 Protocol Buffers 定义
├── search/                # 搜索服务
│    ├── search.api        # 搜索服务 API 定义
│    └── search.proto      # 搜索服务 Protocol Buffers 定义
├── ai/                    # AI 辅助服务
│    ├── ai.api            # AI 辅助服务 API 定义
│    └── ai.proto          # AI 辅助服务 Protocol Buffers 定义
├── map/                   # 地图服务
│    ├── map.api           # 地图服务 API 定义
│    └── map.proto         # 地图服务 Protocol Buffers 定义
├── upload/                # 上传服务
│    ├── upload.api        # 上传服务 API 定义
│    └── upload.proto      # 上传服务 Protocol Buffers 定义
└── gateway/               # API 网关
     └── gateway.api       # API 网关配置
```

## 服务说明

### 用户服务 (User)

用户服务负责处理用户注册、登录、信息管理等功能。主要接口包括：

- 用户注册
- 用户登录（账号密码登录、验证码登录）
- 获取/更新用户信息
- 关注/取消关注用户
- 获取通知列表

### 即时通讯服务 (IM)

即时通讯服务负责处理用户之间的实时通信，包括私聊和群聊功能。主要接口包括：

- WebSocket 连接管理
- 发送/接收消息
- 管理联系人和群组
- 获取聊天历史记录

### 失物招领服务 (LostFound)

失物招领服务用于处理校园内的失物招领信息发布和管理。主要接口包括：

- 发布/更新/删除失物招领信息
- 获取失物招领列表和详情
- 评论功能
- 状态更新（已找到、已认领等）

### 商品服务 (Product)

商品服务用于处理二手商品的发布和交易。主要接口包括：

- 发布/更新/删除商品
- 获取商品列表和详情
- 商品分类管理
- 收藏功能

### 文章服务 (Article)

文章服务用于管理用户发布的文章内容。主要接口包括：

- 发布/更新/删除文章
- 获取文章列表和详情
- 文章分类管理
- 点赞/收藏/评论功能

### 搜索服务 (Search)

搜索服务负责全站内容的搜索功能。主要接口包括：

- 全局搜索
- 分类搜索（商品、文章、失物招领、用户）
- 搜索历史管理
- 热门搜索关键词

### AI 辅助服务 (AI)

AI 辅助服务提供智能化功能增强用户体验。主要接口包括：

- 分析商品图片
- 生成商品描述
- 估算商品价格
- 商品信息优化建议

### 地图服务 (Map)

地图服务提供位置相关的功能。主要接口包括：

- 地址解析和逆地址解析
- 位置搜索和建议
- 周边检索
- 用户地址管理

### 上传服务 (Upload)

上传服务处理各种文件的上传功能。主要接口包括：

- 图片上传
- 文件上传
- 头像上传
- 通过 URL 上传图片

### API 网关 (Gateway)

API 网关是前端访问后端服务的统一入口，负责请求的路由、转发和处理。

## 使用说明

1. 使用 go-zero 工具生成代码：

```bash
# 生成 API 服务代码
goctl api go -api user/user.api -dir ./gen/api/user

# 生成 RPC 服务代码
goctl rpc protoc user/user.proto --go_out=./gen/rpc/user --go-grpc_out=./gen/rpc/user --zrpc_out=./gen/rpc/user
```

2. 实现生成的代码，编写业务逻辑

3. 启动服务

## 部署说明

本项目基于微服务架构，每个服务可以独立部署。推荐使用 Docker 和 Kubernetes 进行容器化部署。

## 开发规范

1. API 命名遵循 RESTful 规范
2. 函数名和变量名使用驼峰命名法
3. 注释遵循 Golang 标准格式
4. 错误处理遵循统一的响应格式

## 注意事项

- 所有 API 接口都需要有相应的权限控制
- 敏感操作需要进行日志记录
- 接口需要有限流保护
- 文件上传需要进行类型和大小限制