package logic

import (
	"context"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/user/api/internal/svc"
	"github.com/liumkssq/goapp/app/user/api/internal/types"
	"github.com/liumkssq/goapp/app/user/rpc/userservice"
	"github.com/liumkssq/goapp/pkg/ctxdata"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &userservice.RegisterRequest{
		Phone:            req.Phone,
		Username:         req.Username,
		Password:         req.Password,
		VerificationCode: req.VerificationCode,
		Campus:           req.Campus,
		College:          req.College,
		Major:            req.Major,
		EnrollmentYear:   int32(req.EnrollmentYear),
		UserRole:         req.UserRole,
		StudentId:        req.StudentId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	var resp types.RegisterResp

	_ = copier.Copy(&resp, registerResp)
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := ctxdata.GetJwtToken(
		l.svcCtx.Config.Auth.AccessSecret,
		now,
		accessExpire,
		strconv.FormatInt(registerResp.UserId, 10),
	)
	resp.AccessToken = accessToken
	resp.AccessExpire = now + accessExpire
	resp.RefreshAfter = now + accessExpire/2
	return &resp, nil
}
