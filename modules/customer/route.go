package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterCustomer struct {
	CustomerRequestHandler RequestHandlerCustomer
}

func NewCustomer(dbCrud *gorm.DB) RouterCustomer {
	return RouterCustomer{
		CustomerRequestHandler: NewCustomerRequestHandler(dbCrud),
	}
}

func (r RouterCustomer) Handle(router *gin.Engine) {
	basepath := "/customer"
	customer := router.Group(basepath)

	customer.POST("/register-customer",
		r.CustomerRequestHandler.CreateCustomer)

	customer.GET("/:id",
		r.CustomerRequestHandler.GetCustomerById,
	)

	customer.PUT("/:id",
		r.CustomerRequestHandler.UpdateCustomerById,
	)

	customer.DELETE("/:id",
		r.CustomerRequestHandler.DeleteCustomerById,
	)
}
