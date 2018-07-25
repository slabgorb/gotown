package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// log levels
const (
	DebugLevel = 1 << iota
	InfoLevel
	ErrorLevel
	FatalLevel
	NoneLevel
)

var level = DebugLevel
var Default Logger
var defaultOut io.Writer = os.Stdout
var currentTime = time.Now()

// Logger is a convenience interface in case packages want to pass around a
// logger themselves, rather than just using the package level functions.
type Logger interface {
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
	SetOutput(out io.Writer)
}

type logger struct {
	*log.Logger
}

// Info logs at the Info level
func (l logger) Info(format string, v ...interface{}) {
	if level < InfoLevel {
		return
	}
	l.SetPrefix("INFO ")
	l.Logger.Printf(addReturn(format), v...)
}

// Debug logs at the Debug level
func (l logger) Debug(format string, v ...interface{}) {
	if level < DebugLevel {
		return
	}
	l.SetPrefix("DEBUG ")
	l.Logger.Printf(addReturn(format), v...)
}

// Error logs at the Error level
func (l logger) Error(format string, v ...interface{}) {
	if level < ErrorLevel {
		return
	}
	l.SetPrefix("ERROR ")
	l.Logger.Printf(addReturn(format), v...)
}

// Fatal logs at the Fatal level, then crashes and burns with the heat of 1000 suns
func (l logger) Fatal(format string, v ...interface{}) {
	if level < FatalLevel {
		return
	}
	l.SetPrefix("FATAL ")
	l.Logger.Printf(addReturn(format), v...)
	panic(fmt.Sprintf(format, v))
}

func (l logger) SetOutput(out io.Writer) {
	l.Logger.SetOutput(out)
}

// New returns a new Logger
func New(out io.Writer) Logger {
	return &logger{log.New(out, "", log.Ldate|log.Lshortfile|log.Ltime)}
}

// SetOutput sets the output for the default logger
func SetOutput(out io.Writer) {
	defaultOut = out
	Default.SetOutput(out)
}

// Info logs using the default logger as a convenience function
func Info(format string, v ...interface{}) {
	Default.Info(format, v)
}

// Debug logs using the default logger as a convenience function
func Debug(format string, v ...interface{}) {
	Default.Debug(format, v)
}

// Error logs using the default logger as a convenience function
func Error(format string, v ...interface{}) {
	Default.Error(format, v)
}

// TimeSet resets the 'stopwatch'. It is used to mark elapsed time along with
// the TimeElapsed function
func TimeSet() {
	currentTime = time.Now()
}

// TimeElapsed messages the duration since the last TimeSet call. It always logs
// at Debug level.
func TimeElapsed(name string) {
	elapsed := time.Since(currentTime)
	Debug("%s took %s", name, elapsed)
}

// SetLogLevel sets the current logging level
func SetLogLevel(l int) {
	if level > NoneLevel || level < 0 {
		return
	}
	level = l
}

func init() {
	Default = New(defaultOut)
}

func addReturn(format string) string {
	return format + "\n"
}
