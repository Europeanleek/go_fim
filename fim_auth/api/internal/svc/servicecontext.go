package svc

import (
	"fim_server/core"
	"fim_server/fim_auth/api/internal/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	gorm_DB := core.InitGorm(c.Mysql.DBSource)
	redis_client := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	return &ServiceContext{
		Config: c,
		DB:     gorm_DB,
		Redis:  redis_client,
	}
}
