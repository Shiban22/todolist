package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"todolist/todo-api/internal/svc"
	"todolist/todo-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTodosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTodosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTodosLogic {
	return &ListTodosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTodosLogic) ListTodos(req *types.Todo) (resp *types.ListResponse, err error) {
	// todo: add your logic here and delete this line
	page := req.Page
	pageSize := req.PageSize
	rowbuilder := squirrel.Select()
	listResponse, err := l.svcCtx.TodoModel.FindPageListByPage(l.ctx, rowbuilder, int64(page), int64(pageSize), "id asc")
	if err != nil {
		return nil, err
	}
	var todos []types.Todo
	for _, todo := range listResponse {
		todos = append(todos, types.Todo{
			Id:    todo.Id,
			Title: todo.Title,
		})
	}
	return &types.ListResponse{
		Todos: todos,
	}, nil
	// 返回分页结果
}
