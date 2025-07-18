package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Logger
}

func InitLogger(loglvl string) *Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
	})

	log.SetOutput(os.Stdout)

	parsedLvl, err := logrus.ParseLevel(loglvl)
	if err != nil {
		log.Info("error while parsing log lvl")
		parsedLvl = logrus.InfoLevel
	}

	log.SetLevel(parsedLvl)

	return &Logger{log}
}
