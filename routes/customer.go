package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/controllers"
	"github.com/tao73bot/A_simple_CRM/middlewares"
)

func CustomerRoutes(r *gin.Engine) {
	r.Use(middlewares.AuthMiddleware())
	r.POST("/create_customer/:lid", controllers.CreateCustomer)
	r.GET("/get_customers", controllers.GetCustomers)
	r.GET("/get_customer/:id", controllers.GetCustomerByID)
	r.GET("/get_customer_of_user", controllers.GetCustomersOfUser)
	r.GET("/get_customer_by_user/:id", controllers.GetCustomersByUserID)
	r.PUT("/update_customer/:id", controllers.UpdateCustomerInfo)
	r.DELETE("/delete_customer/:id", controllers.DeleteCustomer)
}
