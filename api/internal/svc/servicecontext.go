package svc

import (
	"mfrp/config"
	"mfrp/fmgr"

	"github.com/fatedier/frp/server"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config

	FrpService *server.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	svr, err := fmgr.CreateFrpsService(c)
	if err != nil {
		logx.Errorf(err.Error())
		return nil
	}

	return &ServiceContext{
		Config: c,

		FrpService: svr,
	}
}
