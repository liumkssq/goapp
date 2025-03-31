package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/user/api/internal/svc"
	"github.com/liumkssq/goapp/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 自动登录（仅用于开发环境测试）
func NewAutoLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoLoginLogic {
	return &AutoLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AutoLoginLogic) AutoLogin() (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
