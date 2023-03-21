package config

import "github.com/zeromicro/go-zero/rest"

type (
	Config struct {
		rest.RestConf
		SendMq RocketMq
	}

	RocketMq struct {
		Addr  []string
		Topic string
		Group string
		Tag   string `json:",optional"`
	}
)
