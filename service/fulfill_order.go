package service

import (
	"github.com/oybekmuzropov/gymshark-challenge/config"
	"github.com/oybekmuzropov/gymshark-challenge/model"
	"github.com/oybekmuzropov/gymshark-challenge/util"
	"github.com/oybekmuzropov/gymshark-challenge/util/validator"
	"sort"
)

type IFulfillOrderService interface {
	CalculatePacks(req *model.CalculatePackReq) ([]*model.Pack, error)
}

type fulfillOrderService struct {
	config    *config.Config
	validator *validator.Validator
}

func NewFulfillOrderService() IFulfillOrderService {
	return &fulfillOrderService{
		config:    config.Load(),
		validator: validator.NewValidator(),
	}
}

func (s *fulfillOrderService) CalculatePacks(req *model.CalculatePackReq) ([]*model.Pack, error) {
	err := s.validator.Validate(req)
	if err != nil {
		return nil, err
	}

	packs := util.CalculateMinPacksCanBeSent(req.OrderedItemsCount)

	sort.Slice(packs, func(i, j int) bool {
		return packs[i].PackSize > packs[j].PackSize
	})

	return packs, nil
}
