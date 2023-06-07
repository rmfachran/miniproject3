package user

import (
	"crud/dto"
	"crud/entity"
)

type UserParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FindUser struct {
	dto.ResponseMeta
	Data entity.User `json:"data"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data UserParam `json:"data"`
}
