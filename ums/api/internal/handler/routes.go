// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	session "github.com/shrinex/shrinex-admin-backend/ums-api/internal/handler/session"
	"github.com/shrinex/shrinex-admin-backend/ums-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/signUp",
				Handler: session.SignUpHandler(serverCtx),
			},
		},
		rest.WithPrefix("/ums/session"),
	)
}
