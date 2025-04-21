package services

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/dto/requests"
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/internal/repositories"
	pkg_data "yamanmnur/simple-dashboard/pkg/data"
	pkg_requests "yamanmnur/simple-dashboard/pkg/requests"
	"yamanmnur/simple-dashboard/pkg/util"
)

type ICustomerService interface {
	FindById(id uint) (data.CustomerData, error)
	Detail(id uint) (data.CustomerDetailData, error)
	Create(request *requests.CustomerRequest) (data.CustomerDetailData, error)
	Update(userData *data.CustomerData) (data.CustomerData, error)
	UpdatePhotoCustomer(userData *requests.CustomerPhotoRequest) (data.CustomerData, error)
	UploadPhotoCustomer(baseUrl string, file *multipart.FileHeader) (string, error)
	GetPhotoCustomer(filename string) ([]byte, error)
	Delete(id uint) error
	GetCustomersWithPagination(pageRequest pkg_requests.PageRequest) (*pkg_data.PaginateResponse[data.CustomerData], error)
}

type CustomerService struct {
	Repository  repositories.ICustomerRepository
	MinioClient util.IMinioClient
}

func (service *CustomerService) Detail(id uint) (data.CustomerDetailData, error) {
	var customer models.Customer
	var result data.CustomerDetailData

	err := service.Repository.Detail(id, &customer)
	if err != nil {
		return data.CustomerDetailData{}, err
	}

	result.Id = customer.ID
	result.Username = customer.Username
	result.Name = customer.Name
	result.Email = customer.Email
	result.PhoneNumber = customer.PhoneNumber
	result.Address = customer.Address
	result.AccountPurpose = customer.AccountPurpose
	result.SourceOfIncome = customer.SourceOfIncome
	result.IncomePerMonth = customer.IncomePerMonth
	result.Jobs = customer.Jobs
	result.Photo = customer.Photo
	result.Gender = customer.Gender
	result.Position = customer.Position
	result.Industries = customer.Industries
	result.CompanyName = customer.CompanyName
	result.AddressCompany = customer.AddressCompany
	result.Pockets = make([]data.PocketData, len(customer.Pockets))
	result.CreatedAt = customer.CreatedAt.Format("02 January 2006")

	totalBalance := 0.0
	totalDeposits := 0.0
	totalPockets := 0.0
	for i, pocket := range customer.Pockets {
		totalBalance += pocket.Balance
		totalPockets += pocket.Balance
		result.Pockets[i] = data.PocketData{
			Name:     pocket.Name,
			Balance:  pocket.Balance,
			Currency: pocket.Currency,
		}
	}

	bank := customer.BankAccounts[0]
	totalBalance += bank.Balance
	result.TotalBalance = fmt.Sprintf("%.2f", totalBalance)
	result.TotalPockets = fmt.Sprintf("%.2f", totalPockets)

	formatTotalPockets, _ := util.FormatIDRCurrency(result.TotalPockets)
	result.TotalPockets = formatTotalPockets

	formatTotalBalance, _ := util.FormatIDRCurrency(result.TotalBalance)
	result.TotalBalance = formatTotalBalance

	result.Banks = data.BankAccountData{
		Id:             bank.ID,
		AccountNumber:  bank.AccountNumber,
		Balance:        bank.Balance,
		AccountType:    bank.AccountType,
		Cvc:            bank.Cvc,
		CardNumber:     bank.CardNumber,
		Status:         string(bank.Status),
		ExpirationDate: bank.ExpirationDate,
		Deposites:      make([]data.TermDepositData, len(bank.TermDeposit)),
	}
	for j, deposit := range bank.TermDeposit {
		totalDeposits += deposit.Amount
		result.Banks.Deposites[j] = data.TermDepositData{
			Amount:                deposit.Amount,
			InterestRate:          deposit.InterestRate,
			StartDate:             deposit.StartDate,
			MaturityDate:          deposit.MaturityDate,
			Status:                deposit.Status,
			ExtensionInstructions: deposit.ExtensionInstructions,
			TermDepositsTypes: data.TermDepositsTypesData{
				Name:          deposit.TermDepositsType.Name,
				InterestRate:  deposit.TermDepositsType.InterestRate,
				MinAmount:     deposit.TermDepositsType.MinAmount,
				MaxAmount:     deposit.TermDepositsType.MaxAmount,
				TermDays:      deposit.TermDepositsType.TermDays,
				EffectiveDate: deposit.TermDepositsType.EffectiveDate,
			},
		}
	}

	result.TotalDeposits = fmt.Sprintf("%.2f", totalDeposits)
	formatTotalDeposits, _ := util.FormatIDRCurrency(result.TotalDeposits)
	result.TotalDeposits = formatTotalDeposits

	return result, nil
}

