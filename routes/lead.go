package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/controllers"
	"github.com/tao73bot/A_simple_CRM/middlewares"
)

func LeadRoutes(r *gin.Engine) {
	r.Use(middlewares.AuthMiddleware())
	r.POST("/create_lead", controllers.CreateLead)
	r.GET("/get_all_leads", controllers.GetAllLeads)
	r.GET("/get_leads_by_user", controllers.GetLeadsByUser)
	r.GET("/get_lead/:id", controllers.GetLeadByID)
	r.GET("/get_lead_by_name/:name", controllers.GetLeadsByName)
	r.PATCH("/update_lead/:id", controllers.UpdateLeadInfo)
	//r.PATCH("/update_lead_status/:id", controllers.UpdateLeadStatus)
	r.DELETE("/delete_lead/:id", controllers.DeleteLead)
}
