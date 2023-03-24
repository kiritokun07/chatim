package config

import "github.com/zeromicro/go-zero/rest"

type (
	Config struct {
		rest.RestConf
		SendMq RocketMq
		ReadMq RocketMq

		SendTopic SendTopicInfo
	}

	RocketMq struct {
		Addr  []string
		Topic string `json:",optional"`
		Group string `json:",optional"`
		Tag   string `json:",optional"`
	}

	SendTopicInfo struct {
		MtflowerTopic string //美团鲜花下行消息topic
		EbflowerTopic string //饿百鲜花下行消息topic
	}
)
