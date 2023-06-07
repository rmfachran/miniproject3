package actors

import (
	"crud/dto"
	"crud/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerActor struct {
	ctr ControllerActor
}

func NewActorRequestHandler(dbCrud *gorm.DB) RequestHandlerActor {
	return RequestHandlerActor{
		ctr: controllerActor{
			actorUseCase: useCaseActor{
				actorRepo: repository.NewActor(dbCrud),
			},
		},
	}
}

func (h RequestHandlerActor) CreateAdmin(c *gin.Context) {
	request := ActorParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateAdmin(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActor) GetAdmin(c *gin.Context) {
	request := ActorParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	adminId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}
	res, err := h.ctr.ApproveAdmin(uint(adminId), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActor) UpdateAdmin(c *gin.Context) {
	request := ActorParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	adminId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}
	res, err := h.ctr.UpdateAdmin(uint(adminId), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActor) ApproveAdmin(c *gin.Context) {
	request := ActorParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	adminId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}
	res, err := h.ctr.ApproveAdmin(uint(adminId), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActor) LoginAdmin(c *gin.Context) {
	request := ActorParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.LoginAdmin(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActor) LoginSuperAdmin(c *gin.Context) {
	request := ActorParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.LoginSuperAdmin(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}
	c.JSON(http.StatusOK, res)
}

//func loginHandler(c *gin.Context) {
//	var req LoginRequest
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//		return
//	}
//
//	user, err := authenticateUser(req.Username, req.Password)
//	if err != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
//		return
//	}
//
//	token := generateToken(user)
//
//	c.JSON(http.StatusOK, LoginResponse{Token: token})
//}
