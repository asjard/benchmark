package main

import (
	"runtime"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	options *ServerOptions
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

func (s *Server) start() error {
	e := echo.New()
	e.GET("/hello", s.handler)
	return e.Start(s.options.Address)
}

func (s *Server) handler(c echo.Context) error {
	if s.options.CpuBound {
		Pow(s.options.Target)
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
