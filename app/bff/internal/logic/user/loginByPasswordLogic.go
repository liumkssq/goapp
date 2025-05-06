package user

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"
	"github.com/liumkssq/goapp/app/user/rpc/user"
	"github.com/liumkssq/goapp/pkg/ctxdata"
	"github.com/pkg/errors"
	"strconv"
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
	fmt.Printf("req %+v\n ", req)
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
	fmt.Println("userID", loginResponse.UserId)
	Ids := strconv.FormatInt(loginResponse.UserId, 10)
	fmt.Println("Ids", Ids)
	fmt.Printf("%+v\n", Ids)
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := ctxdata.GetJwtToken(
		l.svcCtx.Config.JwtAuth.AccessSecret,
		now,
		accessExpire,
		Ids,
	)
	resp.AccessToken = accessToken
	resp.AccessExpire = now + accessExpire
	resp.RefreshAfter = now + accessExpire/2
	return &resp, nil
}
