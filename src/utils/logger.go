package utils

import "go.uber.org/zap"

func GetLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()

	return logger.Sugar()
}
