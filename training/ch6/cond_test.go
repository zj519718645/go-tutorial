package ch6

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)
var wt sync.WaitGroup

func test(x int) {
	cond.L.Lock()
	fmt.Println("协程", x, "已经被锁")
	cond.Wait()
	fmt.Println("协程", x, "已经被通知可以执行")
	defer func() {
		cond.L.Unlock()
		wt.Done()
	}()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go test(i)
		wt.Add(1)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("------------开始一个一个通知解锁--------------")
	cond.Signal()
	time.Sleep(time.Second * 1)
	cond.Signal()
	time.Sleep(time.Second * 1)
	fmt.Println("-------------------开始广播-------------------")
	cond.Broadcast()
	wt.Wait()
}
