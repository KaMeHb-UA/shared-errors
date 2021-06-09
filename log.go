package sharederrs

import (
	"log"
	"os"
)

const (
	// LogLevelTrace - describing events showing step by step execution
	LogLevelTrace = "TRACE"
	// LogLevelDebug - events considered to be useful during software debugging
	LogLevelDebug = "DEBUG"
	// LogLevelInfo - can be ignored during normal operations
	LogLevelInfo = "INFO"
	// LogLevelWarn - enexpected behavior happened inside the application
	LogLevelWarn = "WARN"
	// LogLevelError - one or more functionalities are not working
	LogLevelError = "ERROR"
	// LogLevelFatal - system doesn’t fulfill the business functionalities
	LogLevelFatal = "FATAL"
)

// ExtendedLog - extended work with log
type ExtendedLog struct{}

// Log - creates a new instance of the extended log
func Log() *ExtendedLog {
	return &ExtendedLog{}
}

func (l *ExtendedLog) write(logLevel string, errInfo string) {
	if errInfo == "" {
		return
	}
	log.Println("[" + logLevel + "] " + errInfo)
	if logLevel == LogLevelFatal {
		os.Exit(1)
	}
}

// TraceOnError - if there is an error, add information about it to the log and continue execution
func (l *ExtendedLog) TraceOnError(err *APIError) {
	l.write(
		LogLevelWarn,
		err.Message+", stack: "+err.GetStack(),
	)
}

// DebugOnError - if there is an error, add information about it to the log and continue execution
func (l *ExtendedLog) DebugOnError(err *APIError) {
	l.write(
		LogLevelDebug,
		err.Message+", stack: "+err.GetStack(),
	)
}

// WarnOnError - if there is an error, add information about it to the log and continue execution
func (l *ExtendedLog) WarnOnError(err error) {
	if err == nil {
		return
	}
	l.write(LogLevelWarn, err.Error())
}

// Error - if there is an error, add information about it to the log and continue execution
func (l *ExtendedLog) Error(err error) {
	if err == nil {
		return
	}
	l.write(LogLevelError, err.Error())
}

// FatalOnError - if there is an error, it will add information about it to the log and abort the execution
func (l *ExtendedLog) FatalOnError(err error) {
	if err == nil {
		return
	}
	l.write(LogLevelFatal, err.Error())
}

// Info - just adds information to the log
func (l *ExtendedLog) Info(info string) {
	l.write(LogLevelInfo, info)
}
