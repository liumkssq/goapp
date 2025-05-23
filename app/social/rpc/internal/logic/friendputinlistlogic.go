package logic

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/social/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/social/rpc/social"
	"github.com/liumkssq/goapp/pkg/xerr"

	"github.com/pkg/errors"

	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInListLogic {
	return &FriendPutInListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInListLogic) FriendPutInList(in *social.FriendPutInListReq) (*social.FriendPutInListResp, error) {
	// todo: add your logic here and delete this line

	userId, err := strconv.ParseUint(in.UserId, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "parse userId err %v req %v", err, in.UserId)
	}
	fmt.Printf("FriendPutInListReq userId %v\n", userId)
	// 查询好友请求列表
	friendReqList, err := l.svcCtx.FriendRequestsModel.ListNoHandler(l.ctx, userId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find list friend req err %v req %v", err, in.UserId)
	}
	fmt.Printf("FriendPutInListReq friendReqList %+v\n", friendReqList)
	var resp []*social.FriendRequests
	copier.Copy(&resp, &friendReqList)
	fmt.Printf("FriendPutInListReq resp %+v\n", resp)

	return &social.FriendPutInListResp{
		List: resp,
	}, nil
}
