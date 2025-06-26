package svc

import (
	"fim_server/core"
	"fim_server/fim_auth/api/internal/config"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/fim_user/user_rpc/user"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	gorm_DB := core.InitGorm(c.Mysql.DBSource)
	redis_client := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	return &ServiceContext{
		Config:  c,
		DB:      gorm_DB,
		Redis:   redis_client,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
