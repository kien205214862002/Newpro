package main

import (
	"go01-airbnb/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	sugarLogger := logger.NewZapLogger()

	sugarLogger.Error("Hello World!!", zap.String("Data", "OK"))
}
