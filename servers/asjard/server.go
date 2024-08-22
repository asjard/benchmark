package asjard

import (
	"context"
	"runtime"
	"time"

	"github.com/asjard/asjard"
	"github.com/asjard/asjard/pkg/server/rest"
	"github.com/asjard/benchmark/servers"
	"github.com/asjard/benchmark/servers/asjard/protobuf/hellopb"
	"github.com/asjard/benchmark/utils"
)

type Server struct {
	options *servers.ServerOptions
	hellopb.UnimplementedHelloServer
}

func StartServer(options *servers.ServerOptions) error {
	server := &Server{options: options}
	return server.start()
}

func (s *Server) Say(ctx context.Context, in *hellopb.HelloReq) (*hellopb.HelloReq, error) {
	if s.options.CpuBound {
		utils.Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	return &hellopb.HelloReq{Message: "hello"}, nil
}

func (s *Server) RestServiceDesc() *rest.ServiceDesc {
	return &hellopb.HelloRestServiceDesc
}

func (s *Server) start() error {
	server := asjard.New()
	server.AddHandler(s, rest.Protocol)
	return server.Start()
}
