package mq

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type (
	Producer struct {
		producer rocketmq.Producer
		//Topic    string
		Tag string
	}

	ProducerConf struct {
		Addr  []string
		Topic string
		Group string
		Tag   string
	}

	ConsumerConf struct {
		Addr  []string
		Group string
		Topic string
		Tag   string
		Fn    SubFunc
	}
)

func NewProducer(conf ProducerConf) (*Producer, error) {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(conf.Addr),
		producer.WithRetry(3),
		producer.WithGroupName(conf.Group),
	)
	if err != nil {
		return nil, err
	}
	err = p.Start()
	if err != nil {
		return nil, err
	}

	return &Producer{
		producer: p,
		//Topic:    conf.Topic,
		Tag: conf.Tag,
	}, err
}

//func (p *Producer) Push(ctx context.Context, v []byte) error {
//	msg := &primitive.Message{
//		Topic: p.Topic,
//		Body:  v,
//	}
//	msg.WithTag(p.Tag)
//	_, err := p.producer.SendSync(ctx, msg)
//	return err
//}

func (p *Producer) PushByTopic(ctx context.Context, topic string, v []byte) error {
	msg := &primitive.Message{
		Topic: topic,
		Body:  v,
	}
	msg.WithTag(p.Tag)
	_, err := p.producer.SendSync(ctx, msg)
	return err
}

type SubFunc func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)

func NewConsumer(conf ConsumerConf) {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer(conf.Addr), // 接入点地址
		consumer.WithGroupName(conf.Group), // 分组名称
	)
	_ = c.Subscribe(conf.Topic, consumer.MessageSelector{Expression: conf.Tag}, conf.Fn)
	_ = c.Start()
}
