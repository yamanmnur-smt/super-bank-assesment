package util_test

import (
	"testing"
	"yamanmnur/simple-dashboard/pkg/util"

	"github.com/stretchr/testify/assert"
)

func TestFormatIDRCurrency_ValidInput(t *testing.T) {
	result, err := util.FormatIDRCurrency("1234567.89")
	assert.NoError(t, err)
	assert.Equal(t, "Rp1.234.567", result)
}

func TestFormatIDRCurrency_ZeroInput(t *testing.T) {
	result, err := util.FormatIDRCurrency("0")
	assert.NoError(t, err)
	assert.Equal(t, "Rp0", result)
}

func TestFormatIDRCurrency_InvalidInput(t *testing.T) {
	result, err := util.FormatIDRCurrency("invalid")
	assert.Error(t, err)
	assert.Equal(t, "", result)
}

func TestFormatIDRCurrency_NegativeInput(t *testing.T) {
	result, err := util.FormatIDRCurrency("-1234567.89")
	assert.NoError(t, err)
	assert.Equal(t, "Rp-1.234.567", result)
}

func TestFormatIDRCurrency_SmallDecimalInput(t *testing.T) {
	result, err := util.FormatIDRCurrency("123.45")
	assert.NoError(t, err)
	assert.Equal(t, "Rp123", result)
}

func TestFormatIDRCurrency_LargeInput(t *testing.T) {
	result, err := util.FormatIDRCurrency("1234567890123.45")
	assert.NoError(t, err)
	assert.Equal(t, "Rp1.234.567.890.123", result)
}
