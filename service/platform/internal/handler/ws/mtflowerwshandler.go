package ws

import (
	"net/http"

	"chatim/service/platform/internal/logic/ws"
	"chatim/service/platform/internal/svc"
	"chatim/service/platform/internal/types"
	"chatim/shared/errorx"
	"chatim/shared/httpresp"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func MtflowerWsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpErr(w, r, errorx.NewStatCodeError(http.StatusNotAcceptable, 2, err.Error()))
			return
		}
		bytes, _ := jsonx.Marshal(req)
		s := string(bytes)
		println("s=" + s)
		l := ws.NewMtflowerWsLogic(r.Context(), svcCtx)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			l.Error(err)
			return
		}
		l.MtflowerWs(conn, req.Token)
	}
}
