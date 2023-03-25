package svc

import (
	"chatim/service/chatim/internal/config"
	"chatim/service/chatim/internal/logic/ws/hub"
	"chatim/shared/mq"
)

type ServiceContext struct {
	Config   config.Config
	WsHub    *hub.Hub
	Producer *mq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	//需要初始化mq生产者，每个平台的mq消费者，初始化hub
	//总共只有一个producer
	producer, err := mq.NewProducer(mq.ProducerConf{
		Addr:  c.RocketMq.Addr,
		Group: c.ProducerInfo.ProducerGroup,
	})
	if err != nil {
		panic(err)
	}

	WsHub := hub.NewHub()

	return &ServiceContext{
		Config:   c,
		WsHub:    WsHub,
		Producer: producer,
	}
}
