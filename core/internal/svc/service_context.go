package svc

import (
	"cloud/core/internal/config"
	"cloud/core/internal/middleware"
	"cloud/core/models"

	"github.com/zeromicro/go-zero/rest"

	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.MySQL.DataSource),
		RDB:    models.InitRedis(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
