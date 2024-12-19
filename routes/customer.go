package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/controllers"
)

func CustomerRoutes(r *gin.Engine) {
	r.POST("/create_customer", controllers.CreateCustomer)
	r.GET("/get_customers", controllers.GetCustomers)
	r.GET("/get_customer/:id", controllers.GetCustomerByID)
	r.PUT("/update_customer/:id", controllers.UpdateCustomerInfo)
	r.DELETE("/delete_customer/:id", controllers.DeleteCustomer)
}
