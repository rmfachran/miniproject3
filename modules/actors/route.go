package actors

import (
	"crud/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterActor struct {
	ActorRequestHandler RequestHandlerActor
}

func NewActor(dbCrud *gorm.DB) RouterActor {
	return RouterActor{
		ActorRequestHandler: NewActorRequestHandler(dbCrud),
	}
}

func (r RouterActor) Handle(router *gin.Engine) {
	basepath := "/admin"
	admin := router.Group(basepath)
	pathSuperAdmin := "/superadmin"
	superAdmin := router.Group(pathSuperAdmin)

	admin.POST("/register-admin",
		r.ActorRequestHandler.CreateAdmin,
	)

	superAdmin.POST("/login",
		r.ActorRequestHandler.LoginSuperAdmin,
	)

	admin.POST("/login",
		r.ActorRequestHandler.LoginAdmin,
	)
	superAdmin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AuthMiddleware())
	//admin
	admin.GET("/:id",
		r.ActorRequestHandler.GetAdmin,
	)
	admin.GET("/customers",
		r.ActorRequestHandler.GetCustomers,
	)
	admin.PUT("/update/:id",
		r.ActorRequestHandler.UpdateAdmin,
	)
	//superadmin
	superAdmin.PUT("/approve/:id",
		r.ActorRequestHandler.ApproveAdmin,
	)

	superAdmin.POST("/register-admin",
		r.ActorRequestHandler.CreateAdmin,
	)
	superAdmin.GET("/:id",
		r.ActorRequestHandler.GetAdmin,
	)
	admin.GET("/customers/",
		r.ActorRequestHandler.GetCustomers,
	)

}
