package user

import (
	"context"
	"fmt"
	"github.com/liumkssq/goapp/app/user/rpc/userservice"
	"strconv"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// finduser
func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserLogic) FindUser(req *types.FindUserReq) (resp *types.FindUserResp, err error) {
	// todo: add your logic here and delete this line
	var ids []string
	for _, v := range req.Ids {
		sId := strconv.Itoa(int(v))
		ids = append(ids, sId)
	}
	Users, err := l.svcCtx.UserRpc.FindUser(l.ctx, &userservice.FindUserReq{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	if Users == nil {
		return nil, err
	}
	fmt.Printf("Users :%+v\n", Users)
	var res types.FindUserResp
	for _, item := range Users.User {
		userResp := types.UserProfileResp{
			UserId:   item.UserId,
			Username: item.Username,
			Avatar:   item.Avatar,
		}
		res.Users = append(res.Users, userResp)
	}
	//copier.Copy(&res, Users.User)
	fmt.Printf("data :%+v\n", res)
	return &res, nil
}
