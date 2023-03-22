package config

import "github.com/zeromicro/go-zero/rest"

type (
	Config struct {
		rest.RestConf
		SendMq RocketMq
		ReadMq RocketMq
	}

	RocketMq struct {
		Addr  []string
		Topic string
		Group string `json:",optional"`
		Tag   string `json:",optional"`
	}
)
