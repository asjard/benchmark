package main

import (
	"context"
	"runtime"
	"time"

	"github.com/asjard/asjard"
	"github.com/asjard/asjard/pkg/server/rest"
	"github.com/asjard/benchmark/servers/bench_asjard/protobuf/hellopb"
)

type Server struct {
	options *ServerOptions
	hellopb.UnimplementedHelloServer
}

func main() {
	if err := StartServer(NewOptions()); err != nil {
		panic(err)
	}
}

func StartServer(options *ServerOptions) error {
	server := &Server{options: options}
	return server.start()
}

func (s *Server) Say(ctx context.Context, in *hellopb.HelloReq) (*hellopb.HelloReq, error) {
	if s.options.CpuBound {
		Pow(s.options.Target)
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
