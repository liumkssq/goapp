package svc

import (
	"github.com/liumkssq/goapp/app/im/api/internal/config"
	"github.com/liumkssq/goapp/app/im/rpc/imclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	imclient.Im
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Im:     imclient.NewIm(zrpc.MustNewClient(c.ImRpc)),
	}
}
