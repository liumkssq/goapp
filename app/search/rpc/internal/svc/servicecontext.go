package svc

import (
	"github.com/liumkssq/goapp/app/product/model"
	"github.com/liumkssq/goapp/app/search/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
