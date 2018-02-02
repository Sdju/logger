package logger

import (
	"fmt"
)

type Settings struct {
	enabled bool
}

type Logger struct {
	Settings Settings
	name string
}

func NewLogger(name string) Logger {
	if _, ok := logs[name]; !ok {
		settings := Settings{enabled: true}
		logs[name] = Logger{name: name, Settings: settings}
	}
	return logs[name]
}

func (l Logger) Log(a ...interface{}) {
	if !l.Settings.enabled {
		return
	}
	fmt.Println(a...)
}

var logs = make(map[string]Logger)
