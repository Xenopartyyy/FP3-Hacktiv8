package main

import (
	controller "FP3-Hacktiv8/controller"
	db "FP3-Hacktiv8/infra/database"
	"FP3-Hacktiv8/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.ConnDB()

	r.Use(middleware.RequiredAuth())

	r.POST("/users/register", controller.RegisterUser)
	r.POST("/users/login", controller.LoginUser)
	r.PUT("/users/update-account", middleware.RequiredAuth(), controller.UpdateUser)
	r.DELETE("/users/delete-account", middleware.RequiredAuth(), controller.DeleteUser)

	r.Use(middleware.CreateCategoryAuthorization())

	r.POST("/categories", middleware.CreateCategoryAuthorization(), controller.Create)

	r.Run(":8080")
}
