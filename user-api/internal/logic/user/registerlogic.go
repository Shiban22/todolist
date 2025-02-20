package user

import (
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"todolist/user-api/internal/svc"
	"todolist/user-api/internal/types"
	"todolist/user-api/model"
	"todolist/utils"

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
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	// err != sqlx.ErrNotFound findone查询的时候 如果结果为空会 反馈dberror，这里要加多一个判断err不为db error
	if err != nil && err != sqlx.ErrNotFound {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("Register user exists mobile")
	}
	var lastuserId int64
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := new(model.User)
		user.Username = req.Username
		user.Password = req.Password
		user.CreateAt = time.Now()
		user.UpdateAt = time.Now()
		insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, user)
		if err != nil {
			return err
		}
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return err
		}
		lastuserId = lastId
		user.Id = lastuserId

		return nil // Return nil if no error occurred
	}); err != nil {
		// Handle the error
		return nil, err
	}
	//create token
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := utils.GenerateToken(l.svcCtx.Config.Auth.AccessSecret, lastuserId, time.Duration(accessExpire)*time.Second)

	if err != nil {
		return nil, err
	}
	resp = &types.BaseResponse{
		Code: 0,
		Msg:  "注册成功",
		Data: map[string]interface{}{
			"token": token,
		},
	}
	return resp, nil
}
