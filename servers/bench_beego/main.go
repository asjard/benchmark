package main

import (
	"net/http"
	"runtime"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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
	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()
	mux := beego.NewControllerRegister()
	mux.Get("/hello", s.handler)
	return http.ListenAndServe(s.options.Address, mux)
}

func (s *Server) handler(w *context.Context) {
	if s.options.CpuBound {
		Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	w.WriteString("hello")
}
