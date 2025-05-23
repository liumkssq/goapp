package svc

import (
	"github.com/liumkssq/goapp/app/product/model"
	"github.com/liumkssq/goapp/app/product/rpc/internal/config"
	model2 "github.com/liumkssq/goapp/app/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
	UserModel    model2.UserModel
	ReportModel  model.ReportModel
	// 这里可以添加其他的模型
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(sqlConn, c.Cache),
		UserModel:    model2.NewUserModel(sqlConn, c.Cache),
		ReportModel:  model.NewReportModel(sqlConn, c.Cache),
		CommentModel: model.NewCommentModel(sqlConn, c.Cache),
	}
}
