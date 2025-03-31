# API与RPC集成方案

本文档描述了如何将HTTP API服务与gRPC服务进行集成，以确保前后端能够高效协同工作。

## 1. 服务通信架构

整个系统采用以下架构进行服务间通信：

```
前端应用
   │
   ▼
API网关 (Gateway)
   │
   ├───► 用户API服务 ───► 用户RPC服务 ───► 数据库/缓存
   │
   ├───► IM API服务 ───► IM RPC服务 ───► 数据库/缓存
   │
   ├───► 商品API服务 ───► 商品RPC服务 ───► 数据库/缓存
   │
   └───► ... 其他服务
```

## 2. 服务之间的调用方式

### 2.1 API服务调用RPC服务

API服务负责处理HTTP/WebSocket请求，然后通过RPC调用内部服务处理业务逻辑。

```go
// API服务中调用RPC服务示例
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
    // 调用RPC服务
    rpcResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
        Username: req.Username,
        Password: req.Password,
    })
    if err != nil {
        return nil, err
    }
    
    // 转换RPC响应为API响应
    return &types.LoginResp{
        UserId:        rpcResp.UserId,
        Username:      rpcResp.Username,
        AccessToken:   rpcResp.AccessToken,
        AccessExpire:  rpcResp.AccessExpire,
        RefreshAfter:  rpcResp.RefreshAfter,
    }, nil
}
```

### 2.2 RPC服务之间的调用

某些场景下，RPC服务可能需要调用其他RPC服务：

```go
// RPC服务中调用另一个RPC服务示例
func (s *UserServer) GetUserWithPosts(ctx context.Context, req *user.GetUserWithPostsRequest) (*user.GetUserWithPostsResponse, error) {
    // 调用用户RPC服务获取用户信息
    userInfo, err := s.GetUser(ctx, &user.GetUserRequest{
        UserId: req.UserId,
    })
    if err != nil {
        return nil, err
    }
    
    // 调用文章RPC服务获取用户的文章
    posts, err := s.articleRpc.GetUserPosts(ctx, &article.GetUserPostsRequest{
        UserId: req.UserId,
    })
    if err != nil {
        return nil, err
    }
    
    // 组合响应
    return &user.GetUserWithPostsResponse{
        User: userInfo.User,
        Posts: posts.Posts,
    }, nil
}
```

## 3. 数据模型转换

API服务和RPC服务都有自己的数据模型，需要进行转换：

### 3.1 HTTP请求 -> RPC请求

```go
// HTTP请求转换为RPC请求
func convertToRpcRequest(apiReq *types.CreateProductReq) *product.CreateProductRequest {
    return &product.CreateProductRequest{
        Title:       apiReq.Title,
        Description: apiReq.Description,
        Price:       apiReq.Price,
        CategoryId:  apiReq.CategoryId,
        Images:      apiReq.Images,
        Tags:        apiReq.Tags,
    }
}
```

### 3.2 RPC响应 -> HTTP响应

```go
// RPC响应转换为HTTP响应
func convertToApiResponse(rpcResp *product.CreateProductResponse) *types.CreateProductResp {
    return &types.CreateProductResp{
        Id:          rpcResp.Id,
        Title:       rpcResp.Title,
        Description: rpcResp.Description,
        Price:       rpcResp.Price,
        CategoryId:  rpcResp.CategoryId,
        Images:      rpcResp.Images,
        Tags:        rpcResp.Tags,
        CreatedAt:   rpcResp.CreatedAt,
    }
}
```

## 4. 错误处理

### 4.1 RPC错误转换为HTTP错误

```go
// RPC错误转换为HTTP错误
func handleRpcError(err error) error {
    if err == nil {
        return nil
    }
    
    st, ok := status.FromError(err)
    if !ok {
        return errorx.NewDefaultError("未知错误")
    }
    
    switch st.Code() {
    case codes.NotFound:
        return errorx.NewDefaultError("资源不存在")
    case codes.AlreadyExists:
        return errorx.NewDefaultError("资源已存在")
    case codes.PermissionDenied:
        return errorx.NewDefaultError("权限不足")
    default:
        return errorx.NewDefaultError(st.Message())
    }
}
```

## 5. 认证和授权

### 5.1 从API服务传递JWT到RPC服务

