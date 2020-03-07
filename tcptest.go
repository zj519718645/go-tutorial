package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	help bool
	v, V bool
	h    string
	p    string
	c    int
)

func init() {
	flag.BoolVar(&help, "help", false, "Show gaping help")
	flag.BoolVar(&v, "v", false, "Show Version and exit")
	flag.BoolVar(&V, "V", false, "Show Version and exit")
	flag.StringVar(&h, "h", "", "Dest host ipaddress")
	flag.StringVar(&p, "p", "", "Dest host port")
	flag.IntVar(&c, "c", 100000000, "Tcp test conuts")
}

// version
func Version() {
	if v || V {
		fmt.Println("gaping  v0.01")
		os.Exit(3)
	}
}

// tcp 端口连通性测试
func TcpConnect(address string) (err error, costTime time.Duration) {
	nets := new(net.Dialer)
	nets.Timeout = time.Millisecond * 3000
	startTime := time.Now()
	conn, err := nets.Dial("tcp", address)
	if err != nil {
		return err, 1000000
	}
	conn.Write([]byte("hello I'm test!"))
	times := time.Since(startTime).Round(time.Microsecond)
	defer conn.Close()
	return nil, times
}

// help
func Help() {
	fmt.Println(`Gaping  v0.01- Copyright (c) 2019 Mike chulinx
Example: pping -h 127.0.0.1 -p 80`)
	flag.PrintDefaults()
	os.Exit(2)
}

// usage
func Usages() {
	if help {
		Help()
	}
}

// 格式化打印输出
func PingPrint(address, port string) (string, error) {
	date := time.Now().Format("2006-01-02 15:04:05")
	err, t := TcpConnect(address)
	if err != nil {
		return "ERROR: connect failed\n", err
	} else {
		return "SUCCESS: connected\n", nil
	}
	time.Sleep(time.Millisecond)
}

func main() {
	flag.Parse()
	address := fmt.Sprintf("%s:%s", h, p)
	Version()
	Usages()
	if h == "" || p == "" {
		Help()
	}
	f, err := os.OpenFile("tcptest.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < c; i++ {
		resp, _ := PingPrint(address, p)
		f.Write([]byte(time.Now().Format("2006/01/02 15:04:05") + "->" + resp))
	}
}
