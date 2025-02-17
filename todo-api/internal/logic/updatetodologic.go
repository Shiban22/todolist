package logic

import (
	"context"

	"todolist/todo-api/internal/svc"
	"todolist/todo-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTodoLogic {
	return &UpdateTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTodoLogic) UpdateTodo(req *types.UpdateRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
