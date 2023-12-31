package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func InitLogger() error {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	var err error
	logger, err = config.Build()
	if err != nil {
		return fmt.Errorf("failed to build logger: %w", err)
	}

	zap.RedirectStdLog(logger)

	return nil
}

func GetLogger() *zap.Logger {
	if logger == nil {
		InitLogger() 
	}
	return logger
}
