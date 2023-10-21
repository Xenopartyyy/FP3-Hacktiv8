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

	// r.Use(middleware.Authentication())

	r.POST("/users/register", controller.RegisterUser)
	r.POST("/users/login", controller.LoginUser)
	r.PUT("/users/update-account", middleware.Authentication(), controller.UpdateUser)
	r.DELETE("/users/delete-account", middleware.Authentication(), controller.DeleteUser)

	// r.Use(middleware.Authorization())

	// r.GET("/categories", controller.GetCategory)
	r.POST("/categories", middleware.Authentication(), middleware.AdminAuthorization(), controller.CreateCategory)
	r.PATCH("/categories/:categoryID", middleware.Authentication(), middleware.AdminAuthorization(), controller.PatchCategory)
	r.DELETE("/categories/:categoryID", middleware.Authentication(), middleware.AdminAuthorization(), controller.DeleteCategory)

	r.POST("/tasks", middleware.Authentication(), controller.CreateTask)
	r.PUT("/tasks/:taskID", middleware.Authentication(), controller.PutPatchTask)
	r.PATCH("/tasks/:taskID", middleware.Authentication(), controller.PutPatchTask)

	r.Run(":8080")
}
