package svc

import (
	Lmodel "github.com/liumkssq/goapp/app/lostfound/model"
	"github.com/liumkssq/goapp/app/product/model"
	"github.com/liumkssq/goapp/app/search/rpc/internal/config"
	Umodel "github.com/liumkssq/goapp/app/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	ProductModel   model.ProductModel
	LostFoundModel Lmodel.LostFoundItemModel
	UserModel      Umodel.UserModel
	// 其他模型
	//UserRpc userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		ProductModel:   model.NewProductModel(sqlConn, c.Cache),
		LostFoundModel: Lmodel.NewLostFoundItemModel(sqlConn, c.Cache),
		UserModel:      Umodel.NewUserModel(sqlConn, c.Cache),
		//UserRpc:        userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
	}
}
