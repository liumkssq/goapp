package main

import (
	"flag"
	"fmt"
	"github.com/liumkssq/goapp/app/task/mq/internal/config"
	"github.com/liumkssq/goapp/app/task/mq/internal/handler"
	"github.com/liumkssq/goapp/app/task/mq/internal/svc"
	"github.com/liumkssq/goapp/pkg/constants"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var configFile = flag.String("f", "etc/dev/task.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	
	// Verify system token is available before proceeding
	rds := redis.MustNewRedis(c.Redisx)
	token, err := rds.Get(constants.REDIS_SYSTEM_ROOT_TOKEN)
	if err != nil {
		logx.Errorf("Error retrieving system token: %v", err)
	}
	
	if token == "" {
		logx.Error("WARNING: System root token is empty. WebSocket message forwarding may fail!")
	} else {
		logx.Infof("System root token is available for WebSocket connections")
	}
	
	ctx := svc.NewServiceContext(c)
	listen := handler.NewListen(ctx)

	serviceGroup := service.NewServiceGroup()
	for _, s := range listen.Services() {
		serviceGroup.Add(s)
	}
	fmt.Println("Starting mqueue at ...")
	serviceGroup.Start()
}
