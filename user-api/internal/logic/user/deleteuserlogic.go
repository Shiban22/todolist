package user

import (
	"context"

	"todolist/user-api/internal/svc"
	"todolist/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.UserIdReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.UserModel.Delete(l.ctx,req.Id)
	if err != nil {
		return nil, err
	}
	return &types.BaseResponse{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}
