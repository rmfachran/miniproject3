package actors

import (
	"github.com/rmfachran/miniproject2/entity"
	"github.com/rmfachran/miniproject2/middleware"
	"github.com/rmfachran/miniproject2/repository"
	"time"
)

type UseCaseActor interface {
	CreateAdmin(actor ActorParam) (entity.Actor, error)
	ApprovedAdmin(id uint) ([]*entity.Actor, error)
	LoginSuperAdmin(username string, password string) (*entity.Actor, error)
	GetAdminById(id uint) (entity.Actor, error)
	LoginAdmin(username string, password string) (*entity.Actor, error)
	GetCustomers(first_name, last_name, email string, page, pageSize int) (*entity.Customer, error)
	SaveCustomersFromAPI() error
}

type useCaseActor struct {
	actorRepo repository.ActorInterfaceRepo
}

func (uc useCaseActor) CreateAdmin(actor ActorParam) (entity.Actor, error) {
	var newActor *entity.Actor

	pass, err := middleware.HashPassword(actor.Password)
	if err != nil {
		return entity.Actor{}, err
	}
	newActor = &entity.Actor{
		Username:   actor.Username,
		Password:   pass,
		RoleId:     2,
		IsVerified: "false",
		IsActive:   "false",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}

	_, err = uc.actorRepo.CreateAdmin(newActor)
	if err != nil {
		return *newActor, err
	}
	return *newActor, nil
}

//func (uc useCaseActor) UpdateActorById(id uint, act ActorParam) (entity.Actor, error) {
//	existingData, err := uc.actorRepo.UpdateAdmin(id)
//	if err != nil {
//		return entity.Actor{}, err
//	}
//	existingData.Username = existingData
//}

func (uc useCaseActor) GetAdminById(id uint) (entity.Actor, error) {
	admin, err := uc.actorRepo.GetAdmin(id)
	if err != nil {
		return entity.Actor{}, err
	}
	return *admin, nil
}

func (uc useCaseActor) UpdateAdmin(id uint, adm ActorParam) (entity.Actor, error) {
	existingData, err := uc.actorRepo.GetAdmin(id)
	if err != nil {
		return entity.Actor{}, err
	}
	existingData.Username = adm.Username
	existingData.Password = adm.Password

	updatedData, err := uc.actorRepo.UpdateAdmin(id, existingData)
	if err != nil {
		return entity.Actor{}, err
	}
	return *updatedData, nil
}

func (uc useCaseActor) ApprovedAdmin(id uint) ([]*entity.Actor, error) {
	request, err := uc.actorRepo.ApprovedAdmin(id)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (uc useCaseActor) LoginAdmin(username string, password string) (*entity.Actor, error) {
	admin, err := uc.actorRepo.LoginAdmin(username)
	if err != nil {
		return nil, err
	}
	//pass, err := middleware.HashPassword(password)
	if err != nil {
		return nil, err
	}
	if admin.Password != password {
		return nil, err
	}

	return admin, nil
}

func (uc useCaseActor) LoginSuperAdmin(username string, password string) (*entity.Actor, error) {
	admin, err := uc.actorRepo.LoginSuperAdmin(username)
	if err != nil {
		return nil, err
	}
	//pass, err := middleware.HashPassword(password)
	if admin.Password != password {
		return nil, err
	}
	return admin, nil
}

func (uc useCaseActor) GetCustomers(first_name, last_name, email string, page, pageSize int) ([]*entity.Customer, error) {
	customers, err := uc.actorRepo.GetCustomers(first_name, last_name, email, page, pageSize)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (uc useCaseActor) SaveCustomersFromAPI() error {
	url := "https://reqres.in/api/users?page=2"

	err := uc.actorRepo.SaveCustomersFromAPI(url)
	if err != nil {
		return err
	}

	return nil
}
