package main

import (
	"net/http"
	"runtime"
	"time"
)

type Server struct {
	options *ServerOptions
}

func main() {
	if err := StartServer(NewOptions()); err != nil {
		panic(err)
	}
}

// 启动服务
func StartServer(options *ServerOptions) error {
	server := &Server{
		options: options,
	}
	return server.start()
}

func (s *Server) start() error {
	http.HandleFunc("/hello", s.handler)
	return http.ListenAndServe(s.options.Address, nil)
}

func (s *Server) handler(w http.ResponseWriter, _ *http.Request) {
	if s.options.CpuBound {
		Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	w.Write([]byte("hello"))
}
