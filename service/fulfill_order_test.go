package service

import (
	"github.com/oybekmuzropov/gymshark-challenge/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFulfillOrderService_CalculatePacks_ErrorNotEmpty(t *testing.T) {
	var expected []*model.Pack

	s := NewFulfillOrderService()

	t.Run("Check Input Not Empty", func(t *testing.T) {
		res, err := s.CalculatePacks(&model.CalculatePackReq{})
		assert.Equal(t, expected, res)
		assert.NotNil(t, err)
	})
}

func TestFulfillOrderService_CalculatePacks_ErrorGreaterThanZero(t *testing.T) {
	var expected []*model.Pack

	s := NewFulfillOrderService()

	t.Run("Check Input Greater Than Zero", func(t *testing.T) {
		res, err := s.CalculatePacks(&model.CalculatePackReq{OrderedItemsCount: 0})
		assert.Equal(t, expected, res)
		assert.NotNil(t, err)
	})
}

func TestFulfillOrderService_CalculatePacks_CheckWithValues(t *testing.T) {
	s := NewFulfillOrderService()

	var testCases = []struct {
		testName string
		input    *model.CalculatePackReq
		expected []*model.Pack
	}{
		{
			testName: "Value is 1",
			input:    &model.CalculatePackReq{OrderedItemsCount: 1},
			expected: []*model.Pack{
				{
					PackSize: 250,
					Count:    1,
				},
			},
		},
		{
			testName: "Value is 250",
			input:    &model.CalculatePackReq{OrderedItemsCount: 250},
			expected: []*model.Pack{
				{
					PackSize: 250,
					Count:    1,
				},
			},
		},
		{
			testName: "Value is 251",
			input:    &model.CalculatePackReq{OrderedItemsCount: 251},
			expected: []*model.Pack{
				{
					PackSize: 500,
					Count:    1,
				},
			},
		},
		{
			testName: "Value is 500",
			input:    &model.CalculatePackReq{OrderedItemsCount: 500},
			expected: []*model.Pack{
				{
					PackSize: 500,
					Count:    1,
				},
			},
		},
		{
			testName: "Value is 501",
			input:    &model.CalculatePackReq{OrderedItemsCount: 501},
			expected: []*model.Pack{
				{
					PackSize: 500,
					Count:    1,
				},
				{
					PackSize: 250,
					Count:    1,
				},
			},
		},
		{
			testName: "Value is 12001",
			input:    &model.CalculatePackReq{OrderedItemsCount: 12001},
			expected: []*model.Pack{
				{
					PackSize: 5000,
					Count:    2,
				},
				{
					PackSize: 2000,
					Count:    1,
				},
				{
					PackSize: 250,
					Count:    1,
				},
			},
		},
		{
			testName: "Value is 15000",
			input:    &model.CalculatePackReq{OrderedItemsCount: 15000},
			expected: []*model.Pack{
				{
					PackSize: 5000,
					Count:    3,
				},
			},
		},
		{
			testName: "Value is 7250",
			input:    &model.CalculatePackReq{OrderedItemsCount: 7250},
			expected: []*model.Pack{
				{
					PackSize: 5000,
					Count:    1,
				},
				{
					PackSize: 2000,
					Count:    1,
				},
				{
					PackSize: 250,
					Count:    1,
				},
			},
		},
		{
			testName: "Value is 3649",
			input:    &model.CalculatePackReq{OrderedItemsCount: 3649},
			expected: []*model.Pack{
				{
					PackSize: 2000,
					Count:    1,
				},
				{
					PackSize: 1000,
					Count:    1,
				},
				{
					PackSize: 500,
					Count:    1,
				},
				{
					PackSize: 250,
					Count:    1,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			res, err := s.CalculatePacks(tc.input)

			assert.Equal(t, tc.expected, res)
			assert.Nil(t, err)
		})
	}
}

// {
//	s := NewFulfillOrderService()
//

//}
