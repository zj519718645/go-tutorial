package ch8

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func WriteFile(path string) {
	file01, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer file01.Close()
	var str string
	for i := 0; i < 10; i++ {
		str = fmt.Sprintf("i = %d\n", i)
		data01, err := file01.WriteString(str)
		if err != nil {
			fmt.Println("err = ", err)
			return
		} else {
			fmt.Println("data01 = ", data01)
		}
	}
}

func TestWriteFile(t *testing.T) {
	path := "./ch8_Writer.txt"
	WriteFile(path)
}

func ReadFile(path string) {
	file01, err := os.Open(path)

	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer file01.Close()

	s1 := make([]byte, 2048)
	long, err := file01.Read(s1)

	if err != nil && err != io.EOF {
		fmt.Println("err = ", err)
		return
	} else {
		fmt.Println(string(s1[0:long]))
	}
}

func TestReadFile(t *testing.T) {
	path := "./ch8_Writer.txt"
	ReadFile(path)
}

func TestReadAll(t *testing.T) {
	path := "./ch8_Writer.txt"
	file01, err := os.Open(path)
	if err == nil {
		data02, err02 := ioutil.ReadAll(file01)
		if err02 != nil {
			fmt.Println("err = ", err02)
		} else {
			fmt.Println("data = ", data02)
		}
	}
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func TestCopy(t *testing.T) {
	CopyFile("./ch8Writer_Copy.txt", "./ch8_Writer.txt")
	fmt.Println("复制完成")
}
