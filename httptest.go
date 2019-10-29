package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url = "http://lw.technicalsupport.cn/1.html"
	for {
		resp, _ := Getresp(url)
		fmt.Println(resp)
	}
}

func Getresp(api string) (string, error) {
	resp, err := http.Get(api)
	if err != nil {
		return "connect failed", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get info didn’t respond 200 OK: %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "get response failed", err
	}
	return string(body), nil
}
