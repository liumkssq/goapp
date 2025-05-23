package svc

import (
	"github.com/liumkssq/goapp/app/im/immodels"
	"github.com/liumkssq/goapp/app/im/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	immodels.ChatLogModel
	immodels.ConversationsModel
	immodels.ConversationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		ChatLogModel:       immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationsModel: immodels.MustConversationsModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel:  immodels.MustConversationModel(c.Mongo.Url, c.Mongo.Db),
	}
}
