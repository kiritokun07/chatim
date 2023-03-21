package main

import (
	"flag"
	"fmt"

	"chatim/service/platform/internal/config"
	"chatim/service/platform/internal/handler"
	"chatim/service/platform/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/threading"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/platform-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	threading.GoSafe(func() {
		ctx.WsHub.Run()
	})
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
