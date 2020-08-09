package ch6

import (
	"fmt"
	"testing"
	"time"
)

//NewTimer阻塞主线程
func TestNewTimer(t *testing.T) {
	fmt.Println(time.Now())
	timer1 := time.NewTimer(time.Second * 2)
	//2秒后会自动将当前时间写入timer1.C, 所以主线程会在阻塞2秒
	m := <-timer1.C
	fmt.Println("timer1 expired:", m)
	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()
	time.Sleep(time.Second * 1)
	stop := timer2.Stop()
	if stop {
		fmt.Println(time.Now())
	}
}
