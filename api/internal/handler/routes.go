// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"shorturl/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/short",
				Handler: ShortHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/expand",
				Handler: ExpandHandler(serverCtx),
			},
		},
	)
}
