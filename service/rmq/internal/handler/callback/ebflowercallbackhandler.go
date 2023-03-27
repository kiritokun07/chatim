package callback

import (
	"net/http"

	"chatim/service/rmq/internal/logic/callback"
	"chatim/service/rmq/internal/svc"
	"chatim/service/rmq/internal/types"
	"chatim/shared/errorx"
	"chatim/shared/httpreq"
	"chatim/shared/httpresp"
)

func EbflowerCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EbDownMsgReq
		if err := httpreq.Parse(r, &req); err != nil {
			httpresp.HttpErr(w, r, errorx.NewStatCodeError(http.StatusNotAcceptable, 2, err.Error()))
			return
		}

		l := callback.NewEbflowerCallbackLogic(r.Context(), svcCtx)
		err := l.EbflowerCallback(&req)
		httpresp.Http(w, r, nil, err)
	}
}
