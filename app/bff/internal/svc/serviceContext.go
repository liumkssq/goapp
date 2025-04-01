package svc

import (
	"github.com/liumkssq/goapp/app/bff/internal/config"
	"github.com/liumkssq/goapp/app/product/rpc/productservice"
	"github.com/liumkssq/goapp/app/user/rpc/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    userservice.UserService
	ProductRpc productservice.ProductService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productservice.NewProductService(zrpc.MustNewClient(c.ProductRpc)),
	}
}
