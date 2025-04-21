package repositories

import (
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/models"
	pkg_data "yamanmnur/simple-dashboard/pkg/data"
	"yamanmnur/simple-dashboard/pkg/db"
	pkg_requests "yamanmnur/simple-dashboard/pkg/requests"

	"gorm.io/gorm"
)

type ICustomerRepository interface {
	FindById(id uint) (models.Customer, error)
	Detail(id uint, detail_data *models.Customer) error
	Create(customer *models.Customer) error
	Update(customerData *data.CustomerData) (models.Customer, error)
	UpdatePatch(customerData *data.CustomerData) (models.Customer, error)
	Delete(id uint) error
	GetCustomersWithPagination(pageRequest pkg_requests.PageRequest) (pkg_data.PaginateData[data.CustomerData], error)
}

type CustomerRepository struct {
	*db.IDbHandler
}

func (repository *CustomerRepository) Detail(id uint, detail_data *models.Customer) error {
	err := repository.DB.Model(&models.Customer{}).
		Where("customers.id = ?", id).
		Preload("BankAccounts").
		Preload("BankAccounts.TermDeposit").
		Preload("BankAccounts.TermDeposit.TermDepositsType").
		Preload("Pockets").
		First(&detail_data).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *CustomerRepository) FindById(id uint) (models.Customer, error) {
	var user models.Customer
	err := repository.DB.Model(&models.Customer{}).
		Where("customers.id = ?", id).
		Preload("BankAccounts").
		Preload("BankAccounts.TermDeposit").
		Preload("Pockets").
		First(&user).Error

	return user, err
}

func (repository *CustomerRepository) Create(customer *models.Customer) error {

	if result := repository.DB.Create(&customer); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *CustomerRepository) Update(customerData *data.CustomerData) (models.Customer, error) {
	var customer models.Customer

	customer.Name = customerData.Name
	customer.Email = customerData.Email
	customer.PhoneNumber = customerData.PhoneNumber
	customer.Address = customerData.Address
	result := repository.DB.Where("id = ?", customerData.Id).Updates(&customer)
	if result.Error != nil {
		return models.Customer{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.Customer{}, gorm.ErrRecordNotFound
	}

	return customer, nil
}

func (repository *CustomerRepository) UpdatePatch(customerData *data.CustomerData) (models.Customer, error) {
	var customer models.Customer

	if customerData.Id == 0 {
		return models.Customer{}, gorm.ErrRecordNotFound
	}

	if customerData.Name != "" {
		customer.Name = customerData.Name
	}

	if customerData.Email != "" {
		customer.Email = customerData.Email
	}

	if customerData.PhoneNumber != "" {
		customer.PhoneNumber = customerData.PhoneNumber
	}

	if customerData.Address != "" {
		customer.Address = customerData.Address
	}

	if customerData.Photo != "" {
		customer.Photo = customerData.Photo
	}

	result := repository.DB.Where("id = ?", customerData.Id).Updates(&customer)
	if result.Error != nil {
		return models.Customer{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.Customer{}, gorm.ErrRecordNotFound
	}

	return customer, nil
}

func (repository *CustomerRepository) Delete(id uint) error {
	result := repository.DB.Delete(&models.Customer{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (repository *CustomerRepository) GetCustomersWithPagination(pageRequest pkg_requests.PageRequest) (pkg_data.PaginateData[data.CustomerData], error) {
	var page_data pkg_data.PaginateData[data.CustomerData]
	var customers []data.CustomerData

	query := repository.DB.Model(&models.Customer{})

	query = query.Select(`
		customers.id, 
		customers.name, 
		customers.email, 
		customers.phone_number, 
		customers.address, 
		TO_CHAR(customers.created_at, 'DD Mon YYYY HH24:MI:SS') as created_at, 
		b.account_number
	`)
	query = query.Joins(`LEFT JOIN LATERAL (
		select 
			bank_accounts.id, 
			bank_accounts.account_number, 
			bank_accounts.customer_id 
		from bank_accounts 
		where bank_accounts.deleted_at is null
		and bank_accounts.customer_id = customers.id
		order by created_at desc  
		limit 1
		) b ON true
	`)

	if pageRequest.Search != "" {
		searchTerm := "%" + pageRequest.Search + "%"
		query = query.Where(`CONCAT_WS(' ', 
		name, 
		email, 
		phone_number, 
		b.account_number,
		TO_CHAR(customers.created_at, 'DD Mon YYYY HH24:MI:SS'),
		address) LIKE ?`, searchTerm)
	}

	if pageRequest.SortBy != "" && pageRequest.SortDirection != "" {
		sortOrder := pageRequest.SortBy + " " + pageRequest.SortDirection
		query = query.Order(sortOrder)
	}

	if err := query.Count(&page_data.PageData.TotalRows).Error; err != nil {
		return pkg_data.PaginateData[data.CustomerData]{}, err
	}

	if err := query.Limit(int(pageRequest.PageSize)).Offset(int((pageRequest.PageNumber) * pageRequest.PageSize)).Find(&customers).Error; err != nil {
		return pkg_data.PaginateData[data.CustomerData]{}, err
	}

	page_data.Data = customers
	page_data.PageData.Limit = int(pageRequest.PageSize)
	page_data.PageData.Page = int(pageRequest.PageNumber)
	page_data.PageData.TotalPages = int((page_data.PageData.TotalRows + int64(pageRequest.PageSize) - 1) / int64(pageRequest.PageSize)) // Calculate total pages

	return page_data, nil
}
