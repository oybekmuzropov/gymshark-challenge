package util

import (
	"github.com/oybekmuzropov/gymshark-challenge/model"
)

func CalculateMinPacksCanBeSent(orderedItemsCount int) []*model.Pack {
	minItemsCountCanBeSent := calculateMinItemsCountCanBeSent(orderedItemsCount)

	var packs []*model.Pack

	i := len(model.PackSizes) - 1
	packSizesMap := make(map[int]int)

	for i >= 0 {
		if minItemsCountCanBeSent >= model.PackSizes[i] {
			packSizesMap[model.PackSizes[i]]++
			minItemsCountCanBeSent -= model.PackSizes[i]
		} else {
			i--
		}
	}

	for k, v := range packSizesMap {
		packs = append(packs, &model.Pack{
			PackSize: k,
			Count:    v,
		})
	}

	return packs
}

func calculateMinItemsCountCanBeSent(orderedItemsCount int) int {
	var minItemsCountCanBeSent int

	i := len(model.PackSizes) - 1

	for i >= 0 {
		if orderedItemsCount >= model.PackSizes[i] {
			orderedItemsCount -= model.PackSizes[i]
			minItemsCountCanBeSent += model.PackSizes[i]
		} else {
			i--
		}
	}

	if orderedItemsCount > 0 {
		minItemsCountCanBeSent += model.PackSizes[0]
	}

	return minItemsCountCanBeSent
}
