package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/controllers"
)

func LeadRoutes(r *gin.Engine) {
	r.POST("/create_lead", controllers.CreateLead)
	r.GET("/get_leads", controllers.GetLeads)
	r.GET("/get_lead/:id", controllers.GetLeadByID)
	r.GET("/get_lead_by_name/:name", controllers.GetLeadsByName)
	r.PUT("/update_lead/:id", controllers.UpdateLeadInfo)
	r.PATCH("/update_lead_status/:id", controllers.UpdateLeadStatus)
	r.DELETE("/delete_lead/:id", controllers.DeleteLead)
}
