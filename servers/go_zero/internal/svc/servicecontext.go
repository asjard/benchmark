package svc

import (
	"github.com/asjard/benchmark/servers"
	"github.com/asjard/benchmark/servers/go_zero/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	Options *servers.ServerOptions
}

func NewServiceContext(c config.Config, options *servers.ServerOptions) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Options: options,
	}
}