func (service *CustomerService) FindById(id uint) (data.CustomerData, error) {
	var result data.CustomerData

	customer, err := service.Repository.FindById(id)
	if err != nil {
		return data.CustomerData{}, err
	}

	result.Id = customer.ID
	result.Name = customer.Name
	result.Email = customer.Email
	result.PhoneNumber = customer.PhoneNumber
	result.Address = customer.Address
	result.Photo = customer.Photo
	result.AccountNumber = customer.BankAccounts[0].AccountNumber

	return result, nil
}

func (service *CustomerService) Create(request *requests.CustomerRequest) (data.CustomerDetailData, error) {
	inputCustomer := models.Customer{
		Name:           request.Name,
		Email:          request.Email,
		PhoneNumber:    request.PhoneNumber,
		Address:        request.Address,
		Photo:          request.Photo,
		Username:       request.Username,
		Gender:         request.Gender,
		AccountPurpose: request.AccountPurpose,
		SourceOfIncome: request.SourceOfIncome,
		IncomePerMonth: request.SourceOfIncome,
		Jobs:           request.Jobs,
		Position:       request.Position,
		Industries:     request.Industries,
		CompanyName:    request.CompanyName,
		AddressCompany: request.AddressCompany,
		BankAccounts: []models.BankAccount{
			{
				CardNumber:     util.GenerateCardNumber(),
				AccountNumber:  util.GenerateAccountNumber(),
				Balance:        0,
				Cvc:            util.GenerateCVC(),
				ExpirationDate: util.GenerateExpirationDate(),
				AccountType:    "Savings",
				Status:         models.ACTIVE_BANK_ACCOUNT,
			},
		},
	}

	err := service.Repository.Create(&inputCustomer)

	if err != nil {
		return data.CustomerDetailData{}, err
	}

	customerDetail, _ := service.Detail(inputCustomer.ID)
	return customerDetail, nil
}

func (service *CustomerService) Update(request *data.CustomerData) (data.CustomerData, error) {
	var result data.CustomerData
	customer, err := service.Repository.Update(request)
	if err != nil {
		return data.CustomerData{}, err
	}

	result.Id = customer.ID
	result.Name = customer.Name
	result.Email = customer.Email
	result.PhoneNumber = customer.PhoneNumber
	result.Address = customer.Address

	return result, nil
}

func (service *CustomerService) UpdatePhotoCustomer(request *requests.CustomerPhotoRequest) (data.CustomerData, error) {
	var result data.CustomerData
	result.Photo = request.Photo
	id, _ := strconv.ParseUint(request.Id, 10, 32)

	result.Id = uint(id)

	_, err := service.Repository.UpdatePatch(&result)
	if err != nil {
		return data.CustomerData{}, err
	}

	customer, _ := service.Repository.FindById(result.Id)

	result.Id = customer.ID
	result.Name = customer.Name
	result.Email = customer.Email
	result.PhoneNumber = customer.PhoneNumber
	result.Address = customer.Address

	return result, nil
}

func (service *CustomerService) UploadPhotoCustomer(baseUrl string, file *multipart.FileHeader) (string, error) {
	multipartFile, err := file.Open()

	if err != nil {
		return "", err
	}
	defer multipartFile.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, multipartFile); err != nil {
		return "", err
	}

	err = service.MinioClient.MakeBucket("customer")
	if err != nil {
		return "", err
	}
	err = service.MinioClient.Upload("customer", file.Filename, buf.Bytes())

	if err != nil {
		return "", err
	}

	location_url := fmt.Sprintf("%s/api/v1/file/customer/photo?filename=%s", baseUrl, file.Filename)

	return location_url, nil
}

func (service *CustomerService) GetPhotoCustomer(filename string) ([]byte, error) {
	bytes, err := service.MinioClient.GetObject("customer", filename)

	return bytes, err
}

func (service *CustomerService) Delete(id uint) error {
	err := service.Repository.Delete(id)
	if err != nil {

		return err
	}

	return nil
}
func (service *CustomerService) GetCustomersWithPagination(pageRequest pkg_requests.PageRequest) (*pkg_data.PaginateResponse[data.CustomerData], error) {

	pageData, err := service.Repository.GetCustomersWithPagination(pageRequest)
	if err != nil {
		return nil, err
	}

	response := &pkg_data.PaginateResponse[data.CustomerData]{
		Data: pageData.Data,
		PageData: pkg_data.PageData{
			Limit:      pageData.PageData.Limit,
			Page:       pageData.PageData.Page,
			Sort:       pageRequest.SortBy + " " + pageRequest.SortDirection,
			TotalRows:  pageData.PageData.TotalRows,
			TotalPages: pageData.PageData.TotalPages,
		},
	}

	return response, nil
}
