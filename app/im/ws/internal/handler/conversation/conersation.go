/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package conversation

import (
	"github.com/liumkssq/goapp/app/im/ws/internal/svc"
	"github.com/liumkssq/goapp/app/im/ws/websocket"
	"github.com/liumkssq/goapp/app/im/ws/ws"
	"github.com/liumkssq/goapp/app/task/mq/mq"
	"github.com/liumkssq/goapp/pkg/constants"
	"github.com/liumkssq/goapp/pkg/wuid"
	"github.com/mitchellh/mapstructure"

	"time"
)

func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		// todo: 私聊
		var data ws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

		if data.ConversationId == "" {
			switch data.ChatType {
			case constants.SingleChatType:
				data.ConversationId = wuid.CombineId(conn.Uid, data.RecvId)
			case constants.GroupChatType:
				data.ConversationId = data.RecvId
			}
		}

		err := svc.MsgChatTransferClient.Push(&mq.MsgChatTransfer{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			SendTime:       time.Now().UnixMilli(),
			MType:          data.Msg.MType,
			Content:        data.Msg.Content,
			MsgId:          msg.Id,
		})
		if err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}
	}
}

func MarkRead(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		// todo: 已读未读处理
		var data ws.MarkRead
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

		err := svc.MsgReadTransferClient.Push(&mq.MsgMarkRead{
			ChatType:       data.ChatType,
			ConversationId: data.ConversationId,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			MsgIds:         data.MsgIds,
		})

		if err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}
	}
}
