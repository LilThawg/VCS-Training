package routes

import (
	"github.com/challenge3/middlewares"
	"github.com/challenge3/controllers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files
	docs "github.com/challenge3/docs"
)

func ServerStart() {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	r.POST("/SignUp", controllers.SignUp)
	r.POST("/SignIn", controllers.SignIn)
	r.POST("/SignOut", controllers.SignOut)

	api := r.Group("/api")
	{
		api.Use(middlewares.SetRoleEmail())
		api.GET("/getAllUsers", controllers.GetAllUsers)
		api.POST("/addUser", controllers.AddUser)
		api.DELETE("/deleteUser", controllers.DeleteUser)
		api.PUT("/updateUser",controllers.UpdateUser)
		api.POST("/addPost", controllers.AddPost)
		api.GET("/getAllPosts", controllers.GetAllPosts)
		api.DELETE("/deletePost", controllers.DeletePost)
		api.PUT("/updatePost",controllers.UpdatePost)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}