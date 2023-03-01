package log

import (
	"io"
	"log"
	"os"
)

const (
	flag       = log.Ldate | log.Ltime
	preDebug   = "[DEBUG]"
	preInfo    = "[INFO]"
	preWarning = "[WARNING]"
	preError   = "[ERROR]"
)

var (
	logFile        io.Writer
	debugLogger    *log.Logger
	infoLogger     *log.Logger
	warningLogger  *log.Logger
	errorLogger    *log.Logger
	defaultLogFile = "/var/log/aixc.log"
)

func init() {
	var err error
	logFile, err = os.OpenFile(defaultLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		defaultLogFile = "./aixc.log"
		logFile, err = os.OpenFile(defaultLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalf("create log file err %+v", err)
		}
	}
	debugLogger = log.New(logFile, preDebug, flag)
	infoLogger = log.New(logFile, preInfo, flag)
	warningLogger = log.New(logFile, preWarning, flag)
	errorLogger = log.New(logFile, preError, flag)
}

func Debugf(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

func Infof(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

func Warningf(format string, v ...interface{}) {
	warningLogger.Printf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}

func SetOutputPath(path string) {
	var err error
	logFile, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {

		log.Fatalf("create log file err %+v", err)
	}
	debugLogger.SetOutput(logFile)
	infoLogger.SetOutput(logFile)
	warningLogger.SetOutput(logFile)
	errorLogger.SetOutput(logFile)
}
