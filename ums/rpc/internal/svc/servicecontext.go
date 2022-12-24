package svc

import (
	"github.com/shrinex/shrinex-admin-backend/ums-rpc/dao"
	"github.com/shrinex/shrinex-admin-backend/ums-rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	AdminUserDao dao.UmsAdminUserDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AdminUserDao: dao.NewUmsAdminUserDao(
			sqlx.NewMysql(c.Database.Datasource),
			c.Cache,
		),
	}
}
