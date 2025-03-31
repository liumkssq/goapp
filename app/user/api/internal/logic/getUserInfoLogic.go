package logic

import (
	"context"
	"fmt"
	"github.com/liumkssq/goapp/pkg/ctxdata"

	"github.com/liumkssq/goapp/app/user/rpc/userservice"

	"github.com/liumkssq/goapp/app/user/api/internal/svc"
	"github.com/liumkssq/goapp/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户信息
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (*types.UserInfoResp, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	userInfo, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userservice.GetUserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("=== %+v", userInfo)
	fmt.Println("abb", userInfo.Avatar)
	return &types.UserInfoResp{
		UserId:   userInfo.UserId,
		Username: userInfo.Username,
		Avatar:   userInfo.Avatar,
		Phone:    userInfo.Phone,
		Gender:   userInfo.Gender,
		Bio:      userInfo.Bio,
	}, nil
}
