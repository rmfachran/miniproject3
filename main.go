package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rmfachran/miniproject2/modules/actors"
	"github.com/rmfachran/miniproject2/modules/customer"
	"github.com/rmfachran/miniproject2/utils/db"
)

func main() {
	router := gin.New()

	// open connection db
	dbCrud := db.GormMysql()

	////check connection
	//checkdb, err := dbCrud.DB()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////ping to database
	//errconn := checkdb.Ping()
	//if err != nil {
	//	log.Fatal(errconn)
	//}

	fmt.Println("database connected..!")

	userHandler := actors.NewActor(dbCrud)
	userHandler.Handle(router)

	custHandler := customer.NewCustomer(dbCrud)
	custHandler.Handle(router)

	errRouter := router.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}

}
