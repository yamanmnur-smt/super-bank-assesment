package services_test

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/dto/requests"
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/internal/services"
	pkg_data "yamanmnur/simple-dashboard/pkg/data"
	pkg_requests "yamanmnur/simple-dashboard/pkg/requests"
	"yamanmnur/simple-dashboard/pkg/util"
	mocks_test "yamanmnur/simple-dashboard/tests/services/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var mockCustomerModel = models.Customer{
	Name:           "John Doe",
	Username:       "johndoe",
	Photo:          "photo.jpg",
	Email:          "jhon@mail.com",
	PhoneNumber:    "123456789",
	Address:        "123 Main St",
	Gender:         "male",
	AccountPurpose: "Personal",
	SourceOfIncome: "Salary",
	IncomePerMonth: "5000",
	Jobs:           "Software Engineer",
	Position:       "Senior",
	Industries:     "IT",
	CompanyName:    "Tech Co",
	AddressCompany: "456 Tech St",
	BankAccounts: []models.BankAccount{
		{
			CardNumber:     "0095-2340-2342-2342",
			AccountNumber:  "00002342345",
			Balance:        200000.0,
			AccountType:    "savings",
			Cvc:            "321",
			ExpirationDate: "2025-12-01",
			Status:         models.ACTIVE_BANK_ACCOUNT,
			TermDeposit: []models.TermDeposit{
				{
					Amount:                500000,
					InterestRate:          0.4,
					ExtensionInstructions: models.NO_ROLLOVER,
					StartDate:             "2024-05-02",
					MaturityDate:          "2024-05-02",
					Status:                models.ACTIVE,
					TermDepositsType: models.TermDepositsTypes{
						Name:          "7 days",
						InterestRate:  0.5,
						MinAmount:     500000,
						TermDays:      7,
						MaxAmount:     10000000,
						EffectiveDate: "2024-05-02",
						Status:        models.TERM_ACTIVE,
					},
				},
			},
		},
	},
	Pockets: []models.Pocket{
		{
			Name:     "Savings",
			Balance:  1000.0,
			Currency: "IDR",
		},
		{
			Name:     "Investments",
			Balance:  2000.0,
			Currency: "IDR",
		},
	},
}

func init() {
	mockCustomerModel.ID = 1
}

func TestCustomerService_FindById(t *testing.T) {

	t.Run("FindById() Success", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock Expected CustomerData
		expectedCustomerData := data.CustomerData{
			Id:            mockCustomerModel.ID,
			Name:          mockCustomerModel.Name,
			Email:         mockCustomerModel.Email,
			PhoneNumber:   mockCustomerModel.PhoneNumber,
			Photo:         mockCustomerModel.Photo,
			Address:       mockCustomerModel.Address,
			AccountNumber: mockCustomerModel.BankAccounts[0].AccountNumber,
		}

		// Mocking the repository method
		mockRepo.On("FindById", uint(1)).Return(mockCustomerModel, nil)

		// Call the service method
		result, err := service.FindById(1)

		// Assertions
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != expectedCustomerData {
			t.Errorf("Expected %v, got %v", expectedCustomerData, result)
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedCustomerData, result)

	})

	t.Run("FindById() Customer Not Found", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock Expected CustomerData
		expectedCustomerData := data.CustomerData{}

		// Mocking the repository method
		mockRepo.On("FindById", uint(1)).Return(models.Customer{}, gorm.ErrRecordNotFound)

		// Call the service method
		result, err := service.FindById(1)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, expectedCustomerData, result)
	})

}

