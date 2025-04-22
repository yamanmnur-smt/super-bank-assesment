package data

import "yamanmnur/simple-dashboard/internal/models"

type TermDepositsTypesData struct {
	Name          string  `json:"name"`
	InterestRate  float64 `json:"interest_rate"`
	MinAmount     float64 `json:"min_amount"`
	MaxAmount     float64 `json:"max_amount"`
	TermDays      uint    `json:"term_days"`
	EffectiveDate string  `json:"effective_date"`
}

type TermDepositData struct {
	Amount                string                          `json:"amount"`
	InterestRate          float64                         `json:"interest_rate"`
	StartDate             string                          `json:"start_date"`
	MaturityDate          string                          `json:"maturity_date"`
	Status                models.TermStatus               `json:"status"`
	ExtensionInstructions models.TermExtensionInstruction `json:"extension_instructions"`
	TermDepositsTypes     TermDepositsTypesData           `json:"term_deposits_types"`
}
