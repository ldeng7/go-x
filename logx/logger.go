package logx

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	INFO = iota
	NOTICE
	WARN
	ERR
	CRIT
	EMERG
)

const (
	fileFlag = os.O_WRONLY | os.O_APPEND | os.O_CREATE
	fileMode = 0666
)

type InitArgs struct {
	Out      io.Writer
	Filename string
	Prefix   string
	Flags    int
	LogLevel int
}

type Logger struct {
	Logger   *log.Logger
	Out      io.Writer
	logLevel int
}

func Init(args *InitArgs) error {
	self := &Logger{
		logLevel: args.LogLevel,
	}

	out := args.Out
	if nil == out {
		file, err := os.OpenFile(args.Filename, fileFlag, fileMode)
		if nil != err {
			return err
		}
		out = file
	}

	self.Out = out
	self.Logger = log.New(out, args.Prefix, args.Flags)
	return nil
}

func (l *Logger) FileRotate() (*os.File, error) {
	file, ok := l.Out.(*os.File)
	if !ok {
		return nil, nil
	}
	fileNew, err := os.OpenFile(file.Name(), fileFlag, fileMode)
	if nil != err {
		return nil, err
	}
	l.Out = fileNew
	l.Logger.SetOutput(fileNew)
	return file, nil
}

func (l *Logger) log(calldepth int, level int, v []interface{}) {
	if nil == l.Logger {
		fmt.Println(v)
		return
	}
	if level < l.logLevel {
		return
	}
	l.Logger.Output(calldepth, fmt.Sprintln(v...))
}

func (l *Logger) Log(level int, v ...interface{}) {
	l.log(3, level, v)
}

func (l *Logger) Info(v ...interface{}) {
	l.log(3, INFO, v)
}

func (l *Logger) Notice(v ...interface{}) {
	l.log(3, NOTICE, v)
}

func (l *Logger) Warn(v ...interface{}) {
	l.log(3, WARN, v)
}

func (l *Logger) Err(v ...interface{}) {
	l.log(3, ERR, v)
}

func (l *Logger) Crit(v ...interface{}) {
	l.log(3, CRIT, v)
}

func (l *Logger) Emerg(v ...interface{}) {
	l.log(3, EMERG, v)
}
