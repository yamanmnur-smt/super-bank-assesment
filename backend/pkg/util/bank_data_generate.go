package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateCVC() string {
	return fmt.Sprintf("%03d", rand.Intn(1000)) // 000 - 999
}

func GenerateAccountNumber() string {
	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(10000000000)
	return fmt.Sprintf("%011d", number)
}

func GenerateExpirationDate() string {
	now := time.Now()
	exp := time.Date(now.Year()+3, now.Month(), 1, 0, 0, 0, 0, time.UTC)
	return exp.Format("2006-01-02")
}

func GenerateCardNumber() string {
	var cardNumber string
	for i := 0; i < 4; i++ {
		part := fmt.Sprintf("%04d", rand.Intn(10000)) // 0000 - 9999
		if i > 0 {
			cardNumber += "-"
		}
		cardNumber += part
	}
	return cardNumber
}
