package data

import "yamanmnur/simple-dashboard/internal/models"

type CustomerData struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Photo         string `json:"photo"`
	Address       string `json:"address"`
	AccountNumber string `json:"account_number"`
	CreatedAt     string `json:"created_at"`
}

type CustomerMessageData struct {
	Type         string          `json:"id"`
	ContentData  CustomerData    `json:"content_data"`
	ContentModel models.Customer `json:"content_model"`
}

type CustomerDetailData struct {
	Id             uint   `json:"id"`
	Photo          string `json:"photo"`
	Username       string `json:"username"`
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
	CreatedAt      string `json:"created_at"`

	Banks   BankAccountData `json:"banks"`
	Pockets []PocketData    `json:"pockets"`
}

type BankAccountData struct {
	Id             uint              `json:"id"`
	AccountNumber  string            `json:"account_number"`
	Balance        float64           `json:"balance"`
	AccountType    string            `json:"account_type"`
	CardNumber     string            `json:"card_number"`
	ExpirationDate string            `json:"expiration_date"`
	Cvc            string            `json:"cvc"`
	Status         string            `json:"status"`
	Deposites      []TermDepositData `json:"deposits"`
}

type PocketData struct {
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}
