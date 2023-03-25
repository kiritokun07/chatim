package config

import (
	"chatim/shared/mq"

	"github.com/zeromicro/go-zero/rest"
)

type (
	Config struct {
		rest.RestConf

		RocketMq mq.RocketMqConf

		ProducerInfo ProducerInfo

		ConsumerInfo ConsumerInfo
	}

	ProducerInfo struct {
		ProducerGroup string //下行消息group
		MtflowerTopic string //美团鲜花下行消息topic
		EbflowerTopic string //饿百鲜花下行消息topic
	}

	ConsumerInfo struct {
		MtflowerGroup string //美团鲜花上行消费者group
		MtflowerTopic string //美团鲜花上行消息topic
	}
)
