package main

import (
	"fmt"
	"net/http"
	"quoteapp/controller"
	"quoteapp/authentication"
	"quoteapp/db"
	"quoteapp/model"
	"quoteapp/routes"
	"strconv"
)

func main() {

	port := authentication.LoadEnv("port")
	portConvert , _ := strconv.Atoi(port)
	db := db.NewDB(
		authentication.LoadEnv("host"), 
		authentication.LoadEnv("username"),
		authentication.LoadEnv("password"),
		authentication.LoadEnv("database"),
		portConvert,
	)

	db.Migrate()
	conn := db.DB()

	qm := model.NewQuoteModel(conn)
	qoute := controller.NewQuoteController(qm)


	usermodel := model.NewUsersModel(conn)
	usercontroller := controller.NewUsersController(usermodel)
	authcontroller := controller.NewAuthController(usermodel)

	router := routes.NewRoute(qoute, usercontroller, authcontroller)

	fmt.Println("starting api server at http://localhost:8080")
	panic(http.ListenAndServe(":8080", router.Run()))
}
