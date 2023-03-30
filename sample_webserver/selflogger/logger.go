package selflogger

import (
	"log"
	"os"
)

var InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
var WarnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
var ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
var CriticalLogger = log.New(os.Stdout, "CRITICAL: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix)
