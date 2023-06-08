package repository

import (
	"encoding/json"
	"errors"
	"github.com/rmfachran/miniproject2/entity"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type Actor struct {
	db *gorm.DB
}

func NewActor(dbCrud *gorm.DB) Actor {
	return Actor{db: dbCrud}
}
//go:generate mockery --name ActorInterfaceRepo
type ActorInterfaceRepo interface {
	CreateAdmin(actor *entity.Actor) (*entity.Actor, error)
	UpdateAdmin(id uint, actor *entity.Actor) (*entity.Actor, error)
	GetAdmin(id uint) (*entity.Actor, error)
	DeleteAdmin(id uint, actor *entity.Actor) error
	LoginSuperAdmin(username string) (*entity.Actor, error)
	LoginAdmin(username string) (*entity.Actor, error)
	SaveCustomersFromAPI(url string) error
	GetCustomers(first_name, last_name, email string, page, pageSize int) ([]*entity.Customer, error)
	ApprovedAdmin(id uint) ([]*entity.Actor, error)
}

func (repo Actor) CreateAdmin(actor *entity.Actor) (*entity.Actor, error) {
	err := repo.db.Model(&entity.Actor{}).Create(actor).Error
	return actor, err
}

func (repo Actor) UpdateAdmin(id uint, actor *entity.Actor) (*entity.Actor, error) {
	err := repo.db.Model(&entity.Actor{}).Where("id = ?", id).Save(actor).Error
	if err != nil {
		return nil, err
	}
	return actor, nil
}

func (repo Actor) GetAdmin(id uint) (*entity.Actor, error) {
	actor := &entity.Actor{}
	err := repo.db.Model(&entity.Actor{}).Where("id = ?", id).First(actor).Error
	if err != nil {
		return nil, err
	}
	return actor, nil
}

func (repo Actor) DeleteAdmin(id uint, actor *entity.Actor) error {
	err := repo.db.Model(&entity.Actor{}).Where("id = ?", id).Delete(id, *actor).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo Actor) LoginAdmin(username string) (*entity.Actor, error) {
	admin := &entity.Actor{}
	err := repo.db.Model(&entity.Actor{}).Where("username = ? AND is_verified = ? AND is_active = ? AND role_id = ?", username, "true", "true", 1).First(admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo Actor) LoginSuperAdmin(username string) (*entity.Actor, error) {
	admin := &entity.Actor{}
	err := repo.db.Model(&entity.Actor{}).Where("username = ? AND is_verified = ? AND is_active = ? AND role_id = ?", username, "true", "true", 0).First(admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo Actor) GetCustomers(first_name, last_name, email string, page, pageSize int) ([]*entity.Customer, error) {
	customer := []*entity.Customer{}
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

type Get struct {
	Customer []*entity.Customer `json:"data"`
}

func (repo Actor) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	err := repo.db.Model(&entity.Customer{}).Create(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo Actor) GetCustomerByEmail(email string) (*entity.Customer, error) {
	customer := &entity.Customer{}

	err := repo.db.Model(&entity.Customer{}).Where("email = ?", email).First(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo Actor) SaveCustomersFromAPI(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	customerAPIResponse := new(Get)

	err = json.Unmarshal(body, customerAPIResponse)
	if err != nil {
		return err
	}

	for _, customer := range customerAPIResponse.Customer {
		_, err := repo.GetCustomerByEmail(customer.Email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newCustomer := &entity.Customer{
					FirstName: customer.FirstName,
					LastName:  customer.LastName,
					Email:     customer.Email,
					Avatar:    customer.Avatar,
				}
				_, err = repo.CreateCustomer(newCustomer)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			// customer already exist, skip saving
		}
	}

	return nil
}

func (repo Actor) ApprovedAdmin(id uint) ([]*entity.Actor, error) {
	var result []*entity.Actor

	err := repo.db.Model(&entity.Actor{}).
		Select("id, username, role_id, is_verified, created_at, updated_at").
		Where("role_id = ? AND is_verified = ?", 2, "false").
		Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