func TestCustomerService_Detail(t *testing.T) {
	t.Run("Detail() Success", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock Expected CustomerData
		expectedCustomerData := data.CustomerDetailData{
			Id:             mockCustomerModel.ID,
			Name:           mockCustomerModel.Name,
			Email:          mockCustomerModel.Email,
			PhoneNumber:    mockCustomerModel.PhoneNumber,
			Photo:          mockCustomerModel.Photo,
			Address:        mockCustomerModel.Address,
			Gender:         mockCustomerModel.Gender,
			AccountPurpose: mockCustomerModel.AccountPurpose,
			SourceOfIncome: mockCustomerModel.SourceOfIncome,
			IncomePerMonth: mockCustomerModel.IncomePerMonth,
			Jobs:           mockCustomerModel.Jobs,
			Position:       mockCustomerModel.Position,
			Industries:     mockCustomerModel.Industries,
			CompanyName:    mockCustomerModel.CompanyName,
			AddressCompany: mockCustomerModel.AddressCompany,
			Username:       mockCustomerModel.Username,
			CreatedAt:      time.Now().Format("02 January 2006"),

			Pockets: make([]data.PocketData, len(mockCustomerModel.Pockets)),
		}

		totalBalance := 0.0
		totalDeposits := 0.0
		totalPockets := 0.0
		for i, pocket := range mockCustomerModel.Pockets {
			totalBalance += pocket.Balance
			totalPockets += pocket.Balance
			expectedCustomerData.Pockets[i] = data.PocketData{
				Name:     pocket.Name,
				Balance:  pocket.Balance,
				Currency: pocket.Currency,
			}
		}

		bank := mockCustomerModel.BankAccounts[0]
		bank.AccountNumber = util.GenerateAccountNumber()
		bank.Cvc = util.GenerateCVC()
		bank.CardNumber = util.GenerateCardNumber()
		totalBalance += bank.Balance

		expectedCustomerData.TotalBalance = fmt.Sprintf("%.2f", totalBalance)
		expectedCustomerData.TotalPockets = fmt.Sprintf("%.2f", totalPockets)

		expectedCustomerData.Banks = data.BankAccountData{
			Id:            bank.ID,
			AccountNumber: bank.AccountNumber,
			Balance:       bank.Balance,
			AccountType:   bank.AccountType,
			Cvc:           bank.Cvc,
			CardNumber:    bank.CardNumber,
			Deposites:     make([]data.TermDepositData, len(bank.TermDeposit)),
		}
		for j, deposit := range bank.TermDeposit {
			totalDeposits += deposit.Amount
			expectedCustomerData.Banks.Deposites[j] = data.TermDepositData{
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

		expectedCustomerData.TotalDeposits = fmt.Sprintf("%.2f", totalDeposits)
		mockCustomerModel.CreatedAt = time.Now()
		expectedCustomerData.CreatedAt = mockCustomerModel.CreatedAt.Format("02 January 2006")

		// Mocking the repository method
		mockRepo.On("Detail", uint(1), &models.Customer{}).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.Customer)
			*arg = mockCustomerModel
		}).Return(nil)

		// Call the service method
		result, err := service.Detail(1)
		expectedCustomerData.Banks.AccountNumber = result.Banks.AccountNumber
		expectedCustomerData.Banks.CardNumber = result.Banks.CardNumber
		expectedCustomerData.Banks.ExpirationDate = result.Banks.ExpirationDate
		expectedCustomerData.Banks.Status = result.Banks.Status
		expectedCustomerData.Banks.Cvc = result.Banks.Cvc
		// Assertions

		assert.NoError(t, err)
		assert.Equal(t, expectedCustomerData, result)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, expectedCustomerData.Photo, result.Photo)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Detail() Customer Not Found", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		mockRepo.On("Detail", uint(1), &models.Customer{}).Return(gorm.ErrRecordNotFound)
		result, err := service.Detail(1)
		// Assertions

		assert.Error(t, err)
		assert.Equal(t, data.CustomerDetailData{}, result)
		mockRepo.AssertExpectations(t)
	})

}
func TestCustomerService_Create(t *testing.T) {
	t.Run("Create() Success", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock input and expected output
		inputCustomerData := requests.CustomerRequest{
			Name:           mockCustomerModel.Name,
			Email:          mockCustomerModel.Email,
			PhoneNumber:    mockCustomerModel.PhoneNumber,
			Address:        mockCustomerModel.Address,
			Photo:          mockCustomerModel.Photo,
			Username:       mockCustomerModel.Username,
			Gender:         mockCustomerModel.Gender,
			AccountPurpose: mockCustomerModel.AccountPurpose,
			SourceOfIncome: mockCustomerModel.SourceOfIncome,
			IncomePerMonth: mockCustomerModel.SourceOfIncome,
			Jobs:           mockCustomerModel.Jobs,
			Position:       mockCustomerModel.Position,
			Industries:     mockCustomerModel.Industries,
			CompanyName:    mockCustomerModel.CompanyName,
			AddressCompany: mockCustomerModel.AddressCompany,
		}
		expectedCustomerData := data.CustomerDetailData{
			Id:             1,
			Name:           inputCustomerData.Name,
			Email:          inputCustomerData.Email,
			PhoneNumber:    inputCustomerData.PhoneNumber,
			Address:        inputCustomerData.Address,
			Photo:          inputCustomerData.Photo,
			Username:       inputCustomerData.Username,
			Gender:         inputCustomerData.Gender,
			AccountPurpose: inputCustomerData.AccountPurpose,
			SourceOfIncome: inputCustomerData.SourceOfIncome,
			IncomePerMonth: mockCustomerModel.IncomePerMonth,
			Jobs:           inputCustomerData.Jobs,
			Position:       inputCustomerData.Position,
			Industries:     inputCustomerData.Industries,
			CompanyName:    mockCustomerModel.CompanyName,
			AddressCompany: mockCustomerModel.AddressCompany,
			Pockets:        []data.PocketData{},
			TotalBalance:   fmt.Sprintf("%.2f", mockCustomerModel.BankAccounts[0].Balance),
			TotalDeposits:  "0.00",
			TotalPockets:   "0.00",
			CreatedAt:      time.Now().Format("02 January 2006"),
			Banks: data.BankAccountData{
				AccountNumber:  mockCustomerModel.BankAccounts[0].AccountNumber,
				AccountType:    mockCustomerModel.BankAccounts[0].AccountType,
				CardNumber:     mockCustomerModel.BankAccounts[0].CardNumber,
				Balance:        mockCustomerModel.BankAccounts[0].Balance,
				Cvc:            mockCustomerModel.BankAccounts[0].Cvc,
				ExpirationDate: mockCustomerModel.BankAccounts[0].ExpirationDate,
				Status:         string(mockCustomerModel.BankAccounts[0].Status),
				Deposites:      []data.TermDepositData{},
			},
		}
		mockCustomerModel.Pockets = []models.Pocket{}
		mockCustomerModel.BankAccounts[0].TermDeposit = []models.TermDeposit{}
		mockCustomerModel.CreatedAt = time.Now()

		// Mocking the repository method
		mockRepo.On("Create", mock.AnythingOfType("*models.Customer")).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Customer)
			*arg = mockCustomerModel
		}).Return(nil)

		mockRepo.On("Detail", uint(1), &models.Customer{}).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.Customer)
			*arg = mockCustomerModel
		}).Return(nil)

		// Call the service method
		result, err := service.Create(&inputCustomerData)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedCustomerData, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create() Error When Create", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock input
		inputCustomerData := requests.CustomerRequest{
			Name:           mockCustomerModel.Name,
			Email:          mockCustomerModel.Email,
			PhoneNumber:    mockCustomerModel.PhoneNumber,
			Address:        mockCustomerModel.Address,
			Photo:          mockCustomerModel.Photo,
			Username:       mockCustomerModel.Username,
			Gender:         mockCustomerModel.Gender,
			AccountPurpose: mockCustomerModel.AccountPurpose,
			SourceOfIncome: mockCustomerModel.SourceOfIncome,
			IncomePerMonth: mockCustomerModel.SourceOfIncome,
			Jobs:           mockCustomerModel.Jobs,
			Position:       mockCustomerModel.Position,
			Industries:     mockCustomerModel.Industries,
			CompanyName:    mockCustomerModel.CompanyName,
			AddressCompany: mockCustomerModel.AddressCompany,
		}
		// Mocking the repository method

		mockRepo.On("Create", mock.AnythingOfType("*models.Customer")).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*models.Customer)
			*arg = models.Customer{}
		}).Return(fmt.Errorf("failed to create customer"))

		result, err := service.Create(&inputCustomerData)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, data.CustomerDetailData{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestCustomerService_Update(t *testing.T) {
	t.Run("Update() Success", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock input and expected output
		inputCustomerData := &data.CustomerData{
			Id:          1,
			Name:        "John Updated",
			Email:       "john.updated@mail.com",
			PhoneNumber: "123456789",
			Address:     "123 Updated St",
		}
		expectedCustomerData := data.CustomerData{
			Id:          inputCustomerData.Id,
			Name:        inputCustomerData.Name,
			Email:       inputCustomerData.Email,
			PhoneNumber: inputCustomerData.PhoneNumber,
			Address:     inputCustomerData.Address,
		}

		mockUpdateCustomerModel := models.Customer{
			Name:        inputCustomerData.Name,
			Email:       inputCustomerData.Email,
			PhoneNumber: inputCustomerData.PhoneNumber,
			Address:     inputCustomerData.Address,
		}
		mockUpdateCustomerModel.ID = 1

		// Mocking the repository method
		mockRepo.On("Update", inputCustomerData).Return(mockUpdateCustomerModel, nil)

		// Call the service method
		result, err := service.Update(inputCustomerData)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedCustomerData, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Update() Error", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock input
		inputCustomerData := &data.CustomerData{
			Id:          1,
			Name:        "John Updated",
			Email:       "john.updated@mail.com",
			PhoneNumber: "123456789",
			Address:     "123 Updated St",
		}

		// Mocking the repository method
		mockRepo.On("Update", inputCustomerData).Return(models.Customer{}, fmt.Errorf("failed to update customer"))

		// Call the service method
		result, err := service.Update(inputCustomerData)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, data.CustomerData{}, result)
		mockRepo.AssertExpectations(t)
	})
}
func TestCustomerService_UpdatePhotoCustomer(t *testing.T) {
	t.Run("UpdatePhotoCustomer() Success", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock input and expected output
		inputRequest := &requests.CustomerPhotoRequest{
			Photo: "updated_photo.jpg",
		}
		expectedCustomerData := data.CustomerData{
			Id:    1,
			Name:  "John Doe",
			Email: "jhon@mail.com",
			Photo: inputRequest.Photo,
		}

		mockUpdatedCustomer := models.Customer{
			Name:  "John Doe",
			Email: "jhon@mail.com",
			Photo: inputRequest.Photo,
		}
		mockUpdatedCustomer.ID = 1

		// Mocking the repository method
		mockRepo.On("UpdatePatch", mock.Anything).Return(mockUpdatedCustomer, nil)
		mockRepo.On("FindById", mock.Anything).Return(mockUpdatedCustomer, nil)

		// Call the service method
		result, err := service.UpdatePhotoCustomer(inputRequest)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedCustomerData, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdatePhotoCustomer() Error", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock input
		inputRequest := &requests.CustomerPhotoRequest{
			Photo: "updated_photo.jpg",
		}

		// Mocking the repository method
		mockRepo.On("UpdatePatch", mock.Anything).Return(models.Customer{}, fmt.Errorf("failed to update photo"))

		// Call the service method
		result, err := service.UpdatePhotoCustomer(inputRequest)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, data.CustomerData{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func createMockFileHeader(fieldName, filename string, content []byte) *multipart.FileHeader {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile(fieldName, filename)
	part.Write(content)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.ParseMultipartForm(10 << 20) // 10MB max

	return req.MultipartForm.File[fieldName][0]
}

func TestCustomerService_UploadPhotoCustomer(t *testing.T) {
	t.Run("UploadPhotoCustomer() Success", func(t *testing.T) {
		mockFileContent := []byte("unit test content")
		mockFile := createMockFileHeader("file", "test-customer.txt", mockFileContent)

		mockRepo := new(mocks_test.MockCustomerRepository)
		mockMinio := new(mocks_test.MockMinioClient)

		service := services.CustomerService{Repository: mockRepo, MinioClient: mockMinio}

		// Mock util.Upload
		mockMinio.On("MakeBucket", "customer").Return(nil)
		mockMinio.On("Upload", "customer", "test-customer.txt", []byte("unit test content")).Return(nil)

		// Call the service method
		location, err := service.UploadPhotoCustomer("http://localhost:3001", mockFile)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "http://localhost:3001/api/v1/file/customer/photo?filename=test-customer.txt", location)
	})

	t.Run("UploadPhotoCustomer() File Open Error", func(t *testing.T) {
		mockFile := &multipart.FileHeader{
			Filename: "test-customer.txt",
		}

		mockRepo := new(mocks_test.MockCustomerRepository)
		mockMinio := new(mocks_test.MockMinioClient)
		service := services.CustomerService{Repository: mockRepo, MinioClient: mockMinio}
		mockMinio.On("Upload", "customer", "test-customer.txt", []byte("unit test content")).Return(nil)

		// Call the service method
		location, err := service.UploadPhotoCustomer("http://localhost:3001", mockFile)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "", location)
		assert.EqualError(t, err, "open : The system cannot find the file specified.")
	})

	t.Run("UploadPhotoCustomer() IO Copy Error", func(t *testing.T) {
		mockFileContent := []byte("unit test content")
		mockFile := createMockFileHeader("file", "test-customer.txt", mockFileContent)

		mockRepo := new(mocks_test.MockCustomerRepository)
		mockMinio := new(mocks_test.MockMinioClient)
		service := services.CustomerService{Repository: mockRepo, MinioClient: mockMinio}
		mockMinio.On("MakeBucket", "customer").Return(nil)
		mockMinio.On("Upload", "customer", "test-customer.txt", []byte("unit test content")).Return(fmt.Errorf("failed to copy file content"))

		// Call the service method
		location, err := service.UploadPhotoCustomer("http://localhost:3001", mockFile)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "", location)
		assert.EqualError(t, err, "failed to copy file content")
	})

	t.Run("UploadPhotoCustomer() MakeBucket Error", func(t *testing.T) {
		mockFileContent := []byte("unit test content")
		mockFile := createMockFileHeader("file", "test-customer.txt", mockFileContent)

		mockRepo := new(mocks_test.MockCustomerRepository)
		mockMinio := new(mocks_test.MockMinioClient)
		service := services.CustomerService{Repository: mockRepo, MinioClient: mockMinio}

		// Mock MakeBucket error
		mockMinio.On("MakeBucket", "customer").Return(fmt.Errorf("failed to create bucket"))
		mockMinio.On("Upload", "customer", "test-customer.txt", []byte("unit test content")).Return(nil)

		// Call the service method
		location, err := service.UploadPhotoCustomer("http://localhost:3001", mockFile)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "", location)
		assert.EqualError(t, err, "failed to create bucket")
	})

	t.Run("UploadPhotoCustomer() Upload Error", func(t *testing.T) {
		mockFileContent := []byte("unit test content")
		mockFile := createMockFileHeader("file", "test-customer.txt", mockFileContent)

		mockRepo := new(mocks_test.MockCustomerRepository)
		mockMinio := new(mocks_test.MockMinioClient)
		service := services.CustomerService{Repository: mockRepo, MinioClient: mockMinio}

		// Mock MakeBucket success
		mockMinio.On("MakeBucket", "customer").Return(nil)

		// Mock Upload error
		mockMinio.On("Upload", "customer", "test-customer.txt", mockFileContent).Return(fmt.Errorf("failed to upload file"))

		// Call the service method
		location, err := service.UploadPhotoCustomer("http://localhost:3001", mockFile)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "", location)
		assert.EqualError(t, err, "failed to upload file")
	})
}

