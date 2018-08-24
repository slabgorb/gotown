package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
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
var currentTime = newCurrentTimeHolder()

type currentTimeHolder struct {
	mux         sync.Mutex
	currentTime map[string]time.Time
}

func newCurrentTimeHolder() *currentTimeHolder {
	return &currentTimeHolder{
		currentTime: make(map[string]time.Time),
	}
}

func (ct *currentTimeHolder) setTime(key string) {
	ct.mux.Lock()
	ct.currentTime[key] = time.Now()
	ct.mux.Unlock()
}

func (ct *currentTimeHolder) getTimeSince(key string) time.Duration {
	return time.Since(ct.currentTime[key])
}

// Logger is a convenience interface in case packages want to pass around a
// logger themselves, rather than just using the package level functions.
type Logger interface {
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
	TimeSet(key string)
	TimeElapsed(key string)
	SetOutput(out io.Writer)
}

type logger struct {
	*log.Logger
}

func (l logger) TimeSet(key string) {
	TimeSet(key)
}

func (l logger) TimeElapsed(key string) {
	TimeElapsed(key)
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
	l.Logger.Output(2, fmt.Sprintf(addReturn(format), v...))
}

// Error logs at the Error level
func (l logger) Error(format string, v ...interface{}) {
	if level < ErrorLevel {
		return
	}
	l.SetPrefix("ERROR ")
	l.Logger.Output(2, fmt.Sprintf(addReturn(format), v...))
}

// Fatal logs at the Fatal level, then crashes and burns with the heat of 1000 suns
func (l logger) Fatal(format string, v ...interface{}) {
	if level < FatalLevel {
		return
	}
	l.SetPrefix("FATAL ")
	l.Logger.Output(2, fmt.Sprintf(addReturn(format), v...))
	panic(fmt.Sprintf(format, v))
}

func (l logger) SetOutput(out io.Writer) {
	l.Logger.SetOutput(out)
}

// New returns a new Logger
func New(out io.Writer) Logger {
	return &logger{log.New(out, "", log.Ldate|log.Ltime)}
}

// SetOutput sets the output for the default logger
func SetOutput(out io.Writer) {
	defaultOut = out
	Default.SetOutput(out)
}

// Info logs using the default logger as a convenience function
func Info(format string, v ...interface{}) {
	Default.Info(format, v...)
}

// Debug logs using the default logger as a convenience function
func Debug(format string, v ...interface{}) {
	Default.Debug(format, v...)
}

// Error logs using the default logger as a convenience function
func Error(format string, v ...interface{}) {
	Default.Error(format, v...)
}

// TimeSet resets the 'stopwatch'. It is used to mark elapsed time along with
// the TimeElapsed function
func TimeSet(key string) {
	currentTime.setTime(key)
}

// TimeElapsed messages the duration since the last TimeSet call. It always log s
// at Debug level. NOT THREADSAFE
func TimeElapsed(key string) {
	elapsed := currentTime.getTimeSince(key)
	Debug("%s took %s", key, elapsed)
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
