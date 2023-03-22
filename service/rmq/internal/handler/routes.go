// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	callback "chatim/service/rmq/internal/handler/callback"
	"chatim/service/rmq/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ebflower",
				Handler: callback.EbflowerCallbackHandler(serverCtx),
			},
		},
		rest.WithPrefix("/rmq"),
	)
}