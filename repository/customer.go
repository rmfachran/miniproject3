package repository

import (
	"crud/entity"
	"gorm.io/gorm"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{db: dbCrud}
}

type CustomerInterfaceRepo interface {
	GetCustomerById(id uint) (*entity.Customer, error)
	CreateCustomer(customer *entity.Customer) (*entity.Customer, error)
	UpdateCustomerById(id uint, customer *entity.Customer) (*entity.Customer, error)
	DeleteById(id uint, customer *entity.Customer) error
}

func (repo Customer) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	err := repo.db.Model(&entity.Customer{}).Create(customer).Error
	return customer, err
}

func (repo Customer) UpdateCustomerById(id uint, customer *entity.Customer) (*entity.Customer, error) {
	err := repo.db.Model(&entity.Customer{}).Where("id = ?", id).Save(customer).Error

	if err != nil {
		return nil, err
	}
	return customer, err
}

func (repo Customer) DeleteById(id uint, customer *entity.Customer) error {
	//customer := &entity.Customer{}
	err := repo.db.Model(&entity.Customer{}).Where("id = ?", id).Delete(id, *customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo Customer) GetCustomerById(id uint) (*entity.Customer, error) {
	customer := &entity.Customer{}
	err := repo.db.Model(&entity.Customer{}).Where("id = ?", id).First(customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}
