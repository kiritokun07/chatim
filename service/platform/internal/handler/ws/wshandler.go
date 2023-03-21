package ws

import (
	"net/http"

	"chatim/service/platform/internal/logic/ws"
	"chatim/service/platform/internal/svc"
	"chatim/service/platform/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func WsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ws.NewWsLogic(r.Context(), svcCtx)
		err := l.Ws(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
