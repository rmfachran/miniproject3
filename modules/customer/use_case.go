package customer

import (
	"crud/entity"
	"crud/repository"
)

type UseCaseCustomer interface {
	GetCustomerById(id uint) (entity.Customer, error)
	CreateCustomer(customer CustomerParam) (entity.Customer, error)
	UpdateCustomerById(id uint, cust CustomerParam) (entity.Customer, error)
	DeleteCustomerById(id uint) error
}

type useCaseCustomer struct {
	customerRepo repository.CustomerInterfaceRepo
}

func (uc useCaseCustomer) CreateCustomer(customer CustomerParam) (entity.Customer, error) {
	var newCustomer *entity.Customer

	newCustomer = &entity.Customer{
		Email:     customer.Email,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Avatar:    customer.Avatar,
	}

	_, err := uc.customerRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}
	return *newCustomer, nil
}

func (uc useCaseCustomer) UpdateCustomerById(id uint, cust CustomerParam) (entity.Customer, error) {
	// Get existing customer by id
	existingData, err := uc.customerRepo.GetCustomerById(id)
	if err != nil {
		return entity.Customer{}, err
	}

	existingData.Email = cust.Email
	existingData.FirstName = cust.FirstName
	existingData.LastName = cust.LastName
	existingData.Avatar = cust.Avatar

	// Updated data
	updatedData, err := uc.customerRepo.UpdateCustomerById(id, existingData)
	if err != nil {
		return entity.Customer{}, err
	}

	return *updatedData, nil

}

func (uc useCaseCustomer) DeleteCustomerById(id uint) error {

	deleteData, err := uc.customerRepo.GetCustomerById(id)
	if err != nil {
		return err
	}

	return uc.customerRepo.DeleteById(id, deleteData)
}

func (uc useCaseCustomer) GetCustomerById(id uint) (entity.Customer, error) {
	//var customer entity.Customer
	customer, err := uc.customerRepo.GetCustomerById(id)
	if err != nil {
		return entity.Customer{}, err
	}
	return *customer, err
}
