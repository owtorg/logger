package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Logger exposes eight methods to write logs to the eight RFC 5424 levels
// (debug, info, notice, warning, error, critical, alert, emergency)
// It additionally exposes a Log endpoint that takes the level as a string
type Logger interface {

	//Init should be run before the logger is used.
	//If the logger is added to a stack then init will be called at Add or Set time
	//This is normally used to run the OnInit functions which are added by the end user
	//Which provide a flexible way to add custom initialization to a generic logger implementation
	Init() error

	//Pass in functions that can be called on init
	OnInit(f ...interface{})

	//Log - Generic logging endpoint that can take a string for level, and the data to output
	Log(level string, v ...interface{})

	//PSR-3
	//The following levels correspond to the PHP PSR-3 log levels

	//Emergency - System is unusable.
	Emergency(v ...interface{})
	//Alert - Action must be taken immediately.
	Alert(v ...interface{})
	//Critical - Critical conditions.
	Critical(v ...interface{})
	//Error - Runtime errors that do not require immediate action but should typically be logged and monitored.
	Error(v ...interface{})
	//Warning - Exceptional occurrences that are not errors. Example: Use of deprecated APIs, poor use of an API, undesirable things that are not necessarily wrong.
	Warning(v ...interface{})
	//Notice - Normal but significant events.
	Notice(v ...interface{})
	//Info - Interesting events.  Example: User logs in, SQL logs.
	Info(v ...interface{})
	//Debug - Detailed debug information.
	Debug(v ...interface{})
}

//LogBase is a generic base that can be used to ease registration of initializers via the generic OnInit function
type LogBase struct {
	initializers []interface{}
}

//OnInit adds initializers to the initializers array
func (l *LogBase) OnInit(f ...interface{}) {
	l.initializers = make([]interface{}, 0)
	l.initializers = append(l.initializers, f...)
}

//Log to fmt
type FmtLog struct {
	LogBase
}

func (s *FmtLog) Init() error {
	for _, fn := range s.initializers {
		funct, ok := fn.(func(s *FmtLog))
		if !ok {
			return errors.New("Init callbacks must have signature func(s Logger)")
		}
		funct(s)
	}
	return nil
}
func (s *FmtLog) Emergency(v ...interface{}) {
	s.Log("Emergency", v...)
}
func (s *FmtLog) Alert(v ...interface{}) {
	s.Log("Alert", v...)
}
func (s *FmtLog) Critical(v ...interface{}) {
	s.Log("Critical", v...)
}
func (s *FmtLog) Error(v ...interface{}) {
	s.Log("Error", v...)
}
func (s *FmtLog) Warning(v ...interface{}) {
	s.Log("Warning", v...)
}
func (s *FmtLog) Notice(v ...interface{}) {
	s.Log("Notice", v...)
}
func (s *FmtLog) Info(v ...interface{}) {
	s.Log("Info", v...)
}
func (s *FmtLog) Debug(v ...interface{}) {
	s.Log("Debug", v...)
}
func (s *FmtLog) Log(level string, v ...interface{}) {
	fmt.Println(level, v)
}

//Log to Log
type StdLog struct {
	LogBase
}

func (s *StdLog) Init() error {
	for _, fn := range s.initializers {
		funct, ok := fn.(func(s *StdLog))
		if !ok {
			return errors.New("Init callbacks must have signature func(s Logger)")
		}
		funct(s)
	}
	return nil
}
func (s *StdLog) Emergency(v ...interface{}) {
	s.Log("Emergency", v...)
}
func (s *StdLog) Alert(v ...interface{}) {
	s.Log("Alert", v...)
}
func (s *StdLog) Critical(v ...interface{}) {
	s.Log("Critical", v...)
}
func (s *StdLog) Error(v ...interface{}) {
	s.Log("Error", v...)
}
func (s *StdLog) Warning(v ...interface{}) {
	s.Log("Warning", v...)
}
func (s *StdLog) Notice(v ...interface{}) {
	s.Log("Notice", v...)
}
func (s *StdLog) Info(v ...interface{}) {
	s.Log("Info", v...)
}
func (s *StdLog) Debug(v ...interface{}) {
	s.Log("Debug", v...)
}
func (s *StdLog) Log(level string, v ...interface{}) {
	log.Println(level, v)
}

//Log to File
type FileLog struct {
	LogBase
	f       *os.File
	logPath string
}

//Init expects the first item passed in to be the log file location.
//If it does not exist ./owtorg-logger will be used
func (s *FileLog) Init() error {
	//Set arbitrary log path, which could be overridden by initializers
	s.logPath = "./owtorg-logger"
	for _, v := range s.initializers {
		funct, ok := v.(func(s *FileLog))
		if !ok {
			return errors.New("Init callbacks must have signature func(s *FileLog)")
		}
		funct(s)
	}
	return nil
}
func (s *FileLog) Emergency(v ...interface{}) {
	s.Log("Emergency", v...)
}
func (s *FileLog) Alert(v ...interface{}) {
	s.Log("Alert", v...)
}
func (s *FileLog) Critical(v ...interface{}) {
	s.Log("Critical", v...)
}
func (s *FileLog) Error(v ...interface{}) {
	s.Log("Error", v...)
}
func (s *FileLog) Warning(v ...interface{}) {
	s.Log("Warning", v...)
}
func (s *FileLog) Notice(v ...interface{}) {
	s.Log("Notice", v...)
}
func (s *FileLog) Info(v ...interface{}) {
	s.Log("Info", v...)
}
func (s *FileLog) Debug(v ...interface{}) {
	s.Log("Debug", v...)
}
func (s *FileLog) Log(level string, v ...interface{}) {
	f, err := os.OpenFile(s.logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(0)
	}
	s.f = f
	defer s.f.Close()

	log.SetOutput(s.f)
	log.Println(level, v)
}
