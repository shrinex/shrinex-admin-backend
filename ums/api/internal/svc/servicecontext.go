package svc

import (
	"github.com/shrinex/shrinex-admin-backend/ums-api/internal/config"
	"github.com/shrinex/shrinex-admin-backend/ums-rpc/ums"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UmsClient ums.Ums
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UmsClient: ums.NewUms(zrpc.MustNewClient(c.UmsClient)),
	}
}
