package ch8

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {

	file := time.Now().Format("2006年1月2日") + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		fmt.Println("生成日志文件出错", err)
	}
	logger := log.New(logFile, "这是一个前缀", log.Ldate|log.Lshortfile)
	logger.Print("打印第一条日志消息")
	fmt.Println("后台输出修改前flag")
	fmt.Println(logger.Flags())
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("后台输出修改后flag")
	fmt.Println(logger.Flags())
	fmt.Println(logger.Prefix())
	logger.SetPrefix("这是修改后的前缀")
	logger.Output(2, "打印第二条日志信息\n")

}

func DivideTestLog(dividend float64, divisor float64) (result float64, err error) {
	if divisor == 0 {
		result = 0.0
		err = errors.New("runtime error: divide by zero")
		return
	}
	result = dividend / divisor
	err = nil
	return
}

func DivideForRecoverLog(dividend float64, divisor float64) (result float64) {
	if divisor == 0 {
		err := errors.New("panic runtime error: divide by zero")
		panic(err)
	}
	result = dividend / divisor
	return
}

func TestLogger2(t *testing.T) {
	file := time.Now().Format("20020115") + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		fmt.Println("生成日志是出现错误信息", err)
	}
	logger := log.New(logFile, "自定义错误处理测试", log.LstdFlags|log.Lshortfile)
	r, err1 := DivideTestLog(6, 0)
	if nil == err1 {
		fmt.Println(r)
	} else {
		logger.Print(err1, "\n")
	}
	logger.SetPrefix("panic 错误处理测试")
	defer func() {
		if err2 := recover(); err2 != nil {
			logger.Print(err2, "\n")
		}
	}()
	r1 := DivideForRecoverLog(5, 0)
	fmt.Println("recover result ", r1)
}
