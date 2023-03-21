package ws

import (
	"context"

	"chatim/service/platform/internal/logic/ws/hub"
	"chatim/service/platform/internal/svc"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/threading"

	"github.com/zeromicro/go-zero/core/logx"
)

type MtflowerWsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMtflowerWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MtflowerWsLogic {
	return &MtflowerWsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MtflowerWsLogic) MtflowerWs(conn *websocket.Conn, token string) {
	client := hub.NewClient(l.svcCtx.WsHub, conn, token)
	println("register" + client.Token)
	client.Hub.Register <- client
	threading.GoSafe(func() {
		client.ReadPump()
	})
	threading.GoSafe(func() {
		client.WritePump()
	})
}
