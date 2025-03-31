package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"
	"github.com/liumkssq/goapp/app/user/rpc/user"
	"github.com/liumkssq/goapp/pkg/ctxdata"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 账号密码登录
func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByPasswordLogic) LoginByPassword(req *types.LoginByPasswordReq) (*types.LoginResp, error) {
	loginResponse, err := l.svcCtx.UserRpc.LoginByPassword(l.ctx, &user.LoginByPasswordRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	var resp types.LoginResp
	_ = copier.Copy(&resp, loginResponse)
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := ctxdata.GetJWTToken(
		l.svcCtx.Config.Auth.AccessSecret,
		now,
		accessExpire,
		loginResponse.UserId,
	)
	resp.AccessToken = accessToken
	resp.AccessExpire = now + accessExpire
	resp.RefreshAfter = now + accessExpire/2
	return &resp, nil
}
