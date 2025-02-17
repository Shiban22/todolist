package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"todolist/todo-api/internal/logic"
	"todolist/todo-api/internal/svc"
	"todolist/todo-api/internal/types"
)

func listTodosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Todo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListTodosLogic(r.Context(), svcCtx)
		resp, err := l.ListTodos(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
