package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(env string) error {
	var config zap.Config

	// init logger config
	if env == "development" {
		config = zap.NewDevelopmentConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	} else {
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	// build logger
	log, err := config.Build()
	if err != nil {
		return err
	}

	logger = log
	return nil
}

func GetLogger() *zap.Logger {
	return logger
}
