package logic

import (
	"context"

	"github.com/liumkssq/goapp/app/user/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByVerificationCodeLogic {
	return &LoginByVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证码登录
func (l *LoginByVerificationCodeLogic) LoginByVerificationCode(in *user.LoginByVerificationCodeRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &user.LoginResponse{}, nil
}
