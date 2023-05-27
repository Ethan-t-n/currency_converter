package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCurrency(t *testing.T) {
	rates := map[string]float64{
		"EUR": 0.85,
		"GBP": 0.72,
		"CAD": 1.21,
	}

	amount := 100.0
	targetCurrency := "EUR"
	expectedAmount := 85.0

	convertedAmount := ConvertCurrency(amount, targetCurrency, rates)

	assert.Equal(t, expectedAmount, convertedAmount)
}
