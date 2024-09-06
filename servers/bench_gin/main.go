package main

import (
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
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
	gin.SetMode(gin.ReleaseMode)
	mux := gin.New()
	mux.GET("/hello", s.handler)
	return mux.Run(s.options.Address)
}

func (s *Server) handler(c *gin.Context) {
	if s.options.CpuBound {
		Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	c.Writer.Write([]byte("hello"))
}
