package logic

import (
	"context"
	"time"

	"todolist/todo-api/internal/svc"
	"todolist/todo-api/internal/types"
	"todolist/todo-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTodoLogic {
	return &CreateTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTodoLogic) CreateTodo(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line

	result, err := l.svcCtx.TodoModel.Insert(l.ctx, &model.Todo{
		Title:     req.Title,
		Completed: 0, // 默认未完成
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &types.CreateResponse{Id: id}, nil
}
