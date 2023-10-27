package main

import (
	"errors"
	"log"
	"os"
)

var (
	l loggers
)

func init() {
	flags := log.LstdFlags | log.Lshortfile

	fileInfo, _ := os.OpenFile("log_info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	fileWarn, _ := os.OpenFile("log_warn.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	fileErr, _ := os.OpenFile("log_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	logInfo := log.New(fileInfo, "INFO:\t", flags)
	logWarn := log.New(fileWarn, "WARN:\t", flags)
	logErr := log.New(fileErr, "ERR:\t", flags)

	l = loggers{
		logInfo: logInfo,
		logWarn: logWarn,
		logErr:  logErr,
	}

}

type loggers struct {
	logInfo *log.Logger
	logWarn *log.Logger
	logErr  *log.Logger
}

func main() {

	l.info("123Some information", 1346253)
	l.warn("123Warn about something")
	l.err(errors.New("123Some error"))

}

func (l *loggers) info(v ...interface{}) {
	l.logInfo.Println(v...)
}

func (l *loggers) warn(v ...interface{}) {
	l.logWarn.Println(v...)
}

func (l *loggers) err(v ...interface{}) {
	l.logErr.Println(v...)
}
