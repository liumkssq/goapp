package svc

import (
	"github.com/liumkssq/goapp/app/bff/internal/config"
	"github.com/liumkssq/goapp/app/bff/internal/middleware"
	"github.com/liumkssq/goapp/app/im/rpc/imclient"
	"github.com/liumkssq/goapp/app/lostfound/rpc/lostfoundservice"
	"github.com/liumkssq/goapp/app/product/rpc/productservice"
	"github.com/liumkssq/goapp/app/search/rpc/searchservice"
	"github.com/liumkssq/goapp/app/social/rpc/socialclient"
	"github.com/liumkssq/goapp/app/user/rpc/userservice"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	UserRpc               userservice.UserService
	ProductRpc            productservice.ProductService
	LostFoundRpc          lostfoundservice.LostFoundService
	SearchRpc             searchservice.SearchService
	SocialRpc             socialclient.Social
	ImRpc                 imclient.Im
	IdempotenceMiddleware rest.Middleware
	LimitMiddleware       rest.Middleware
	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		UserRpc:               userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc:            productservice.NewProductService(zrpc.MustNewClient(c.ProductRpc)),
		LostFoundRpc:          lostfoundservice.NewLostFoundService(zrpc.MustNewClient(c.LostFoundRpc)),
		SearchRpc:             searchservice.NewSearchService(zrpc.MustNewClient(c.SearchRpc)),
		IdempotenceMiddleware: middleware.NewIdempotenceMiddleware().Handle,
		LimitMiddleware:       middleware.NewLimitMiddleware().Handle,
		SocialRpc:             socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		ImRpc:                 imclient.NewIm(zrpc.MustNewClient(c.ImRpc)),
	}
}
