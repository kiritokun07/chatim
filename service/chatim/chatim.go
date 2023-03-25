package main

import (
	"context"
	"flag"
	"fmt"

	"chatim/service/chatim/internal/config"
	"chatim/service/chatim/internal/handler"
	"chatim/service/chatim/internal/logic/ws/consumer"
	"chatim/service/chatim/internal/svc"
	"chatim/shared/mq"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chatim-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	logx.DisableStat()
	ctx := svc.NewServiceContext(c)
	threading.GoSafe(func() {
		ctx.WsHub.Run()
	})

	consumerLogic := consumer.NewConsumerLogic(context.TODO(), ctx)
	//美团鲜花下行消息消费者
	mq.NewConsumer(mq.ConsumerConf{
		Addr:  c.RocketMq.Addr,
		Group: c.ConsumerInfo.MtflowerGroup,
		Topic: c.ConsumerInfo.MtflowerTopic,
		//Tag:   "",
		Fn: consumerLogic.Consumer(),
	})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
