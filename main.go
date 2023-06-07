package main

import (
	_ ".github/joho/godotenv/autoload"
	"crud/modules/actors"
	"crud/modules/customer"
	"crud/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
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
