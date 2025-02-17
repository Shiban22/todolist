// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"todolist/todo-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/todos",
				Handler: createTodoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/todos",
				Handler: updateTodoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/todos",
				Handler: listTodosHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/todos/:id",
				Handler: deleteTodoHandler(serverCtx),
			},
		},
	)
}
