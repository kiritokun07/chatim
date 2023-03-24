package ws

import (
	"net/http"

	"chatim/service/chatim/internal/logic/ws"
	"chatim/service/chatim/internal/svc"
	"chatim/service/chatim/internal/types"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ws.NewWsLogic(r.Context(), svcCtx)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			l.Error(err)
			return
		}
		l.Ws(conn, req.Token)
	}
}
