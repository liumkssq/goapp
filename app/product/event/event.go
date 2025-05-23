package event

import (
	"flag"
	"github.com/liumkssq/goapp/app/product/event/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/product.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	// log、prometheus、trace、metricsUrl.
	//if err := c.SetUp(); err != nil {
	//	panic(err)
	//}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	//
	//for _, mq := range listen.Mqs(c) {
	//	serviceGroup.Add(mq)
	//}

	serviceGroup.Start()

}
