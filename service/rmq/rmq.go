package main

import (
	"flag"
	"fmt"

	"chatim/service/rmq/internal/config"
	"chatim/service/rmq/internal/handler"
	"chatim/service/rmq/internal/svc"
	"chatim/shared/mq"

	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/rmq-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	logx.DisableStat()
	rlog.SetLogLevel("error")
	svcCtx := svc.NewServiceContext(c)
	//美团鲜花上行消息消费者
	mq.NewConsumer(mq.ConsumerConf{
		Addr:  c.RocketMq.Addr,
		Group: c.ConsumerInfo.MtflowerGroup,
		Topic: c.ConsumerInfo.MtflowerTopic,
		Fn:    svcCtx.MtflowerWs.ConsumeMsg(),
	})
	handler.RegisterHandlers(server, svcCtx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
