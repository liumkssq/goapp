package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/liumkssq/goapp/app/im/ws/internal/config"
	"github.com/liumkssq/goapp/app/im/ws/internal/handler"
	"github.com/liumkssq/goapp/app/im/ws/internal/svc"
	"github.com/liumkssq/goapp/app/im/ws/websocket"
	"github.com/liumkssq/goapp/pkg/constants"
	"github.com/liumkssq/goapp/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/im.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c)
	token, err := ctxdata.GetJwtToken(c.JwtAuth.AccessSecret, time.Now().Unix(), 3153600000, fmt.Sprintf("ws:%s", time.Now().Unix()))
	if err != nil {
		panic(err)
	}

	// Create headers with token in both Authorization and sec-websocket-protocol
	authHeaders := http.Header{
		"Authorization":          []string{token},
		"sec-websocket-protocol": []string{token},
	}

	opts := []websocket.ServerOptions{
		websocket.WithServerAuthentication(handler.NewJwtAuth(ctx)),
		websocket.WithServerAck(websocket.NoAck),
		websocket.WithServerDiscover(websocket.NewRedisDiscover(authHeaders, constants.REDIS_DISCOVER_SRV, c.Redisx)),
	}
	//srv := websocket.NewServer(c.ListenOn,
	//	websocket.WithServerAuthentication(handler.NewJwtAuth(ctx)),
	//	//websocket.WithServerAck(websocket.RigorAck),
	//	websocket.WithServerAck(websocket.NoAck),
	//	websocket.WithServerMaxConnectionIdle(10*60*time.Second),
	//	//websocket.WithServerDiscover(&websocket.nopDiscover{}), //
	//	websocket.WithServerCors("*"))
	//

	srv := websocket.NewServer(c.ListenOn, opts...)
	defer srv.Stop()
	handler.RegisterHandlers(srv, ctx)
	srv.Start()
}
