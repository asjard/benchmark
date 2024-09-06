package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/asjard/benchmark/servers/bench_go_zero/internal/config"
	"github.com/asjard/benchmark/servers/bench_go_zero/internal/handler"
	"github.com/asjard/benchmark/servers/bench_go_zero/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Server struct {
	options *svc.ServerOptions
}

func main() {
	if err := StartServer(svc.NewOptions()); err != nil {
		panic(err)
	}
}

func StartServer(options *svc.ServerOptions) error {
	server := &Server{options: options}
	return server.start()
}

func (s *Server) start() error {
	wd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	var c config.Config
	conf.MustLoad(filepath.Join(wd, "etc", "gozero-api.yaml"), &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c, s.options)
	handler.RegisterHandlers(server, ctx)
	logx.Disable()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	return nil
}
