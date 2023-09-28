package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/oybekmuzropov/gymshark-challenge/config"
	"github.com/oybekmuzropov/gymshark-challenge/handler"
	"github.com/oybekmuzropov/gymshark-challenge/logger"
	"github.com/oybekmuzropov/gymshark-challenge/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.Load()
	l := logger.Load()
	fulFillOrderService := service.NewFulfillOrderService()

	fulFillOrderHandler := handler.NewFulFillOrderHandler(cfg, l, fulFillOrderService)

	e := echo.New()
	if cfg.LogLevel == "debug" {
		if l, ok := e.Logger.(*echoLog.Logger); ok {
			l.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
		}
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")

	fulFillRoutes := v1.Group("/fulfill-order")
	fulFillRoutes.POST("/calculate", fulFillOrderHandler.CalculatePacksCanBeSent)

	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