func TestCustomerService_GetPhotoCustomer(t *testing.T) {
	t.Run("GetPhotoCustomer() Success", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		mockMinio := new(mocks_test.MockMinioClient)
		service := services.CustomerService{Repository: mockRepo, MinioClient: mockMinio}

		expectedResult := []byte("unit test content")
		mockMinio.On("GetObject", "customer", "test-customer.txt").Return(expectedResult, nil)

		// Call the service method
		data, err := service.GetPhotoCustomer("test-customer.txt")

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, []byte("unit test content"), data)
	})

	t.Run("GetPhotoCustomer() Error", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)

		mockMinio := new(mocks_test.MockMinioClient)
		service := services.CustomerService{Repository: mockRepo, MinioClient: mockMinio}

		// Mock util.Download
		mockMinio.On("GetObject", "customer", "test-customer-not-found.txt").Return([]byte{}, fmt.Errorf("download failed"))

		// Call the service method
		data, err := service.GetPhotoCustomer("test-customer-not-found.txt")

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, []byte{}, data)
	})
}
func TestCustomerService_Delete(t *testing.T) {
	t.Run("Delete() Success", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mocking the repository method
		mockRepo.On("Delete", uint(1)).Return(nil)

		// Call the service method
		err := service.Delete(1)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Delete() Error", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mocking the repository method
		mockRepo.On("Delete", uint(1)).Return(fmt.Errorf("failed to delete customer"))

		// Call the service method
		err := service.Delete(1)

		// Assertions
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to delete customer")
		mockRepo.AssertExpectations(t)
	})
}
func TestCustomerService_GetCustomersWithPagination(t *testing.T) {
	t.Run("GetCustomersWithPagination() Success", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock input and expected output
		pageRequest := pkg_requests.PageRequest{
			Search:        "",
			PageNumber:    0,
			PageSize:      10,
			SortBy:        "name",
			SortDirection: "asc",
		}

		mockPageData := pkg_data.PaginateData[data.CustomerData]{
			Data: []data.CustomerData{
				{
					Id:          1,
					Name:        "John Doe",
					Email:       "johndoe@mail.com",
					PhoneNumber: "123456789",
					Address:     "123 Main St",
				},
				{
					Id:          2,
					Name:        "Jane Doe",
					Email:       "janedoe@mail.com",
					PhoneNumber: "987654321",
					Address:     "456 Main St",
				},
			},
			PageData: pkg_data.PageData{
				Limit:      10,
				Page:       1,
				TotalRows:  2,
				TotalPages: 1,
			},
		}

		// Mocking the repository method
		mockRepo.On("GetCustomersWithPagination", pageRequest).Return(mockPageData, nil)

		// Call the service method
		_, err := service.GetCustomersWithPagination(pageRequest)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetCustomersWithPagination() Error", func(t *testing.T) {
		mockRepo := new(mocks_test.MockCustomerRepository)
		service := services.CustomerService{Repository: mockRepo}

		// Mock input
		pageRequest := pkg_requests.PageRequest{
			Search:        "",
			PageNumber:    0,
			PageSize:      10,
			SortBy:        "name",
			SortDirection: "asc",
		}

		// Mocking the repository method
		mockRepo.On("GetCustomersWithPagination", pageRequest).Return(pkg_data.PaginateData[data.CustomerData]{}, fmt.Errorf("failed to fetch customers"))

		// Call the service method
		result, err := service.GetCustomersWithPagination(pageRequest)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "failed to fetch customers")
		mockRepo.AssertExpectations(t)
	})
}
