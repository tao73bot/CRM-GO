package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tao73bot/A_simple_CRM/initializers"
	"github.com/tao73bot/A_simple_CRM/models"
)

func CreateCustomer(c *gin.Context) {
	var body struct {
		LeadID      uuid.UUID `json:"lead_id"`
		UserID      uuid.UUID `json:"user_id"`
		Address     string    `json:"address"`
		CompanyName string    `json:"company_name"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}
	// Save the customer to the database
	customer := models.Customer{
		LeadID: 	body.LeadID,
		UserID: 	body.UserID,
		Address: 	body.Address,
		CompanyName: body.CompanyName,
	}

	result := initializers.DB.Create(&customer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error saving the customer",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer created successfully",
	})
}

func GetCustomers(c *gin.Context) {
	var customers []models.Customer

	result := initializers.DB.Find(&customers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching customers",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"customers": customers,
	})
}

func GetCustomerByID(c *gin.Context) {
	var customer models.Customer
	id := uuid.MustParse(c.Param("id"))
	result := initializers.DB.First(&customer, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching customer",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"customer": customer,
	})
}

func UpdateCustomerInfo(c *gin.Context) {
	var body struct {
		Address     string `json:"address"`
		CompanyName string `json:"company_name"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}
	id := uuid.MustParse(c.Param("id"))
	result := initializers.DB.Model(&models.Customer{}).Where("id = ?", id).Updates(body)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error updating the customer",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer updated successfully",
	})
}

func DeleteCustomer(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	result := initializers.DB.Where("id = ?", id).Delete(&models.Customer{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error deleting the customer",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer deleted successfully",
	})
}