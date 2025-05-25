package logger

import (
	"fmt"
	"io"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type logger struct {
	out io.Writer
}

var singleLogger *logger;

func GetLoggerInstance(out io.Writer) *logger {
	if singleLogger == nil {
		lock.Lock()
		defer lock.Unlock()

		if (singleLogger == nil) {
			singleLogger = &logger{
				out: out,
			}
			fmt.Println("Logger instance is created!")
		} else {
			fmt.Println("Logger instance is already created!")
		}
	} else {
		fmt.Println("Logger instance is already created!")
	}

	return singleLogger
}

func resetSingleLogger() {
	singleLogger = nil
}

func (lg logger) Info(message string) {
	lg.outputLog(INFO, message)
}

func (lg logger) Debug(message string) {
	lg.outputLog(DEBUG, message)
}

func (lg logger) Error(message string) {
	lg.outputLog(ERROR, message)
}

func (lg logger) Warn(message string) {
	lg.outputLog(WARN, message)
}

func (lg logger) outputLog(l level, message string) {
	log := log{
		timestamp: time.Now(),
		level: l,
		message: message,
	}

	lg.out.Write([]byte(log.String()))
}
