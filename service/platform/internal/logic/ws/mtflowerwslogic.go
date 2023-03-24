package ws

import (
	"context"

	"chatim/service/platform/internal/logic/ws/hub"
	"chatim/service/platform/internal/svc"

	"github.com/gorilla/websocket"
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
	println("register ready")
	client := hub.NewClient(l.svcCtx.WsHub, conn, token)
	println("register" + client.Token)
	client.Run()
}
