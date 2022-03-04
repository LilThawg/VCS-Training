package routes

import (
	"github.com/challenge2/middlewares"
	"github.com/challenge2/controllers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files
	docs "github.com/challenge2/docs"
)

func ServerStart() {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	r.POST("/SignUp", controllers.SignUp)
	r.POST("/SignIn", controllers.SignIn)
	r.POST("/SignOut", controllers.SignOut)

	admin := r.Group("/admin")
	{
		admin.Use(middlewares.IsAdminMiddleWare())
		
		admin.GET("/getAllUsers", controllers.GetAllUsers)
		admin.POST("/addUser", controllers.AddUser)
		admin.DELETE("/deleteUser", controllers.DeleteUser)
		admin.PUT("/updateUser",controllers.UpdateUser)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}