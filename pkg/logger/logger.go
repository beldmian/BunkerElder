package logger

import (
	"log"

	"go.uber.org/zap"
)

func ProvideLogger() *zap.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	return l
}
