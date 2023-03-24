package ws

import (
	"context"

	"chatim/shared/mq"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func (ws *MtflowerWs) ConsumeMsg() mq.SubFunc {
	return func(ctx context.Context, messages ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		//收到消息后通过ws发给platform
		//TODO 这里需要将消息存表
		for _, m := range messages {
			err := ws.Send(m.Body)
			if err != nil {
				return consumer.ConsumeRetryLater, err
			}
		}
		return consumer.ConsumeSuccess, nil
	}
}
