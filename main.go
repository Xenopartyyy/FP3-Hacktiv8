package main

import (
	usercontroller "FP3-Hacktiv8/controller"
	db "FP3-Hacktiv8/infra/database"
	"FP3-Hacktiv8/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.ConnDB()

	r.POST("/users/register", usercontroller.Register)
	r.POST("/users/login", usercontroller.Login)
	r.Use(middleware.RequiredAuth())
	r.PUT("/users/update-account", middleware.RequiredAuth(), usercontroller.Update)
	r.DELETE("/users/delete-account", middleware.RequiredAuth(), usercontroller.Delete)

	r.Run(":8080")
}
