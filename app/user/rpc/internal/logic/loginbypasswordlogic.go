package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/user/model"
	"github.com/liumkssq/goapp/pkg/tool"
	"github.com/liumkssq/goapp/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/liumkssq/goapp/app/user/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 账号密码登录
func (l *LoginByPasswordLogic) LoginByPassword(in *user.LoginByPasswordRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	u, err := l.svcCtx.UsersModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "username:%s, password:%s", in.Username, in.Password)
	}
	if u == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "username:%s, password:%s", in.Username, in.Password)
	}
	//
	isVaild := tool.ValidatePasswordHash(in.Password, u.PasswordHash)
	if !isVaild {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "username:%s, password:%s", in.Username, in.Password)
	}
	var resp user.LoginResponse
	_ = copier.Copy(&resp, u)
	resp.UserId = int64(u.Id)
	return &resp, nil
}
