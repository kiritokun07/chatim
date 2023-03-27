package message

import (
	"net/http"

	"chatim/service/platform/internal/logic/message"
	"chatim/service/platform/internal/svc"
	"chatim/service/platform/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EbflowerMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EbMsgReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := message.NewEbflowerMessageLogic(r.Context(), svcCtx)
		resp, err := l.EbflowerMessage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
