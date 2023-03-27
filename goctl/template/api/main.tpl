package main

import (
	"flag"
	"fmt"

	{{.importPackages}}
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

func main() {
	flag.Parse()

    var c config.Config
    conf.MustLoad(*configFile, &c)

    server := rest.MustNewServer(c.RestConf)
    defer server.Stop()

    logx.DisableStat()
    svcCtx := svc.NewServiceContext(c)
    handler.RegisterHandlers(server, svcCtx)

    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}