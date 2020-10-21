// DO NOT EDIT, generated by goctl
package handler

import (
	"net/http"

	"book_store/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/add",
			Handler: addHandler(serverCtx),
		},
		{
			Method:  http.MethodGet,
			Path:    "/check",
			Handler: checkHandler(serverCtx),
		},
	})
}