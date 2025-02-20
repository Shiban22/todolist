package user

import (
	"context"
	"errors"
	"time"

	"todolist/user-api/internal/svc"
	"todolist/user-api/internal/types"
	"todolist/user-api/model"
	"todolist/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	var user *model.User
	var userid int64
	user, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != req.Password {
		return nil, errors.New("password error")
	}
	userid = user.Id
	//create token
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := utils.GenerateToken(l.svcCtx.Config.Auth.AccessSecret, userid, time.Duration(accessExpire)*time.Second)
	if err != nil {
		return nil, err
	}
	resp = &types.BaseResponse{
		Code: 0,
		Msg:  "登陆成功",
		Data: map[string]interface{}{
			"token": token,
		},
	}
	return resp, nil
}
