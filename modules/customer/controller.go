package customer

import (
	"github.com/rmfachran/miniproject2/dto"
)

type ControllerCustomer interface {
	CreateCustomer(req CustomerParam) (any, error)
	GetCustomerById(id uint) (any, error)
	UpdateCustomerById(id uint, cust CustomerParam) (any, error)
	DeleteCustomerById(id uint) error
}

type controllerCustomer struct {
	customerUseCase useCaseCustomer
}

func (uc controllerCustomer) CreateCustomer(req CustomerParam) (any, error) {

	customer, err := uc.customerUseCase.CreateCustomer(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create customer",
			Message:      "Success register",
			ResponseTime: "",
		},
		Data: CustomerParam{
			Email:     customer.Email,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}

func (uc controllerCustomer) UpdateCustomerById(id uint, cust CustomerParam) (any, error) {
	//var res FindCustomer
	customer, err := uc.customerUseCase.UpdateCustomerById(id, cust)
	if err != nil {
		return dto.ErrorResponse{}, err
	}
	res := FindCustomer{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success update customer",
			Message:      "success",
			ResponseTime: "",
		},
		Data: CustomerParam{
			Email:     customer.Email,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}

func (uc controllerCustomer) DeleteCustomerById(id uint) error {
	err := uc.customerUseCase.DeleteCustomerById(id)
	if err != nil {
		return err
	}
	return err
}

func (uc controllerCustomer) GetCustomerById(id uint) (any, error) {
	//var res FindCustomer
	customer, err := uc.customerUseCase.GetCustomerById(id)
	if err != nil {
		return FindCustomer{}, err
	}
	res := FindCustomer{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success get customer",
			Message:      "success",
			ResponseTime: "",
		},
		Data: CustomerParam{
			Email:     customer.Email,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}
