package repository

import (
	"crud/entity"
	"gorm.io/gorm"
)

type Actor struct {
	db *gorm.DB
}

func NewActor(dbCrud *gorm.DB) Actor {
	return Actor{db: dbCrud}
}

type ActorInterfaceRepo interface {
	CreateAdmin(actor *entity.Actor) (*entity.Actor, error)
	UpdateAdmin(id uint, actor *entity.Actor) (*entity.Actor, error)
	GetAdmin(id uint) (*entity.Actor, error)
	DeleteAdmin(id uint, actor *entity.Actor) error
	LoginSuperAdmin(username string) (*entity.Actor, error)
	LoginAdmin(username string) (*entity.Actor, error)
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
