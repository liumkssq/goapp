/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package svc

import (
	"fmt"
	"net/http"

	"github.com/liumkssq/goapp/app/im/immodels"
	"github.com/liumkssq/goapp/app/im/ws/websocket"
	"github.com/liumkssq/goapp/app/social/rpc/socialclient"
	"github.com/liumkssq/goapp/app/task/mq/internal/config"
	"github.com/liumkssq/goapp/pkg/constants"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	config.Config

	WsClient websocket.Client
	*redis.Redis

	socialclient.Social
	immodels.ChatLogModel
	immodels.ConversationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:            c,
		Redis:             redis.MustNewRedis(c.Redisx),
		ChatLogModel:      immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel: immodels.MustConversationModel(c.Mongo.Url, c.Mongo.Db),

		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
	}

	token, err := svc.GetSystemToken()
	if err != nil {
		logx.Errorf("Failed to get system token: %v", err)
		panic(err)
	}

	if token == "" {
		logx.Error("System token is empty, this will cause authentication failures")
		panic("empty token")
	}

	// Create headers with token in both Authorization and sec-websocket-protocol
	header := http.Header{
		"Authorization":          []string{token},
		"sec-websocket-protocol": []string{token},
	}

	fmt.Printf("Initializing WebSocket client with token: %s\n", token)

	discover := websocket.NewRedisDiscover(header, constants.REDIS_DISCOVER_SRV, c.Redisx)
	svc.WsClient = websocket.NewClient(c.Ws.Host,
		websocket.WithClientHeader(header),
		websocket.WithClientDiscover(discover),
	)

	return svc
}

func (svc *ServiceContext) GetSystemToken() (string, error) {
	token, err := svc.Redis.Get(constants.REDIS_SYSTEM_ROOT_TOKEN)
	if err != nil || token == "" {
		// If we can't get the token or it's empty, log this important error
		fmt.Printf("Warning: Failed to get system token: %v, token: %s\n", err, token)
	}
	return token, err
}
