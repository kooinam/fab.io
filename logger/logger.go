package logger

import (
	"fmt"
	"log"
	"os"
)

var base = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

func doPrintf(format string, a ...interface{}) {
	base.Output(3, fmt.Sprintf(format, a...))
}

func Debug(format string, a ...interface{}) {
	doPrintf(format, a...)
}
