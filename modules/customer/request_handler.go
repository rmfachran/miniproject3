package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/rmfachran/miniproject2/dto"
	"github.com/rmfachran/miniproject2/repository"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerCustomer struct {
	ctr ControllerCustomer
}

func NewCustomerRequestHandler(dbCrud *gorm.DB) RequestHandlerCustomer {
	return RequestHandlerCustomer{
		ctr: controllerCustomer{
			customerUseCase: useCaseCustomer{
				customerRepo: repository.NewCustomer(dbCrud),
			}},
	}
}

func (h RequestHandlerCustomer) CreateCustomer(c *gin.Context) {
	request := CustomerParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) UpdateCustomerById(c *gin.Context) {
	request := CustomerParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	customerId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}
	res, err := h.ctr.UpdateCustomerById(uint(customerId), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) DeleteCustomerById(c *gin.Context) {
	request := CustomerParam{}
	//customerId := c.Param("id")

	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	customerId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	err = h.ctr.DeleteCustomerById(uint(customerId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}
	c.JSON(http.StatusOK, nil)
}

func (h RequestHandlerCustomer) GetCustomerById(c *gin.Context) {
	request := CustomerParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	CustomerId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetCustomerById(uint(CustomerId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}
