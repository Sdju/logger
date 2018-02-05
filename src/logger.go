package logger

import (
	"fmt"
)

const basePrelude = "%NAME> "

type Settings struct {
	Enabled bool
	HasPrelude bool
	PreludeString string
}

type Logger struct {
	Settings Settings
	name string
}

func NewLogger(name string) Logger {
	if _, ok := logs[name]; !ok {
		settings := Settings{Enabled: true, HasPrelude: true}
		logs[name] = Logger{name: name, Settings: settings}
	}
	return logs[name]
}

func (l Logger) Log(a ...interface{}) {
	if !l.Settings.Enabled {
		return
	}
	if l.Settings.HasPrelude {
		fmt.Print(l.name, "> ")
	}
	fmt.Println(a...)
}

func (l Logger) Name()string {
	return l.name
}

var logs = make(map[string]Logger)
