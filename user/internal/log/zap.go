package log

import (
	"go.uber.org/zap"
	"os"
)

type Logger struct {
	*zap.Logger
}

func GetLogger(env string) Logger {
	env = os.Getenv("ENV")

	var (
		log *zap.Logger
		err error
	)

	switch env {
	case "prod":
		log, err = zap.NewProduction()
		if err != nil {
			panic("failed to get logger: " + err.Error())
		}
	default:
		log, err = zap.NewDevelopment()
		if err != nil {
			panic("failed to get logger: " + err.Error())
		}
	}
	log.Info("successfully got logger!")
	return Logger{log}
}
