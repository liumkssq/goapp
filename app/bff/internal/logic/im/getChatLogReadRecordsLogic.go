package im

import (
	"context"
	"github.com/liumkssq/goapp/app/im/rpc/imclient"
	"github.com/liumkssq/goapp/app/social/rpc/socialclient"
	"github.com/liumkssq/goapp/app/user/rpc/user"
	"github.com/liumkssq/goapp/app/user/rpc/userservice"
	"github.com/liumkssq/goapp/pkg/bitmap"
	"github.com/liumkssq/goapp/pkg/constants"
	"strconv"

	"github.com/liumkssq/goapp/app/bff/internal/svc"
	"github.com/liumkssq/goapp/app/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatLogReadRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatLogReadRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatLogReadRecordsLogic {
	return &GetChatLogReadRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetChatLogReadRecords 获取消息已读记录 1.获取聊天记录 2.获取已读记录 3.获取未读记录
func (l *GetChatLogReadRecordsLogic) GetChatLogReadRecords(req *types.GetChatLogReadRecordsReq) (resp *types.GetChatLogReadRecordsResp, err error) {
	chatlogs, err := l.svcCtx.ImRpc.GetChatLog(l.ctx, &imclient.GetChatLogReq{
		MsgId: req.MsgId,
	})

	if err != nil || len(chatlogs.List) == 0 {
		return nil, err
	}
	var (
		chatlog = chatlogs.List[0]
		reads   = []string{chatlog.SendId}
		unreads []string
		ids     []string
	)
	switch constants.ChatType(chatlog.ChatType) {
	case constants.SingleChatType:
		if len(chatlog.ReadRecords) == 0 || chatlog.ReadRecords[0] == 0 {
			unreads = []string{chatlog.RecvId}
		} else {
			reads = append(reads, chatlog.RecvId)
		}
		ids = []string{chatlog.RecvId, chatlog.SendId}
	case constants.GroupChatType:
		groupUsers, err := l.svcCtx.SocialRpc.GroupUsers(l.ctx, &socialclient.GroupUsersReq{
			GroupId: chatlog.RecvId,
		})
		if err != nil {
			return nil, err
		}
		bitmaps := bitmap.Load(chatlog.ReadRecords)
		for _, members := range groupUsers.List {
			ids = append(ids, members.UserId)

			if members.UserId == chatlog.SendId {
				continue
			}

			if bitmaps.IsSet(members.UserId) {
				reads = append(reads, members.UserId)
			} else {
				unreads = append(unreads, members.UserId)
			}
		}
	}
	userEntities, err := l.svcCtx.UserRpc.FindUser(l.ctx, &userservice.FindUserReq{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	userEntitySet := make(map[string]*user.UserInfoResponse, len(userEntities.User))
	for _, userEntity := range userEntities.User {
		userEntitySet[strconv.FormatInt(userEntity.UserId, 10)] = userEntity
	}

	for i, read := range reads {
		if u := userEntitySet[read]; u != nil {
			reads[i] = u.Phone
		}
	}
	// 设置手机号码
	for i, unread := range unreads {
		if u := userEntitySet[unread]; u != nil {
			unreads[i] = u.Phone
		}
	}
	return &types.GetChatLogReadRecordsResp{
		Reads:   reads,
		UnReads: unreads,
	}, nil
}
