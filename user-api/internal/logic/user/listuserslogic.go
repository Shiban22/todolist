package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

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
	allowedFields := map[string]bool{
		"username": true,
		"email":    true,
		"phone":    true,
	}
	conditions := make(map[string]string)
	for _, kw := range req.Keyword {
		parts := strings.SplitN(kw, ":", 2)
		if len(parts) != 2 {
			return nil, errors.New("关键字格式错误，应为'字段:值'")
		}
		field := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		// 验证字段合法性
		if !allowedFields[field] {
			return nil, fmt.Errorf("无效查询字段: %s", field)
		}
		// 存储有效条件（覆盖重复字段）
		conditions[field] = value
	}
	//keyword := req.Keyword
	rowbuilder := squirrel.Select()
	for field, value := range conditions {
		rowbuilder = rowbuilder.Where(fmt.Sprintf("%s = ?", field), value) // 自动处理 AND 逻辑
	}
	//parts := strings.SplitN(req.Keyword, ":", 2)
	//switch keyword != "" {
	//case parts[0] == "username":
	//	rowbuilder = rowbuilder.Where(squirrel.Eq{parts[0]: parts[1]})
	//case parts[0] == "email":
	//	rowbuilder = rowbuilder.Where(squirrel.Eq{parts[0]: parts[1]})
	//case parts[0] == "phone":
	//	rowbuilder = rowbuilder.Where(squirrel.Eq{parts[0]: parts[1]})
	//	//rowbuilder = rowbuilder.Where(squirrel.Or{
	//	//	squirrel.Like{"username": "%" + keyword + "%"},
	//	//	squirrel.Like{"email": "%" + keyword + "%"},
	//	//	squirrel.Like{"phone": "%" + keyword + "%"},
	//	//})
	//}
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
