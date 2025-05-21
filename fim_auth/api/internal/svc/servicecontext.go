package svc

import (
	"fim_server/core"
	"fim_server/fim_auth/api/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	gorm_DB := core.InitGorm(c.Mysql.DBSource)
	return &ServiceContext{
		Config: c,
		DB:     gorm_DB,
	}
}
