// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	ws "chatim/service/chatim/internal/handler/ws"
	"chatim/service/chatim/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: ws.WsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/chatim/ws"),
	)
}
