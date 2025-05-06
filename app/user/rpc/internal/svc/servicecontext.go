package svc

import (
	"github.com/liumkssq/goapp/app/user/model"
	"github.com/liumkssq/goapp/app/user/rpc/internal/config"
	"github.com/liumkssq/goapp/pkg/constants"
	"github.com/liumkssq/goapp/pkg/ctxdata"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UserModel
	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		Redis:      redis.MustNewRedis(c.Redisx),
		UsersModel: model.NewUserModel(sqlConn, c.Cache),
	}
}
func (svc *ServiceContext) SetRootToken() error {
	// 生成jwt
	systemToken, err := ctxdata.GetJwtToken(svc.Config.JwtAuth.AccessSecret, time.Now().Unix(), 999999999, constants.SYSTEM_ROOT_UID)
	if err != nil {
		return err
	}
	// 写入到redis

	return svc.Redis.Set(constants.REDIS_SYSTEM_ROOT_TOKEN, systemToken)
}
