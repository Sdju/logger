package logger

import (
	"fmt"
)

type Settings struct {
}

type Logger struct {
	settings Settings
}

func NewLogger(name string) Logger {
	if _, ok := logs[name]; !ok {
		logs[name] = Logger{}
	}
	return logs[name]
}

func (l Logger) Log(a ...interface{}) {
	fmt.Println(a...)
}

var logs = make(map[string]Logger)
