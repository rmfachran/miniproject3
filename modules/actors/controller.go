package actors

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rmfachran/miniproject2/dto"
	"github.com/rmfachran/miniproject2/middleware"
	"time"
)

type ControllerActor interface {
	CreateAdmin(req ActorParam) (any, error)
	ApprovedAdmin(id uint) (any, error)
	GetAdmin(id uint, act ActorParam) (any, error)
	UpdateAdmin(id uint, adm ActorParam) (any, error)
	LoginSuperAdmin(username string, password string) (any, error)
	LoginAdmin(username string, password string) (any, error)
	GetCustomers(first_name, last_name, email string, page, pageSize int) (interface{}, error)
}

type controllerActor struct {
	actorUseCase useCaseActor
}

func (uc controllerActor) CreateAdmin(req ActorParam) (any, error) {
	actor, err := uc.actorUseCase.CreateAdmin(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success create Admin",
			Message:      "success register",
			ResponseTime: "",
		},
		Data: ActorParam{
			Username:   actor.Username,
			Password:   actor.Password,
			RoleId:     actor.RoleId,
			IsVerified: actor.IsVerified,
			IsActive:   actor.IsActive,
		},
	}
	return res, nil
}

func (uc controllerActor) UpdateAdmin(id uint, adm ActorParam) (any, error) {
	admin, err := uc.actorUseCase.UpdateAdmin(id, adm)
	if err != nil {
		return dto.ErrorResponse{}, err
	}
	res := FindAdmin{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success approve admin",
			Message:      "success",
			ResponseTime: "",
		},
		Data: ActorParam{
			Username:   admin.Username,
			Password:   admin.Password,
			RoleId:     admin.RoleId,
			IsVerified: admin.IsVerified,
			IsActive:   admin.IsActive,
		},
	}
	return res, nil
}

func (uc controllerActor) GetAdmin(id uint, act ActorParam) (any, error) {
	admin, err := uc.actorUseCase.GetAdminById(id)
	if err != nil {
		return dto.ErrorResponse{}, err
	}
	res := FindAdmin{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success get admin",
			Message:      "success",
			ResponseTime: "",
		},
		Data: ActorParam{
			Username:   admin.Username,
			Password:   admin.Password,
			RoleId:     admin.RoleId,
			IsVerified: admin.IsVerified,
			IsActive:   admin.IsActive,
			CreatedAt:  admin.CreatedAt,
			UpdatedAt:  admin.UpdatedAt,
		},
	}
	return res, nil
}

func (uc controllerActor) ApprovedAdmin(id uint) (any, error) {
	req, err := uc.actorUseCase.ApprovedAdmin(id)
	if err != nil {
		return dto.ErrorResponse{}, err
	}
	res := SuccessApproveAdmin{

		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "admin approved",
			Message:      "approved",
			ResponseTime: "",
		},
		Data: req,
	}
	return res, nil
}

func (uc controllerActor) LoginAdmin(username string, password string) (any, error) {
	admin, err := uc.actorUseCase.LoginAdmin(username, password)
	if err != nil {
		return nil, err
	}

	response := SuccessLoginAdmin{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success login admin",
			Message:      "success",
			ResponseTime: "",
		},
		Username: admin.Username,
	}
	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = admin.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set the token expiration time (e.g., 24 hours)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(middleware.JwtSecret))
	if err != nil {
		return response, errors.New("failed to generate JWT token")
	}

	// Return the token in the response
	response.Token = tokenString
	// Other response data

	return response, nil
}

func (uc controllerActor) LoginSuperAdmin(username string, password string) (any, error) {
	super, err := uc.actorUseCase.LoginSuperAdmin(username, password)
	if err != nil {
		return nil, err
	}

	response := SuccessLoginAdmin{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success login super admin",
			Message:      "success",
			ResponseTime: "",
		},
		Username: super.Username,
	}
	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = super.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set the token expiration time (e.g., 24 hours)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(middleware.JwtSecret))
	if err != nil {
		return response, errors.New("failed to generate JWT token")
	}

	// Return the token in the response
	response.Token = tokenString
	// Other response data
	return response, nil
}

func (uc controllerActor) GetCustomers(first_name, last_name, email string, page, pageSize int) (interface{}, error) {
	request, err := uc.actorUseCase.GetCustomers(first_name, last_name, email, page, pageSize)
	if err != nil {
		return SuccessGetCustomers{}, err
	}

	response := SuccessGetCustomers{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success get all customers",
			Message:      "success",
			ResponseTime: "",
		},
		Data: request,
	}
	return response, nil
}
