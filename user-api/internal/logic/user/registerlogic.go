package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"todolist/user-api/internal/svc"
	"todolist/user-api/internal/types"
	"todolist/user-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("Register user exists mobile")
	}
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := new(model.User)
		insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, user)
		if err != nil {
			return err
		}
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return err
		}
		userId = lastId

		return nil // Return nil if no error occurred
	}); err != nil {
		// Handle the error
		return nil, err
	}

	return
}
