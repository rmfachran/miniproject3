package customer

import (
	"crud/dto"
)

type CustomerParam struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type FindCustomer struct {
	dto.ResponseMeta
	Data CustomerParam `json:"data"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data CustomerParam `json:"data"`
}

type SuccessUpdate struct {
	dto.ResponseMeta
	Data CustomerParam `json:"data"`
}
