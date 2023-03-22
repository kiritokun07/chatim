package svc

import (
	"chatim/mq"
	"chatim/service/rmq/internal/config"
	"chatim/service/rmq/internal/ws"
)

type ServiceContext struct {
	Config     config.Config
	MtflowerWs *ws.MtflowerWs
}

func NewServiceContext(c config.Config) *ServiceContext {

	producer, err := mq.NewProducer(mq.ProducerConf{
		Addr:  c.SendMq.Addr,
		Group: c.SendMq.Group,
		Topic: c.SendMq.Topic,
	})
	if err != nil {
		panic(err)
	}

	wsUrl := "ws://127.0.0.1:8888/platform/ws/mtflower?token=111&platformType=1"
	//wsUrl := ws.GetMtflowerWsUrl("4833", "wo5328fLL-MKvXYeqJzQmvRqMx7zXKW-Jfh6NVvDBjb2XU-wo")
	mtflowerWs, err := ws.NewMtflowerWs(wsUrl, producer)
	if err != nil {
		panic(err)
	}

	//TODO 一个平台new一个consumer 一个consumer监听多个源的消息，每个源一个tag
	return &ServiceContext{
		Config:     c,
		MtflowerWs: mtflowerWs,
	}
}
