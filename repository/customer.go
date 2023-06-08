package repository

import (
	"github.com/rmfachran/miniproject2/entity"
	"gorm.io/gorm"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{db: dbCrud}
}

//go:generate mockery --name CustomerInterfaceRepo
type CustomerInterfaceRepo interface {
	GetCustomerById(id uint) (*entity.Customer, error)
	CreateCustomer(customer *entity.Customer) (*entity.Customer, error)
	UpdateCustomerById(id uint, customer *entity.Customer) (*entity.Customer, error)
	DeleteById(id uint, customer *entity.Customer) error
	GetCustomerByEmail(email string) (*entity.Customer, error)
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

func (repo Customer) GetCustomerByEmail(email string) (*entity.Customer, error) {
	customer := &entity.Customer{}
	err := repo.db.Model(&entity.Customer{}).Where("email = ?", email).First(customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (repo Customer) GetCustomer(first_name, last_name, email string, page, pageSize int) (*entity.Customer, error) {
	customer := &entity.Customer{}
	query := repo.db.Model(&entity.Customer{})
	if first_name != "" {
		query = query.Where("first_name = ?", "%"+first_name+"%")
	} else if last_name != "" {
		query = query.Where("last_name = ?", "%"+last_name+"%")
	} else if email != "" {
		query = query.Where("email = ?", "%"+email+"%")
	}
	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}
