package logger

import (
	"fmt"
	"go.uber.org/zap"
)

type Logger interface {
	Info(string)
	InfoW(string, map[any]any)
	Debug(string)
	DebugW(string, map[any]any)
	Warn(string, error)
	WarnW(string, error)
}

type customLogger struct {
	logger *zap.Logger
}

func NewCustomLogger(loggerType string) (Logger, error) {
	var logger *zap.Logger
	var err error
	if loggerType == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}

	return &customLogger{logger: logger}, nil
}

func (l *customLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *customLogger) InfoW(msg string, fields map[any]any) {
	l.logger.Info(msg, convertFields(fields)...)
}

func (l *customLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *customLogger) DebugW(msg string, fields map[any]any) {
	l.logger.Debug(msg, convertFields(fields)...)
}

func (l *customLogger) Warn(msg string, err error) {
	if err != nil {
		l.logger.Warn(msg, zap.Error(err))
	} else {
		l.logger.Warn(msg)
	}
}

func (l *customLogger) WarnW(msg string, err error) {
	if err != nil {
		l.logger.Warn(msg, zap.Error(err))
	} else {
		l.logger.Warn(msg)
	}
}

func convertFields(fields map[any]any) []zap.Field {
	if fields == nil {
		return nil
	}
	f := make([]zap.Field, len(fields))
	index := 0
	for k, v := range fields {
		f[index] = zap.Any(fmt.Sprint(k), v)
		index++
	}
	return f
}
