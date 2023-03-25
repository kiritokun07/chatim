package ws

import (
	"context"

	"chatim/service/chatim/internal/logic/ws/hub"
	"chatim/service/chatim/internal/svc"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type WsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsLogic {
	return &WsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsLogic) Ws(conn *websocket.Conn, token string) {
	println("Ws~")
	client := hub.NewClient(l.svcCtx.WsHub, conn, token)
	client.Run()
	threading.GoSafe(func() {
		l.handleRead(client)
	})
}

// 处理来自web客户端的上行消息
func (l *WsLogic) handleRead(c *hub.Client) {
	for {
		select {
		case message := <-c.Read:
			threading.GoSafe(func() {
				l.handleMsg(message)
			})
		}
	}
}

func (l *WsLogic) handleMsg(message []byte) {
	logx.Infov(string(message))
	//TODO 消息分发
	if err := l.svcCtx.Producer.PushByTopic(l.ctx, l.svcCtx.Config.ProducerInfo.MtflowerTopic, message); err != nil {
		l.Error(err)
	}
}
