package customer

import (
	"github.com/rmfachran/miniproject2/dto"
)

type CustomerParam struct {
	ID        uint   `json:"id"`
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
