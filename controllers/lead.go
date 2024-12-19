package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tao73bot/A_simple_CRM/initializers"
	"github.com/tao73bot/A_simple_CRM/models"
)


func CreateLead(c *gin.Context) {
	var body struct {
		UserID  uuid.UUID `json:"user_id"`
		Name    string    `json:"name"`
		Email   string    `json:"email"`
		Phone   string    `json:"phone"`
		Sources string    `json:"sources"`
		Status  string    `json:"status"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}
	// Save the lead to the database
	lead := models.Lead{
		UserID: body.UserID,
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

func GetLeads(c *gin.Context) {
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

func GetLeadsByName(c *gin.Context) {
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
	var body struct {
		Name   string `json:"name"`
		Email  string `json:"email"`
		Phone  string `json:"phone"`
		Source string `json:"source"`
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
	lead.Name = body.Name
	lead.Email = body.Email
	lead.Phone = body.Phone
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
