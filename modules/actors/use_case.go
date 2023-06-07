package actors

import (
	"crud/entity"
	"crud/repository"
	"time"
)

type UseCaseActor interface {
	CreateAdmin(actor ActorParam) (entity.Actor, error)
	ApproveAdmin(id uint, act ActorParam) (entity.Actor, error)
	LoginSuperAdmin(username string, password string) (*entity.Actor, error)
	LoginAdmin(username string, password string) (*entity.Actor, error)
}

type useCaseActor struct {
	actorRepo repository.ActorInterfaceRepo
}

func (uc useCaseActor) CreateAdmin(actor ActorParam) (entity.Actor, error) {
	var newActor *entity.Actor

	newActor = &entity.Actor{
		Username:   actor.Username,
		Password:   actor.Password,
		RoleId:     2,
		IsVerified: "false",
		IsActive:   "false",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}

	_, err := uc.actorRepo.CreateAdmin(newActor)
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

func (uc useCaseActor) GetAdmin(id uint) (entity.Actor, error) {
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

func (uc useCaseActor) ApproveAdmin(id uint, act ActorParam) (entity.Actor, error) {
	existingData, err := uc.actorRepo.GetAdmin(id)
	if err != nil {
		return entity.Actor{}, err
	}
	existingData.RoleId = 1
	existingData.IsVerified = "true"

	updatedData, err := uc.actorRepo.UpdateAdmin(id, existingData)
	if err != nil {
		return entity.Actor{}, err
	}
	return *updatedData, nil
}

func (uc useCaseActor) LoginAdmin(username string, password string) (*entity.Actor, error) {
	admin, err := uc.actorRepo.LoginAdmin(username)
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
	if admin.Password != password {
		return nil, err
	}
	return admin, nil
}
