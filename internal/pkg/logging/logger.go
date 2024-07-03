package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppLogger struct{}

var instance *AppLogger = &AppLogger{}

func GetLogger() *AppLogger {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return instance
}

func (l *AppLogger) Fatalf(message string, v ...any) {
	log.Fatal().Msgf(message, v...)
}

func (l *AppLogger) Errorf(message string, v ...any) {
	log.Error().Msgf(message, v...)
}

func (l *AppLogger) Warnf(message string, v ...any) {
	log.Warn().Msgf(message, v...)
}

func (l *AppLogger) Infof(message string, v ...any) {
	log.Info().Msgf(message, v...)
}

func (l *AppLogger) Debugf(message string, v ...any) {
	log.Debug().Msgf(message, v...)
}

func (l *AppLogger) Tracef(message string, v ...any) {
	log.Trace().Msgf(message, v...)
}

func (l *AppLogger) Fatal(message string, v ...any) {
	l.Fatalf(message, v...)
}

func (l *AppLogger) Error(message string, v ...any) {
	l.Errorf(message, v...)
}

func (l *AppLogger) Warn(message string, v ...any) {
	l.Warnf(message, v...)
}

func (l *AppLogger) Info(message string, v ...any) {
	l.Infof(message, v...)
}

func (l *AppLogger) Debug(message string, v ...any) {
	l.Debugf(message, v...)
}

func (l *AppLogger) Trace(message string, v ...any) {
	l.Tracef(message, v...)
}

func (l *AppLogger) Warningf(message string, v ...any) {
	log.Warn().Msgf(message, v...)
}

func (l *AppLogger) Errors(errs ...error) {
	for _, v := range errs {
		if v != nil {
			l.Error("%s", errs)
		}
	}
}
