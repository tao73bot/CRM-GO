package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/controllers"
	"github.com/tao73bot/A_simple_CRM/middlewares"
)

func InteractionRoutes(r *gin.Engine) {
	r.Use(middlewares.AuthMiddleware())
	r.POST("/create_interaction_with_lead/:lid", controllers.CreateInteractionWithLead)
	// r.POST("/create_interaction_with_customer/:cid", controllers.CreateInteractionWithCustomer)
	r.PUT("/update_interaction/:iid", controllers.UpdateNoteOfaInteraction)
}
