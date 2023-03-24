package ws

import (
	"context"

	"chatim/service/chatim/internal/logic/ws/hub"
	"chatim/service/chatim/internal/svc"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
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
}
