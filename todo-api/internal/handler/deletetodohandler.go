package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"todolist/todo-api/internal/logic"
	"todolist/todo-api/internal/svc"
	"todolist/todo-api/internal/types"
)

func deleteTodoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Todo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteTodoLogic(r.Context(), svcCtx)
		err := l.DeleteTodo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
