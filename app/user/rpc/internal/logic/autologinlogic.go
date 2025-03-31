package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/user/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAutoLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoLoginLogic {
	return &AutoLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自动登录（仅用于开发环境测试）
func (l *AutoLoginLogic) AutoLogin(in *user.AutoLoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &user.LoginResponse{}, nil
}
