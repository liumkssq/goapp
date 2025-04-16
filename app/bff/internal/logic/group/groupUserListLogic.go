package group

import (
	"context"
	"github.com/liumkssq/goapp/app/social/rpc/socialclient"
	"github.com/liumkssq/goapp/app/user/rpc/user"
	"github.com/liumkssq/goapp/app/user/rpc/userservice"
	"strconv"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 成员列表列表
func NewGroupUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUserListLogic {
	return &GroupUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUserListLogic) GroupUserList(req *types.GroupUserListReq) (resp *types.GroupUserListResp, err error) {
	// todo: add your logic here and delete this line
	groupUsers, err := l.svcCtx.SocialRpc.GroupUsers(l.ctx, &socialclient.GroupUsersReq{
		GroupId: req.GroupId,
	})

	// 还需要获取用户的信息
	uids := make([]string, 0, len(groupUsers.List))
	for _, v := range groupUsers.List {
		uids = append(uids, v.UserId)
	}

	// 获取用户信息
	userList, err := l.svcCtx.UserRpc.FindUser(l.ctx, &userservice.FindUserReq{Ids: uids})
	if err != nil {
		return nil, err
	}

	userRecords := make(map[string]*user.UserInfoResponse, len(userList.User))
	for i, _ := range userList.User {
		userRecords[strconv.FormatInt(userList.User[i].UserId, 10)] = userList.User[i]
	}

	respList := make([]*types.GroupMembers, 0, len(groupUsers.List))
	for _, v := range groupUsers.List {

		member := &types.GroupMembers{
			Id:        int64(v.Id),
			GroupId:   v.GroupId,
			UserId:    v.UserId,
			RoleLevel: int(v.RoleLevel),
		}
		if u, ok := userRecords[v.UserId]; ok {
			member.Nickname = u.Nickname
			member.UserAvatarUrl = u.Avatar
		}
		respList = append(respList, member)
	}

	return &types.GroupUserListResp{List: respList}, err
}
