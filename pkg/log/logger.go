package log

import (
	"fmt"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

// InitLogger initializes the logger.
func InitLogger() error {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		return err
	}
	return nil
}

// Info logs an info level message.
func Info(msg string) {
	logger.Info(msg)
}

// Error logs an error level message.
func Error(msg string) {
	logger.Error(msg)
}

// Error logs an error level message.
func Errorf(msg string, args ...interface{}) {
	logger.Error(fmt.Sprintf(msg, args...))
}

// CloseLogger closes the logger.
func CloseLogger() error {
	err := logger.Sync()
	if err != nil {
		return err
	}
	return nil
}
