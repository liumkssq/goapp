package friend

import (
	"context"
	"github.com/liumkssq/goapp/app/social/rpc/socialclient"
	"github.com/liumkssq/goapp/app/user/rpc/user"
	"github.com/liumkssq/goapp/app/user/rpc/userservice"
	"github.com/liumkssq/goapp/pkg/ctxdata"
	"strconv"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友列表
func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)

	friends, err := l.svcCtx.SocialRpc.FriendList(l.ctx, &socialclient.FriendListReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	if len(friends.List) == 0 {
		return &types.FriendListResp{}, nil
	}

	// 根据好友id获取好友信息
	uids := make([]string, 0, len(friends.List))
	for _, i := range friends.List {
		uids = append(uids, i.FriendUid)
	}

	// 根据uids查询用户信息
	users, err := l.svcCtx.UserRpc.FindUser(l.ctx, &userservice.FindUserReq{
		Ids: uids,
	})
	if err != nil {
		return &types.FriendListResp{}, nil
	}
	userRecords := make(map[string]*user.UserInfoResponse, len(users.User))
	for i, _ := range users.User {
		userRecords[strconv.FormatInt(users.User[i].UserId, 10)] = users.User[i]
	}
	respList := make([]*types.Friends, 0, len(friends.List))
	for _, v := range friends.List {
		friend := &types.Friends{
			Id:        v.Id,
			FriendUid: v.FriendUid,
		}

		if u, ok := userRecords[v.FriendUid]; ok {
			friend.Nickname = u.Nickname
			friend.Avatar = u.Avatar
		}
		respList = append(respList, friend)
	}

	return &types.FriendListResp{
		List: respList,
	}, nil
}
