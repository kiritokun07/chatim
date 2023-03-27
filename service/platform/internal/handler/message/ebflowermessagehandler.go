package message

import (
	"net/http"

	"chatim/service/platform/internal/logic/message"
	"chatim/service/platform/internal/svc"
	"chatim/service/platform/internal/types"
	"chatim/shared/errorx"
	"chatim/shared/httpreq"
	"chatim/shared/httpresp"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EbflowerMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EbUpMsgReq
		if err := httpreq.Parse(r, &req); err != nil {
			httpresp.HttpErr(w, r, errorx.NewStatCodeError(http.StatusNotAcceptable, 2, err.Error()))
			return
		}

		l := message.NewEbflowerMessageLogic(r.Context(), svcCtx)
		resp, err := l.EbflowerMessage(&req)
		//这里不用包装一次
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
