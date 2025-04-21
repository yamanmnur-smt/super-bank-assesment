package util

import (
	"strconv"
	"strings"
)

func FormatIDRCurrency(balance string) (string, error) {
	amount, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		return "", err
	}

	intAmount := int64(amount)

	result := formatWithDotSeparator(intAmount)

	return "Rp" + result, nil
}

func formatWithDotSeparator(n int64) string {
	s := strconv.FormatInt(n, 10)
	var result strings.Builder

	length := len(s)
	for i, digit := range s {
		if i != 0 && (length-i)%3 == 0 {
			result.WriteRune('.')
		}
		result.WriteRune(digit)
	}
	return result.String()
}
