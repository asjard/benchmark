package go_zero

import (
	"fmt"
	"os"

	"github.com/asjard/benchmark/servers"
	"github.com/asjard/benchmark/servers/go_zero/internal/config"
	"github.com/asjard/benchmark/servers/go_zero/internal/handler"
	"github.com/asjard/benchmark/servers/go_zero/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Server struct {
	options *servers.ServerOptions
}

func StartServer(options *servers.ServerOptions) error {
	server := &Server{options: options}
	return server.start()
}

func (s *Server) start() error {
	var c config.Config
	conf.MustLoad(os.Getenv("GO_ZERO_CONF_FILE"), &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c, s.options)
	handler.RegisterHandlers(server, ctx)
	logx.Disable()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	return nil
}
