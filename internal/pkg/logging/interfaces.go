package logging

//counterfeiter:generate . Logger
type Logger interface {
	Fatalf(message string, v ...any)
	Errorf(message string, v ...any)
	Warnf(message string, v ...any)
	Infof(message string, v ...any)
	Debugf(message string, v ...any)
	Tracef(message string, v ...any)

	Fatal(message string, v ...any)
	Error(message string, v ...any)
	Warn(message string, v ...any)
	Info(message string, v ...any)
	Debug(message string, v ...any)
	Trace(message string, v ...any)

	Warningf(message string, v ...any)
	Errors(errs ...error)
}
