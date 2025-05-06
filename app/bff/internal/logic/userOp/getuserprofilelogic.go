package userOp

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/user/rpc/userservice"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户详细资料
func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProfileLogic) GetUserProfile(req *types.UserProfileReq) (resp *types.UserProfileResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("=== %+v", req)
	userProfile, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userservice.GetUserInfoRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var res types.UserProfileResp
	fmt.Printf("=== %+v", userProfile)
	copier.Copy(&res, userProfile)
	fmt.Printf("=== %+v", res)
	return &res, nil
}
