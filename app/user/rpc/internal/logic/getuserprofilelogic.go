package logic

import (
	"context"
	"github.com/pkg/errors"

	"github.com/liumkssq/goapp/app/user/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户详细资料
func (l *GetUserProfileLogic) GetUserProfile(in *user.GetUserProfileRequest) (*user.UserProfileResponse, error) {
	one, err := l.svcCtx.UsersModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, err
	}
	if one == nil {
		return nil, errors.New("用户不存在")
	}
	// 组装数据
	userProfile := user.UserProfileResponse{
		UserId:   int64(one.Id),
		Username: one.Username,
		Nickname: one.Nickname.String,
		Avatar:   one.AvatarUrl.String,
	}

	return &userProfile, nil
}
