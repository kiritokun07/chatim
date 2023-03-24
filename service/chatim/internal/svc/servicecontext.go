package svc

import (
	"chatim/service/chatim/internal/config"
	"chatim/service/chatim/internal/logic/ws/hub"
)

type ServiceContext struct {
	Config config.Config
	WsHub  *hub.Hub
}

func NewServiceContext(c config.Config) *ServiceContext {
	WsHub := hub.NewHub()
	return &ServiceContext{
		Config: c,
		WsHub:  WsHub,
	}
}
