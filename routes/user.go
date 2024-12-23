package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/controllers"
	"github.com/tao73bot/A_simple_CRM/middlewares"
)

func UserRoutes(r *gin.Engine) {
	r.Use(middlewares.AuthMiddleware())
	r.POST("/logout", controllers.Logout)
	r.POST("isLogged", controllers.IsUserLoggedIN)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:uid", controllers.GetUser)
}
