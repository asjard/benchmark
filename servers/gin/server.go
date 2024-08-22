package gin

import (
	"runtime"
	"time"

	"github.com/asjard/benchmark/servers"
	"github.com/asjard/benchmark/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	options *servers.ServerOptions
}

func StartServer(options *servers.ServerOptions) error {
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
		utils.Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	c.Writer.Write([]byte("hello"))
}
