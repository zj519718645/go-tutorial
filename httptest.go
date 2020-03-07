package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	//var url = "http://118.190.156.110:30080/"
	//var url = "http://lw.technicalsupport.cn/1.html"
	var url string
	var repeats int
	flag.StringVar(&url, "url", "", "api url")
	flag.IntVar(&repeats, "c", 1, "run counts")
	flag.Parse()

	f, err := os.OpenFile("httptest.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < repeats; i++ {
		resp, _ := Getresp(url)
		//fmt.Println(time.Now(), resp)
		f.Write([]byte(time.Now().Format("2006/01/02 15:04:05") + "->" + resp))
		time.Sleep(time.Millisecond)
	}
}

func Getresp(api string) (string, error) {
	resp, err := http.Get(api)
	if err != nil {
		return "ERROR: connect failed\n", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "ERROR: get response failed\n", fmt.Errorf("ERROR: get info didn’t respond 200 OK: %s", resp.Status)
	}
	return "SUCCESS: connected\n", nil
}
