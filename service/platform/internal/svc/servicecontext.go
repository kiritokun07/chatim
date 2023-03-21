package svc

import (
	"chatim/service/platform/internal/config"
	"chatim/service/platform/internal/logic/ws/hub"
)

type ServiceContext struct {
	Config config.Config
	WsHub  *hub.Hub
}

func NewServiceContext(c config.Config) *ServiceContext {
	wsHub := hub.NewHub()
	return &ServiceContext{
		Config: c,
		WsHub:  wsHub,
	}
}
