package logger

import (
	"fmt"
	"log"
)

const red = string("\033[31m")
const green = string("\033[32m")
const yellow = string("\033[33m")
const blue = string("\033[34m")
const purple = string("\033[35m")
const cyan = string("\033[36m")
const white = string("\033[37m")
const reset = string("\033[0m")

type stdLogger struct {
	instanceName string
}

func NewSTDLogger(instanceName string) *stdLogger {
	return &stdLogger{"[" + instanceName + "]"}
}

func (l *stdLogger) Info(format string, i ...interface{}) {
	fmt.Printf(cyan + l.instanceName + reset)
	fmt.Printf(blue + " [INFO] " + reset)
	log.Printf(format, i...)
}

func (l *stdLogger) Warn(format string, i ...interface{}) {
	fmt.Printf(cyan + l.instanceName + reset)
	fmt.Printf(yellow + " [WARNING] " + reset)
	log.Printf(format, i...)

}

func (l *stdLogger) Err(format string, i ...interface{}) {
	fmt.Printf(cyan + l.instanceName + reset)
	fmt.Printf(red + " [ERROR] " + reset)
	log.Printf(format, i...)

}

func (l *stdLogger) Debug(format string, i ...interface{}) {
	fmt.Printf(cyan + l.instanceName + reset)
	fmt.Printf(purple + " [DEBUG] " + reset)
	log.Printf(format, i...)
}
