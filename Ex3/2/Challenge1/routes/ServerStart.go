package routes

import (
	"github.com/challenge1/controllers"
	"github.com/gin-gonic/gin"
	docs "github.com/challenge1/docs"
	swaggerfiles "github.com/swaggo/files"
   	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API CRUD 
// @description Đây là API của challenge 1
// @version 1.0
// @host localhost:8080
// @BasePath /api

func ServerStart() {

	r := gin.Default()

	
	docs.SwaggerInfo.BasePath = "/api"
	api := r.Group("/api")
	{
		api.GET("/getAllUsers", controllers.GetAllUsers)
		api.POST("/addUser", controllers.AddUser)
		api.DELETE("/deleteUser", controllers.DeleteUser)
		api.PUT("/updateUser", controllers.UpdateUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}