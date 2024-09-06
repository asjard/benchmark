package main

import (
	"runtime"
	"time"

	"github.com/valyala/fasthttp"
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
		Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	ctx.Write([]byte("hello"))
}
