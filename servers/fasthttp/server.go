package fasthttp

import (
	"runtime"
	"time"

	"github.com/asjard/benchmark/servers"
	"github.com/asjard/benchmark/utils"
	"github.com/valyala/fasthttp"
)

type Server struct {
	options *servers.ServerOptions
}

func StartServer(options *servers.ServerOptions) error {
	server := &Server{options: options}
	return server.start()
}

func (s *Server) start() error {
	server := &fasthttp.Server{
		Handler:                       s.handler,
		GetOnly:                       true,
		NoDefaultDate:                 true,
		NoDefaultContentType:          true,
		DisableHeaderNamesNormalizing: true,
	}
	return server.ListenAndServe(s.options.Address)
}

func (s *Server) handler(ctx *fasthttp.RequestCtx) {
	if s.options.CpuBound {
		utils.Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	ctx.Write([]byte("hello"))
}
