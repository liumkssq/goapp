/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package handler

import (
	"github.com/liumkssq/goapp/app/im/ws/internal/handler/conversation"
	"github.com/liumkssq/goapp/app/im/ws/internal/handler/push"
	"github.com/liumkssq/goapp/app/im/ws/internal/handler/user"
	"github.com/liumkssq/goapp/app/im/ws/internal/svc"
	"github.com/liumkssq/goapp/app/im/ws/websocket"
)

func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	srv.AddRoutes([]websocket.Route{
		{
			Method:  "user.online",
			Handler: user.OnLine(svc),
		},
		{
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
		{
			Method:  "conversation.markChat",
			Handler: conversation.MarkRead(svc),
		},
		{
			Method:  "push",
			Handler: push.Push(svc),
		},
	})
}
