package consumer

import (
	"context"

	"chatim/service/chatim/internal/svc"
	"chatim/shared/mq"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	ConsumerLogic struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
)

func NewConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConsumerLogic {
	return &ConsumerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumerLogic) Consumer() mq.SubFunc {
	return func(ctx context.Context, messages ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, m := range messages {
			logx.Info(string(m.Body))
			l.HandleDownMsg(m.Body)
		}
		return consumer.ConsumeSuccess, nil
	}
}

// HandleDownMsg 处理下行消息 发送到web客户端
func (l *ConsumerLogic) HandleDownMsg(message []byte) {
	l.Dispatch(message)
}

func (l *ConsumerLogic) Dispatch(message []byte) {
	for _, client := range l.svcCtx.WsHub.Clients {
		client.Send <- message
	}
}
