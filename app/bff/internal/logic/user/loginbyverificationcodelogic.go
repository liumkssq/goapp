package user

import (
	"context"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByVerificationCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证码登录
func NewLoginByVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByVerificationCodeLogic {
	return &LoginByVerificationCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByVerificationCodeLogic) LoginByVerificationCode(req *types.LoginByVerificationCodeReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
