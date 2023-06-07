package actors

import (
	"crud/dto"
	"crud/middleware"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type ControllerActor interface {
	CreateAdmin(req ActorParam) (any, error)
	ApproveAdmin(id uint, act ActorParam) (any, error)
	GetAdmin(id uint, act ActorParam) (any, error)
	UpdateAdmin(id uint, adm ActorParam) (any, error)
	LoginSuperAdmin(username string, password string) (any, error)
	LoginAdmin(username string, password string) (any, error)
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
	admin, err := uc.actorUseCase.GetAdmin(id)
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

func (uc controllerActor) ApproveAdmin(id uint, act ActorParam) (any, error) {
	admin, err := uc.actorUseCase.ApproveAdmin(id, act)
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
			MessageTitle: "success login admin",
			Message:      "success",
			ResponseTime: "",
		},
		Username: super.Username,
		Token:    "asdasdasid",
	}
	return response, nil
}
