package main

import (
	"go.uber.org/zap"
)

func foo(log *zap.Logger) {
	log.Error("foo func ERROR")
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("Hello logger, INFO")
	logger.Debug("Hello logger, DEBUG")
	logger.Warn("Hello logger, WARN")
	logger.Error("Hello logger, ERROR")

	foo(logger)
}
