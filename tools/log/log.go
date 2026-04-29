package log

import (
	"log"
	"os"
	"sync"
)

const flags = log.Ldate | log.Ltime

var (
	once sync.Once
	fp   *os.File

	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func Init(path string) error {
	var err error
	if path == "" {
		path = "./aixc.log"
	}

	once.Do(func() {
		fp, err = os.OpenFile(path,
			os.O_CREATE|os.O_APPEND|os.O_RDWR,
			0666,
		)
		if err != nil {
			return
		}
		debugLogger = log.New(fp, "[DEBUG] ", flags)
		infoLogger = log.New(fp, "[INFO] ", flags)
		warningLogger = log.New(fp, "[WARN] ", flags)
		errorLogger = log.New(fp, "[ERROR] ", flags)
	})

	return err
}
func Debugf(format string, v ...any) {
	if debugLogger != nil {
		debugLogger.Printf(format, v...)
	}
}

func Infof(format string, v ...any) {
	if infoLogger != nil {
		infoLogger.Printf(format, v...)
	}
}

func Warningf(format string, v ...any) {
	if warningLogger != nil {
		warningLogger.Printf(format, v...)
	}
}

func Errorf(format string, v ...any) {
	if errorLogger != nil {
		errorLogger.Printf(format, v...)
	}
}

func Close() error {
	if fp != nil {
		return fp.Close()
	}
	return nil
}
