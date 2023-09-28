package config

import (
	"github.com/alexflint/go-arg"
	"sync"
)

var (
	config Config
	once   sync.Once
)

type Config struct {
	LogLevel string `arg:"env:LOG_LEVEL"`
	HTTPAddr string `arg:"env:HTTP_ADDR"`
}

func Load() *Config {
	once.Do(func() {
		arg.Parse(&config)
	})

	return &config
}
