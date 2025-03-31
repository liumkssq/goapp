package logic

import (
	"context"
	"fmt"
	"github.com/liumkssq/goapp/app/user/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/user/rpc/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取当前用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.UserInfoResponse, error) {
	fmt.Printf("uid: %d", in.Id)
	userInfo, err := l.svcCtx.UsersModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, err
	}
	var avatar string
	if userInfo.AvatarUrl.Valid {
		avatar = userInfo.AvatarUrl.String
	} else {
		avatar = "https://pic2.zhimg.com/v2-3bfd3fdf24511924291b909580081065_r.jpg"
	}
	var bio string
	if userInfo.Bio.Valid {
		bio = userInfo.Bio.String
	} else {
		bio = `用户什么都没有写`
	}
	return &user.UserInfoResponse{
		UserId:   int64(userInfo.Id),
		Username: userInfo.Username,
		Avatar:   avatar,
		Phone:    userInfo.Phone,
		Gender:   strconv.FormatInt(userInfo.Sex, 10),
		Bio:      bio,
	}, nil
}
