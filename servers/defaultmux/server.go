package defaultmux

import (
	"net/http"
	"runtime"
	"time"

	"github.com/asjard/benchmark/servers"
	"github.com/asjard/benchmark/utils"
)

type Server struct {
	options *servers.ServerOptions
}

// 启动服务
func StartServer(options *servers.ServerOptions) error {
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
		utils.Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	w.Write([]byte("hello"))
}
