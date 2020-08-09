package ch7

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	//双向通道
	ch1 := make(chan string, 1)
	ch1 <- "content"
	fmt.Println(<-ch1)

	//单向只写通道
	//ch2 := make(chan<- string, 1)
	//ch2 <- "content2"
	//fmt.Println(<-ch2)

	//无缓存会阻塞
	ch3 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		ch3 <- "content"
	}()
	fmt.Println("开始阻塞")
	fmt.Println(<-ch3)
	fmt.Println("阻塞完成")

}

func TestChannelClose(t *testing.T) {
	var ch1 chan int
	//close(ch1)
	ch1 = make(chan int, 10)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	//ch1 <- 3
	val, ok := <-ch1
	fmt.Println(val, ok)
	val, ok = <-ch1
	fmt.Println(val, ok)
	val, ok = <-ch1
	fmt.Println(val, ok)
}

func TestSelect(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "content"
	}()
	time.Sleep(time.Millisecond * 5)
	select {
	case <-ch1:
		fmt.Println("ch1 received msg")
	case <-ch2:
		fmt.Println("ch2 received msg")
	default:
		fmt.Println("no new msg received")
	}
}

func TestSelectTimeout(t *testing.T) {
	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "msg"
	}()
	select {
	case <-ch1:
		fmt.Println("ch1 received msg")
	case <-time.After(time.Millisecond * 1500):
		fmt.Println("timeout")
	}
}

func TestDeadLock(t *testing.T) {
	messages := make(chan string, 1)
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	}
	go func() {
		fmt.Println("received messages", <-messages)
	}()
}