```go
// 在API服务中添加JWT到RPC上下文
func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfoResp, err error) {
    // 从JWT中获取用户ID
    userId := l.ctx.Value("userId").(int64)
    
    // 创建带有metadata的上下文
    md := metadata.New(map[string]string{
        "userId": strconv.FormatInt(userId, 10),
    })
    ctx := metadata.NewOutgoingContext(l.ctx, md)
    
    // 使用新上下文调用RPC
    rpcResp, err := l.svcCtx.UserRpc.GetUserInfo(ctx, &user.GetUserInfoRequest{})
    if err != nil {
        return nil, err
    }
    
    // 转换响应
    return &types.UserInfoResp{
        User: types.UserInfo{
            UserId:   rpcResp.User.UserId,
            Username: rpcResp.User.Username,
            Avatar:   rpcResp.User.Avatar,
            // ... 其他字段
        },
    }, nil
}
```

### 5.2 在RPC服务中验证身份

```go
// RPC服务中间件，验证身份
func AuthInterceptor() grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
        md, ok := metadata.FromIncomingContext(ctx)
        if !ok {
            return nil, status.Error(codes.Unauthenticated, "无元数据")
        }
        
        userIds := md.Get("userId")
        if len(userIds) == 0 {
            return nil, status.Error(codes.Unauthenticated, "缺少用户ID")
        }
        
        userId, err := strconv.ParseInt(userIds[0], 10, 64)
        if err != nil {
            return nil, status.Error(codes.Unauthenticated, "用户ID无效")
        }
        
        // 将用户ID添加到上下文
        newCtx := context.WithValue(ctx, "userId", userId)
        
        // 继续处理请求
        return handler(newCtx, req)
    }
}
```

## 6. 分布式追踪

使用jaeger进行分布式追踪，确保能够跟踪整个请求链路：

```go
// API服务中配置追踪
func main() {
    flag.Parse()
    
    var c config.Config
    conf.MustLoad(*configFile, &c)
    
    // 配置jaeger
    tracer, closer, err := trace.NewJaegerTracer("user-apis", c.Jaeger.Endpoint)
    if err != nil {
        panic(err)
    }
    defer closer.Close()
    
    ctx := svc.NewServiceContext(c)
    server := rest.MustNewServer(c.RestConf)
    defer server.Stop()
    
    handler.RegisterHandlers(server, ctx)
    
    // ... 其他配置
    
    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}
```

## 7. 前端API调用

前端应用应当通过API网关调用后端服务，以确保统一的认证和路由：

```javascript
// 前端API调用示例
async function loginUser(username, password) {
  try {
    const response = await fetch('/api/user/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username,
        password,
      }),
    });
    
    if (!response.ok) {
      throw new Error('登录失败');
    }
    
    const data = await response.json();
    
    // 保存token
    localStorage.setItem('token', data.accessToken);
    
    return data;
  } catch (error) {
    console.error('登录错误:', error);
    throw error;
  }
}
```

## 8. WebSocket集成

对于WebSocket通信，前端需要单独处理：

```javascript
// WebSocket连接示例
function connectWebSocket() {
  const token = localStorage.getItem('token');
  const ws = new WebSocket(`ws://localhost:8889/api/im/ws/connect?token=${token}`);
  
  ws.onopen = () => {
    console.log('WebSocket连接已建立');
  };
  
  ws.onmessage = (event) => {
    const message = JSON.parse(event.data);
    handleIncomingMessage(message);
  };
  
  ws.onclose = () => {
    console.log('WebSocket连接已关闭');
    // 可以在这里实现重连逻辑
    setTimeout(connectWebSocket, 3000);
  };
  
  ws.onerror = (error) => {
    console.error('WebSocket错误:', error);
  };
  
  // 返回WebSocket实例，以便在其他地方使用
  return ws;
}
```

## 9. 部署考虑事项

在将API和RPC服务部署到生产环境时，需要考虑以下几点：

1. 使用Docker容器化每个服务
2. 使用Kubernetes进行服务编排
3. 配置Etcd进行服务发现
4. 设置适当的资源限制和请求
5. 配置健康检查和就绪探针
6. 实现自动扩展策略

## 10. 总结

通过上述集成方案，我们建立了一个清晰的API与RPC服务之间的通信机制，确保了：

1. 前端应用可以通过HTTP/WebSocket与后端通信
2. API服务负责处理外部请求并转发给RPC服务
3. RPC服务负责实现核心业务逻辑
4. 所有服务都遵循统一的认证、授权和错误处理机制
5. 整个系统可以进行分布式追踪，便于排查问题