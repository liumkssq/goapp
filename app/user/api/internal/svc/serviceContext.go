package svc

import (
	"github.com/liumkssq/goapp/app/user/api/internal/config"
	"github.com/liumkssq/goapp/app/user/rpc/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserService
	//CrosMiddleWare rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		//CrosMiddleWare: middleware.NewCorsMiddleware().Handle,
	}
}
