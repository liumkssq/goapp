package logic

import (
	"context"
	"database/sql"

	"github.com/jinzhu/copier"
	"github.com/liumkssq/goapp/app/social/rpc/internal/svc"
	"github.com/liumkssq/goapp/app/social/rpc/social"
	"github.com/liumkssq/goapp/app/social/socialmodels"
	"github.com/liumkssq/goapp/pkg/xerr"

	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	// 解析用户ID
	userId, err := strconv.ParseUint(in.UserId, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewMsg("parse uint error"), "parse uint error: %v", err)
	}

	reqUserId, err := strconv.ParseUint(in.ReqUid, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewMsg("参数错误"), "parse req id error: %v", err)
	}

	// 查询是否已经是好友
	_, err = l.svcCtx.FriendsModel.FindByUidAndFid(l.ctx, uint64ToString(userId), uint64ToString(reqUserId))
	if err == nil {
		return nil, errors.Wrapf(xerr.NewMsg("已经是好友"), "already is friend: %v", in)
	}

	// 查询是否已经发送过请求
	req := &socialmodels.FriendRequests{}
	copier.Copy(req, &in)

	_, err = l.svcCtx.FriendRequestsModel.FindByReqUidAndUserId(l.ctx, reqUserId, userId)
	if err == nil {
		return nil, errors.Wrapf(xerr.NewMsg("已经发送过请求"), "already send request: %v", in)
	}

	// 插入好友请求
	fr := &socialmodels.FriendRequests{
		UserId:       userId,
		ReqUserId:    reqUserId,
		ReqMsg:       sql.NullString{Valid: true, String: in.ReqMsg},
		ReqTime:      time.Now(),
		HandleResult: sql.NullInt64{Valid: false, Int64: 0},
		HandleMsg:    sql.NullString{Valid: false, String: ""},
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}

	_, err = l.svcCtx.FriendRequestsModel.Insert(l.ctx, fr)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert friend request db fail, err:%v, req: %v", err, fr)
	}

	return &social.FriendPutInResp{}, nil
}

// 辅助函数：将uint64转换为字符串
func uint64ToString(v uint64) string {
	return strconv.FormatUint(v, 10)
}
