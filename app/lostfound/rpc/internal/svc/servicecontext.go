package svc

import (
	"github.com/liumkssq/goapp/app/lostfound/model"
	"github.com/liumkssq/goapp/app/lostfound/rpc/internal/config"
	usermodel "github.com/liumkssq/goapp/app/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	LostFoundItemModel model.LostFoundItemModel
	UserModel          usermodel.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConnn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:             c,
		LostFoundItemModel: model.NewLostFoundItemModel(sqlConnn, c.Cache),
		UserModel:          usermodel.NewUserModel(sqlConnn, c.Cache),
	}
}
