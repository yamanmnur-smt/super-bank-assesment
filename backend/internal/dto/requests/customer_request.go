package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CustomerRequest struct {
	Photo          string `json:"photo"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone"`
	Address        string `json:"address"`
	Gender         string `json:"gender"`
	AccountPurpose string `json:"account_purpose"`
	SourceOfIncome string `json:"source_of_income"`
	IncomePerMonth string `json:"income_per_month"`
	Jobs           string `json:"jobs"`
	Position       string `json:"position"`
	Industries     string `json:"industries"`
	CompanyName    string `json:"company_name"`
	AddressCompany string `json:"address_company"`
	TotalBalance   string `json:"total_balance"`
	TotalDeposits  string `json:"total_deposits"`
	TotalPockets   string `json:"total_pockets"`
	Username       string `json:"username"`
}

func (r CustomerRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.PhoneNumber, validation.Required),
		validation.Field(&r.Address, validation.Required),
		validation.Field(&r.Username, validation.Required),
	)
}
