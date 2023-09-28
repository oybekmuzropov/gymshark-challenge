package logger

import (
	"github.com/oybekmuzropov/gymshark-challenge/config"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	logger log.Logger
	once   sync.Once
)

func Load() *log.Logger {
	once.Do(func() {
		cfg := config.Load()

		lvl, err := log.ParseLevel(cfg.LogLevel)
		if err != nil {
			lvl = log.InfoLevel
		}
		log.SetLevel(lvl)
	})

	return &logger
}
