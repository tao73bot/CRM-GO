package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/controllers"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	r.GET("/is_logged_in", controllers.IsUserLoggedIN)
}
