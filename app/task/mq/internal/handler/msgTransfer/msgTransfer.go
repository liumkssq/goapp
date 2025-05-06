/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package msgTransfer

import (
	"context"
	"fmt"
	"github.com/liumkssq/goapp/app/im/ws/websocket"
	"github.com/liumkssq/goapp/app/im/ws/ws"
	"github.com/liumkssq/goapp/app/social/rpc/socialclient"
	"github.com/liumkssq/goapp/app/task/mq/internal/svc"
	"github.com/liumkssq/goapp/pkg/constants"

	"github.com/zeromicro/go-zero/core/logx"
)

type baseMsgTransfer struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseMsgTransfer(svc *svc.ServiceContext) *baseMsgTransfer {
	return &baseMsgTransfer{
		svcCtx: svc,
		Logger: logx.WithContext(context.Background()),
	}
}

func (m *baseMsgTransfer) Transfer(ctx context.Context, data *ws.Push) error {
	fmt.Println("data : ", data)
	var err error
	switch data.ChatType {
	case constants.GroupChatType:
		fmt.Println("group data : ", data)
		err = m.group(ctx, data)
	case constants.SingleChatType:
		fmt.Println("single data : ", data)
		err = m.single(ctx, data)
	}
	return err
}

func (m *baseMsgTransfer) single(ctx context.Context, data *ws.Push) error {
	fmt.Println("single data : ", data)
	return m.svcCtx.WsClient.Send(websocket.Message{
		FrameType: websocket.FrameData,
		Method:    "push",
		FormId:    constants.SYSTEM_ROOT_UID,
		Data:      data,
	})
}

func (m *baseMsgTransfer) group(ctx context.Context, data *ws.Push) error {
	// 就要查询，群的用户
	users, err := m.svcCtx.Social.GroupUsers(ctx, &socialclient.GroupUsersReq{
		GroupId: data.RecvId,
	})
	if err != nil {
		return err
	}
	data.RecvIds = make([]string, 0, len(users.List))

	for _, members := range users.List {
		if members.UserId == data.SendId {
			continue
		}

		data.RecvIds = append(data.RecvIds, members.UserId)
	}

	return m.svcCtx.WsClient.Send(websocket.Message{
		FrameType: websocket.FrameData,
		Method:    "push",
		FormId:    constants.SYSTEM_ROOT_UID,
		Data:      data,
	})
}
