package ch4_1

import (
	"fmt"
	"runtime"
	"testing"
)

type Log string

func (l *Log) Input(s string) {
	*l = Log(s)
}

func (l Log) Output(s string) {
	fmt.Println(l)
}

type LeveledLog struct {
	Log
	header string
	level  int
}

func NewLeveledLog() *LeveledLog {
	leveledLog := new(LeveledLog)
	_, leveledLog.header, _, _ = runtime.Caller(1)
	leveledLog.level = 0
	return leveledLog
}

func (l *LeveledLog) SetLevel(level int) {
	l.level = level
}

//覆盖Log Ouput方法
func (l LeveledLog) Output() {
	if l.level <= 0 {
		fmt.Print("DEBUG: ")
	} else if l.level == 1 {
		fmt.Print("INFO: ")
	} else if l.level == 2 {
		fmt.Print("WARNING: ")
	} else {
		fmt.Print("ERROR: ")
	}
	fmt.Println(l.header, ":", l.Log)
}

func TestInheritance(t *testing.T) {
	llog := NewLeveledLog()
	llog.SetLevel(1)
	//调用继承的方法
	llog.Input("Using input method of Log type")
	llog.Output()
}
