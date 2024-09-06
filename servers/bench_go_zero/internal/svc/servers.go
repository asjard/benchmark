package svc

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

var (
	h bool   // show help
	p int    // port
	s string // sleepTime
	c bool   // cpu bound
)

func init() {
	flag.BoolVar(&h, "h", false, "show help")
	flag.IntVar(&p, "p", 7030, "server listen port")
	flag.StringVar(&s, "s", "", "process time")
	flag.BoolVar(&c, "c", false, "cpu bound")
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprint(os.Stderr, `gowebbenchmark version:
gowebbenchmark/1.0.0`)
	flag.PrintDefaults()
}

func NewOptions() *ServerOptions {
	if h {
		flag.Usage()
		os.Exit(1)
	}
	sleepTime := time.Duration(0)
	if s != "" {
		duration, err := time.ParseDuration(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sleepTime = duration
	}
	options := &ServerOptions{
		Target:    15,
		CpuBound:  c,
		SleepTime: sleepTime,
		Address:   fmt.Sprintf("127.0.0.1:%d", p),
	}
	go func() {
		time.Sleep(5 * time.Second)
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		var u uint64 = 1024 * 1024
		fmt.Printf("TotalAlloc: %d\n", mem.TotalAlloc/u)
		fmt.Printf("Alloc: %d\n", mem.Alloc/u)
		fmt.Printf("HeapAlloc: %d\n", mem.HeapAlloc/u)
		fmt.Printf("HeapSys: %d\n", mem.HeapSys/u)
	}()
	return options
}
