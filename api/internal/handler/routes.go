// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	mfrptest "mfrp/api/internal/handler/mfrp/test"
	"mfrp/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: mfrptest.MfrpPingHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)
}
