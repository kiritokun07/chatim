package svc

import (
	"chatim/service/rmq/internal/config"
	"chatim/service/rmq/internal/ws"
	"chatim/shared/mq"
)

type ServiceContext struct {
	Config     config.Config
	Producer   *mq.Producer
	MtflowerWs *ws.MtflowerWs
}

func NewServiceContext(c config.Config) *ServiceContext {

	//总共只有一个producer
	producer, err := mq.NewProducer(mq.ProducerConf{
		Addr:  c.RocketMq.Addr,
		Group: c.ProducerInfo.ProducerGroup,
	})
	if err != nil {
		panic(err)
	}

	wsUrl := "ws://127.0.0.1:8888/platform/ws/mtflower?token=111&platformType=1"
	//wsUrl := ws.GetMtflowerWsUrl("4833", "wo5328fLL-MKvXYeqJzQmvRqMx7zXKW-Jfh6NVvDBjb2XU-wo")
	mtflowerWs, err := ws.NewMtflowerWs(wsUrl, c.ProducerInfo.MtflowerTopic, producer)
	if err != nil {
		panic(err)
	}

	//TODO 一个平台new一个consumer 一个consumer监听多个源的消息，每个源一个tag
	return &ServiceContext{
		Config:     c,
		Producer:   producer,
		MtflowerWs: mtflowerWs,
	}
}
