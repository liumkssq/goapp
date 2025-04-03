/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package svc

import (
	"github.com/liumkssq/goapp/app/im/immodels"
	"github.com/liumkssq/goapp/app/im/ws/internal/config"
	"github.com/liumkssq/goapp/app/task/mq/mqclient"
)

type ServiceContext struct {
	Config config.Config
	immodels.ChatLogModel
	mqclient.MsgChatTransferClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		MsgChatTransferClient: mqclient.NewMsgChatTransferClient(c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
		ChatLogModel:          immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
	}
}
