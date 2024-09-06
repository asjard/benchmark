package svc

import (
	"github.com/asjard/benchmark/servers/bench_go_zero/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	Options *ServerOptions
}

func NewServiceContext(c config.Config, options *ServerOptions) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Options: options,
	}
}
