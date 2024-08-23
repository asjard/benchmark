package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/asjard/benchmark/servers"
	"github.com/asjard/benchmark/servers/asjard"
	"github.com/asjard/benchmark/servers/beego"
	"github.com/asjard/benchmark/servers/defaultmux"
	"github.com/asjard/benchmark/servers/echo"
	"github.com/asjard/benchmark/servers/fasthttp"
	"github.com/asjard/benchmark/servers/gin"
	"github.com/asjard/benchmark/servers/go_chassis"
	"github.com/asjard/benchmark/servers/go_zero"
	// "github.com/qiangxue/fasthttp-routing"
)

var (
	port              = 7030
	sleepTime         = 0
	cpuBound          bool
	target            = 15
	sleepTimeDuration time.Duration
	message           = []byte("hello world")
	messageStr        = "hello world"
	samplingPoint     = 10 // seconds
)

// server [default] [10] [8080]
func main() {
	args := os.Args
	argsLen := len(args)
	webFramework := "default"
	if argsLen > 1 {
		webFramework = args[1]
	}
	if argsLen > 2 {
		sleepTime, _ = strconv.Atoi(args[2])
		if sleepTime == -1 {
			cpuBound = true
			sleepTime = 0
		}
	}
	if argsLen > 3 {
		port, _ = strconv.Atoi(args[3])
	}
	if argsLen > 4 {
		samplingPoint, _ = strconv.Atoi(args[4])
	}
	sleepTimeDuration = time.Duration(sleepTime) * time.Millisecond
	samplingPointDuration := time.Duration(samplingPoint) * time.Second

	go func() {
		time.Sleep(samplingPointDuration)
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		var u uint64 = 1024 * 1024
		fmt.Printf("TotalAlloc: %d\n", mem.TotalAlloc/u)
		fmt.Printf("Alloc: %d\n", mem.Alloc/u)
		fmt.Printf("HeapAlloc: %d\n", mem.HeapAlloc/u)
		fmt.Printf("HeapSys: %d\n", mem.HeapSys/u)
	}()
	options := &servers.ServerOptions{
		Target:    target,
		CpuBound:  cpuBound,
		SleepTime: time.Duration(sleepTime) * time.Millisecond,
		Address:   ":" + strconv.Itoa(port),
	}
	var err error
	switch webFramework {
	case "default":
		err = defaultmux.StartServer(options)
	case "beego":
		err = beego.StartServer(options)
	case "echo":
		err = echo.StartServer(options)
	case "fasthttp":
		err = fasthttp.StartServer(options)
	case "gin":
		err = gin.StartServer(options)
	case "asjard":
		err = asjard.StartServer(options)
	case "go_zero":
		err = go_zero.StartServer(options)
	case "go_chassis":
		err = go_chassis.StartServer(options)
	// case "jupiter":
	// 	err = jupiter.StartServer(options)
	default:
		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("------------- Unknown framework given!!! Check libs.sh -------------")
		fmt.Println("------------- Unknown framework given!!! Check libs.sh -------------")
		fmt.Println("------------- Unknown framework given!!! Check libs.sh -------------")
		fmt.Println("--------------------------------------------------------------------")
	}
	if err != nil {
		log.Fatal(err)
	}
}
