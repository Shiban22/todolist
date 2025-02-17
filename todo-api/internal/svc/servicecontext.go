package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"todolist/todo-api/internal/config"
	"todolist/todo-api/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	TodoModel   model.TodoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		TodoModel: model.NewTodoModel(conn, c.Cache),
	}
}
