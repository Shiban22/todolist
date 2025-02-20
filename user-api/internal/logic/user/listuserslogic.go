package user

import (
	"context"

	"todolist/user-api/internal/svc"
	"todolist/user-api/internal/types"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUsersLogic {
	return &ListUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUsersLogic) ListUsers(req *types.UserListReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	keyword := req.Keyword
	rowbuilder := squirrel.Select()
	if keyword != "" {
		rowbuilder = rowbuilder.Where(squirrel.Or{
			squirrel.Like{"username": "%" + keyword + "%"},
			squirrel.Like{"email": "%" + keyword + "%"},
			squirrel.Like{"phone": "%" + keyword + "%"},
		})
	}
	BaseResponse, err := l.svcCtx.UserModel.FindPageListByPage(l.ctx, rowbuilder, int64(req.Page), int64(req.PageSize), "id desc")
	if err != nil {
		return nil, err
	}
	var users []types.UserInfo
	for _, user := range BaseResponse {
		users = append(users, types.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email.String,
			Phone:    user.Phone.String,
			CreateAt: user.CreateAt.String(),
			UpdateAt: user.UpdateAt.String(),
		})
	}
	resp = &types.BaseResponse{
		Code: 0,
		Msg:  "success",
		Data: types.UserListResp{
			List:        users,
			CurrentPage: req.Page,
			PageSize:    req.PageSize,
			Total:       int64(len(users)),
		},
	}
	return resp, nil
}
