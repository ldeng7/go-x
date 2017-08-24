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

var levelsStr = map[int]string{
	INFO:   "[\x1b[1;42;37mINFO\x1b[0m]",
	NOTICE: "[\x1b[1;42;37mNOTICE\x1b[0m]",
	WARN:   "[\x1b[1;43;37mWARN\x1b[0m]",
	ERR:    "[\x1b[1;41;37mERR\x1b[0m]",
	CRIT:   "[\x1b[1;45;37mCRIT\x1b[0m]",
	EMERG:  "[\x1b[1;45;37mEMERG\x1b[0m]",
}

const (
	fileFlag = os.O_WRONLY | os.O_APPEND | os.O_CREATE
	fileMode = 0666
)

type InitArgs struct {
	Writer   io.Writer
	Filename string
	Prefix   string
	Flags    int
	LogLevel int
}

type Logger struct {
	Logger   *log.Logger
	LogLevel int
	writer   io.Writer
}

func Init(args *InitArgs) (*Logger, error) {
	self := &Logger{
		LogLevel: args.LogLevel,
	}

	writer := args.Writer
	if nil == writer {
		file, err := os.OpenFile(args.Filename, fileFlag, fileMode)
		if nil != err {
			return nil, err
		}
		writer = file
	}

	self.writer = writer
	self.Logger = log.New(writer, args.Prefix, args.Flags)
	return self, nil
}

func (l *Logger) GetWriter() io.Writer {
	return l.writer
}

func (l *Logger) SetWriter(writer io.Writer) {
	l.writer = writer
	l.Logger.SetOutput(writer)
}

func (l *Logger) FileRotate() (*os.File, error) {
	file, ok := l.writer.(*os.File)
	if !ok {
		return nil, nil
	}
	fileNew, err := os.OpenFile(file.Name(), fileFlag, fileMode)
	if nil != err {
		return nil, err
	}
	l.SetWriter(fileNew)
	return file, nil
}

func (l *Logger) log(calldepth int, level int, v []interface{}) {
	if nil == l.Logger {
		fmt.Println(v)
		return
	}
	if level < l.LogLevel {
		return
	}
	levelStr, ok := levelsStr[level]
	if !ok {
		levelStr = fmt.Sprintf("[level %d]", level)
	}
	l.Logger.Output(calldepth, fmt.Sprintf("%s%s", levelStr, fmt.Sprintln(v...)))
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
