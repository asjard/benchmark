package main

import (
	"context"
	"runtime"
	"time"

	"github.com/asjard/benchmark/servers/bench_kratos/helloworld"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "helloworld"
	// Version is the version of the compiled software.
	// Version = "v1.0.0"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	options *ServerOptions
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	if s.options.CpuBound {
		Pow(s.options.Target)
	} else {
		if s.options.SleepTime > 0 {
			time.Sleep(s.options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	return &helloworld.HelloReply{Message: "Hello kratos"}, nil
}

func main() {
	if err := StartServer(NewOptions()); err != nil {
		panic(err)
	}
}

func StartServer(options *ServerOptions) error {
	s := &server{
		options: options,
	}
	return s.start()
}

func (s *server) start() error {
	httpSrv := http.NewServer(
		http.Address(s.options.Address),
		http.Middleware(
			recovery.Recovery(),
		),
	)
	helloworld.RegisterGreeterHTTPServer(httpSrv, s)

	app := kratos.New(
		kratos.Name(Name),
		kratos.Server(
			httpSrv,
		),
	)
	return app.Run()
}
