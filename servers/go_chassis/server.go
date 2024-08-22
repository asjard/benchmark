package go_chassis

import (
	"net/http"
	"runtime"
	"time"

	"github.com/asjard/benchmark/servers"
	"github.com/asjard/benchmark/utils"
	"github.com/go-chassis/go-chassis/v2"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

type Server struct {
	options *servers.ServerOptions
}

func StartServer(options *servers.ServerOptions) error {
	server := &Server{options: options}
	return server.start()
}

func (s *Server) start() error {
	chassis.RegisterSchema("rest", s)
	if err := chassis.Init(); err != nil {
		return err
	}
	return chassis.Run()
}

// URLPatterns helps to respond for corresponding API calls
func (s *Server) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/hello", ResourceFunc: s.Hello},
	}
}

// Hello
func (s *Server) Hello(b *rf.Context) {
	if s.options.CpuBound {
		utils.Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	b.Write([]byte("hi from hello"))
}
