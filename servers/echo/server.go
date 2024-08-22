package echo

import (
	"runtime"
	"time"

	"github.com/asjard/benchmark/servers"
	"github.com/asjard/benchmark/utils"
	"github.com/labstack/echo/v4"
)

type Server struct {
	options *servers.ServerOptions
}

func StartServer(options *servers.ServerOptions) error {
	server := &Server{options: options}
	return server.start()
}

func (s *Server) start() error {
	e := echo.New()
	e.GET("/hello", s.handler)
	return e.Start(s.options.Address)
}

func (s *Server) handler(c echo.Context) error {
	if s.options.CpuBound {
		utils.Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	c.Response().Write([]byte("hello"))
	return nil
}
