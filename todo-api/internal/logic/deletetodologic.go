package logic

import (
	"context"

	"todolist/todo-api/internal/svc"
	"todolist/todo-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTodoLogic {
	return &DeleteTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTodoLogic) DeleteTodo(req *types.Todo) error {
	// todo: add your logic here and delete this line

	return nil
}
