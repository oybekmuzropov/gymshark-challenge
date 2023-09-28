package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/oybekmuzropov/gymshark-challenge/config"
	"github.com/oybekmuzropov/gymshark-challenge/model"
	"github.com/oybekmuzropov/gymshark-challenge/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type FulFillOrderHandler struct {
	log     *log.Logger
	service service.IFulfillOrderService
	config  *config.Config
}

func NewFulFillOrderHandler(cfg *config.Config, log *log.Logger, service service.IFulfillOrderService) *FulFillOrderHandler {
	return &FulFillOrderHandler{
		config:  cfg,
		log:     log,
		service: service,
	}
}

func (h *FulFillOrderHandler) CalculatePacksCanBeSent(ctx echo.Context) error {
	h.log.Infof("Received request %s %s", ctx.Request().URL, ctx.Request().Method)

	var req model.CalculatePackReq
	err := ctx.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not decode request data")
	}

	res, err := h.service.CalculatePacks(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	h.log.Debugf("Calculated Packs Can Be Sent '%v'", res)

	return ctx.JSON(http.StatusOK, res)
}
