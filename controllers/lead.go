package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tao73bot/A_simple_CRM/helpers"
	"github.com/tao73bot/A_simple_CRM/initializers"
	"github.com/tao73bot/A_simple_CRM/models"
)

func CreateLead(c *gin.Context) {
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

	var body struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Sources string `json:"sources"`
		Status  string `json:"status"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}
	// Save the lead to the database
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	lead := models.Lead{
		UserID: userID,
		Name:   body.Name,
		Email:  body.Email,
		Phone:  body.Phone,
		Source: body.Sources,
		Status: body.Status,
	}
	result := initializers.DB.Create(&lead)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error saving the lead",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Lead created successfully",
	})
}

func GetAllLeads(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to view all leads",
		})
		return
	}
	tokenString := c.Request.Header.Get("Authorization")
	_, err := helpers.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var leads []models.Lead

	result := initializers.DB.Find(&leads)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching leads",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"leads": leads,
	})
}

func GetLeadsByUser(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to view the leads",
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
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	var leads []models.Lead
	result := initializers.DB.Where("user_id = ?", userID).Find(&leads)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching leads",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"leads": leads,
	})
}

func GetLeadsByName(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to view the leads",
		})
		return
	}
	tokenString := c.Request.Header.Get("Authorization")
	_, err := helpers.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	name := c.Param("name")
	var lead models.Lead

	result := initializers.DB.Where("name = ?", name).Find(&lead)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching the lead",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"lead": lead,
	})
}

func GetLeadByID(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to view the leads",
		})
		return
	}
	tokenString := c.Request.Header.Get("Authorization")
	_, err := helpers.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid lead ID",
		})
		return
	}
	var lead models.Lead

	result := initializers.DB.Where("lead_id = ?", id).Find(&lead)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching the lead",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"lead": lead,
	})
}

func UpdateLeadInfo(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to view the leads",
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
	var body struct {
		Name   string `json:"name"`
		Email  string `json:"email"`
		Phone  string `json:"phone"`
		Source string `json:"source"`
		Status string `json:"status"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid lead ID",
		})
		return
	}
	var lead models.Lead

	result := initializers.DB.Where("lead_id = ?", id).First(&lead)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching the lead",
		})
		return
	}
	if lead.UserID != uuid.MustParse(claims.UserID) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to update this lead",
		})
		return
	}
	if body.Name != "" {
		lead.Name = body.Name
	}
	if body.Email != "" {
		lead.Email = body.Email
	}
	if body.Phone != "" {
		lead.Phone = body.Phone
	}
	if body.Source != "" {
		lead.Source = body.Source
	}
	if body.Status != "" {
		lead.Status = body.Status
	}
	result = initializers.DB.Save(&lead)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error updating the lead",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Lead updated successfully",
	})
}

func UpdateLeadStatus(c *gin.Context) {
	var body struct {
		Status string `json:"status"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Status is required",
		})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid lead ID",
		})
		return
	}
	var lead models.Lead

	result := initializers.DB.Where("lead_id = ?", id).First(&lead)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching the lead",
		})
		return
	}
	lead.Status = body.Status
	result = initializers.DB.Save(&lead)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error updating the lead",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Lead status updated successfully",
	})
}

func DeleteLead(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to delete the lead",
		})
		return
	}
	tokenString := c.Request.Header.Get("Authorization")
	claims, err := helpers.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You token is invalid or expired",
		})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid lead ID",
		})
		return
	}
	var lead models.Lead

	result := initializers.DB.Where("lead_id = ?", id).First(&lead)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching the lead",
		})
		return
	}
	if lead.UserID != uuid.MustParse(claims.UserID) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to delete this lead",
		})
		return
	}
	result = initializers.DB.Delete(&lead)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error deleting the lead",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Lead deleted successfully",
	})
}
