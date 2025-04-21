package util_test

import (
	"regexp"
	"testing"
	"yamanmnur/simple-dashboard/pkg/util"
)

func TestGenerateCVC(t *testing.T) {
	cvc := util.GenerateCVC()
	if len(cvc) != 3 {
		t.Errorf("Expected CVC length to be 3, got %d", len(cvc))
	}
	match, _ := regexp.MatchString(`^\d{3}$`, cvc)
	if !match {
		t.Errorf("CVC format is invalid: %s", cvc)
	}
}

func TestGenerateAccountNumber(t *testing.T) {
	accountNumber := util.GenerateAccountNumber()
	if len(accountNumber) != 11 {
		t.Errorf("Expected account number length to be 11, got %d", len(accountNumber))
	}
	match, _ := regexp.MatchString(`^\d{11}$`, accountNumber)
	if !match {
		t.Errorf("Account number format is invalid: %s", accountNumber)
	}
}

func TestGenerateExpirationDate(t *testing.T) {
	expirationDate := util.GenerateExpirationDate()
	match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, expirationDate)
	if !match {
		t.Errorf("Expiration date format is invalid: %s", expirationDate)
	}
}

func TestGenerateCardNumber(t *testing.T) {
	cardNumber := util.GenerateCardNumber()
	if len(cardNumber) != 19 {
		t.Errorf("Expected card number length to be 19, got %d", len(cardNumber))
	}
	match, _ := regexp.MatchString(`^\d{4}-\d{4}-\d{4}-\d{4}$`, cardNumber)
	if !match {
		t.Errorf("Card number format is invalid: %s", cardNumber)
	}
}
