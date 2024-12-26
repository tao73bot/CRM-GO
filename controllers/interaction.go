package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tao73bot/A_simple_CRM/helpers"
	"github.com/tao73bot/A_simple_CRM/initializers"
	"github.com/tao73bot/A_simple_CRM/models"
)

func CreateInteractionWithLead(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	tokenString := c.Request.Header.Get("Authorization")
	claims, err := helpers.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	// Check if the lead exists
	lid := uuid.MustParse(c.Param("lid"))
	var lead models.Lead
	result := initializers.DB.First(&lead, lid)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Lead not found",
		})
		return
	}
	uid := uuid.MustParse(claims.UserID)
	var body struct {
		Type  string `json:"type"`
		Notes string `json:"notes"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}

	// Save the interaction to the database
	interaction := models.Interaction{
		LeadID:     lid,
		UserID:     uid,
		Type:       body.Type,
		Notes:      body.Notes,
	}
	result = initializers.DB.Create(&interaction)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":     "Interaction created",
		"interaction": interaction,
	})
}

// func CreateInteractionWithCustomer(c *gin.Context) {
// 	if err := helpers.CheckUserRoles(c, "user"); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Unauthorized",
// 		})
// 		return
// 	}
// 	tokenString := c.Request.Header.Get("Authorization")
// 	claims, err := helpers.ValidateToken(tokenString)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Unauthorized",
// 		})
// 		return
// 	}
// 	// Check if the customer exists
// 	cid := uuid.MustParse(c.Param("cid"))
// 	var customer models.Customer
// 	result := initializers.DB.First(&customer, cid)
// 	if result.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "Customer not found",
// 		})
// 		return
// 	}
// 	uid := uuid.MustParse(claims.UserID)
// 	var body struct {
// 		Type  string `json:"type"`
// 		Notes string `json:"notes"`
// 	}
// 	if c.Bind(&body) != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "All fields are required",
// 		})
// 		return
// 	}
// 	// Save the interaction to the database
// 	interaction := models.Interaction{
// 		CustomerID: cid,
// 		LeadID:     customer.LeadID,
// 		UserID:     uid,
// 		Type:       body.Type,
// 		Notes:      body.Notes,
// 	}
// 	initializers.DB.Create(&interaction)
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message":     "Interaction created",
// 		"interaction": interaction,
// 	})
// }

func UpdateNoteOfaInteraction(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	tokenString := c.Request.Header.Get("Authorization")
	claims, err := helpers.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	iid := uuid.MustParse(c.Param("iid"))
	var interaction models.Interaction
	result := initializers.DB.First(&interaction, iid)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Interaction not found",
		})
		return
	}
	uid := uuid.MustParse(claims.UserID)
	if interaction.UserID != uid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	var body struct {
		Notes string `json:"notes"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}
	interaction.Notes = body.Notes
	initializers.DB.Save(&interaction)
	c.JSON(http.StatusOK, gin.H{
		"message": "Interaction updated",
	})
}
