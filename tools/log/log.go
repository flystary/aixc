package log

import (
	"io"
	"log"
	"os"
)

const flags = log.Ldate | log.Ltime

type Logger struct {
	file  *os.File
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
}

var (
	logFile       io.Writer
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func New(path string) (*Logger, error) {
	if path == "" {
		path = "./aixc.log"
	}

	fp, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("create log file err %+v", err)
	}

	return &Logger{
		file:  fp,
		debug: log.New(fp, "[DEBUG] ", flags),
		info:  log.New(fp, "[INFO] ", flags),
		warn:  log.New(fp, "[WARN] ", flags),
		err:   log.New(fp, "[ERROR] ", flags),
	}, nil
}

func (l *Logger) Debugf(format string, v ...any) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...any) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.err.Printf(format, v...)
}

func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}
