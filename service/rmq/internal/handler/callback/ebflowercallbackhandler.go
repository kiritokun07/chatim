package callback

import (
	"net/http"

	"chatim/service/rmq/internal/logic/callback"
	"chatim/service/rmq/internal/svc"
	"chatim/service/rmq/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EbflowerCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EbMsg
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := callback.NewEbflowerCallbackLogic(r.Context(), svcCtx)
		err := l.EbflowerCallback(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
