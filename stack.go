package logger

import "errors"

//Stack - A stack is a group of loggers that also implements the logger interface
//loggers will be called in the order they are added
type Stack struct {
	LogBase
	loggers []interface{}
}

//Add a logger to the stack
func (s *Stack) Add(l ...interface{}) {
	//Init the loggers and then add them to the stack
	for _, v := range l {
		lg := v.(Logger)
		lg.Init()
	}
	s.loggers = append(s.loggers, l...)
}

//Set the loggers in the stack
func (s *Stack) Set(Loggers []interface{}) {
	s.loggers = Loggers
	//Initialize all the loggers
	for _, v := range s.loggers {
		lg := v.(Logger)
		lg.Init()
	}
}

//Init - expects input to be a list of func(s *Stack) which will be called on initialization
func (s *Stack) Init() error {
	s.loggers = make([]interface{}, 1)
	for _, fn := range s.initializers {
		funct, ok := fn.(func(s Logger))
		if !ok {
			return errors.New("Init callbacks must have signature func(s Logger)")
		}
		funct(s)
	}
	return nil
}

func (s *Stack) Emergency(v ...interface{}) {
	for _, lg := range s.loggers {
		lg := lg.(Logger)
		lg.Emergency(v...)
	}
}
func (s *Stack) Alert(v ...interface{}) {
	for _, lg := range s.loggers {
		lg := lg.(Logger)
		lg.Alert(v...)
	}
}
func (s *Stack) Critical(v ...interface{}) {
	for _, lg := range s.loggers {
		lg := lg.(Logger)
		lg.Critical(v...)
	}
}
func (s *Stack) Error(v ...interface{}) {
	for _, lg := range s.loggers {
		lg := lg.(Logger)
		lg.Error(v...)
	}
}
func (s *Stack) Warning(v ...interface{}) {
	for _, lg := range s.loggers {
		lg := lg.(Logger)
		lg.Warning(v...)
	}
}
func (s *Stack) Notice(v ...interface{}) {
	for _, lg := range s.loggers {
		lg := lg.(Logger)
		lg.Notice(v...)
	}
}
func (s *Stack) Info(v ...interface{}) {
	for _, lg := range s.loggers {
		lg := lg.(Logger)
		lg.Info(v...)
	}
}
func (s *Stack) Debug(v ...interface{}) {
	for _, lg := range s.loggers {
		lg := lg.(Logger)
		lg.Debug(v...)
	}
}
func (s *Stack) Log(level string, v ...interface{}) {
	for _, lg := range s.loggers {
		lg := lg.(Logger)
		lg.Log(level, v...)
	}
}
