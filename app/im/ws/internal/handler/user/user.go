/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package user

import (
	"github.com/liumkssq/goapp/app/im/ws/internal/svc"
	"github.com/liumkssq/goapp/app/im/ws/websocket"
)

func OnLine(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		srv.Infof("[USER.GO DEBUG] OnLine handler triggered for client msg ID: %s", msg.Id)
		uids := srv.GetUsers()
		u := srv.GetUsers(conn)

		// 创建响应消息
		responseMsg := websocket.NewMessage(u[0], uids)
		// !!! 关键改动：将客户端原始请求的 ID 复制到响应消息中 !!!
		responseMsg.Id = msg.Id
		srv.Infof("[USER.GO DEBUG] Echoing original ID in response: %s", responseMsg.Id)

		// 发送响应
		err := srv.Send(responseMsg, conn)
		// 记录发送错误（如果需要）
		if err != nil {
			srv.Errorf("Failed to send online user list to UID %s: %v", conn.Uid, err)
		}
	}
}
