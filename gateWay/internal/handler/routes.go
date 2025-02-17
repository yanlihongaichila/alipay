// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"gateWay/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/order/notif",
				Handler: OrderNotifHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/order/create",
				Handler: OrderCreateHandler(serverCtx),
			},
		},
	)
}
