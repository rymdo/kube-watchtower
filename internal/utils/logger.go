package utils

import (
	"go.uber.org/zap"
)

func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return &Logger{
		logger.Sugar(),
	}
}
