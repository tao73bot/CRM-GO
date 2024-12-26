package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tao73bot/A_simple_CRM/helpers"
	"github.com/tao73bot/A_simple_CRM/initializers"
	"github.com/tao73bot/A_simple_CRM/models"
)

func CreateCustomer(c *gin.Context) {
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
	lid := uuid.MustParse(c.Param("lid"))
	uid := uuid.MustParse(claims.UserID)
	// Check lead status
	var lead models.Lead
	result := initializers.DB.First(&lead, lid)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Lead not found",
		})
		return
	}
	if lead.UserID != uid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to create customer for this lead",
		})
		return
	}
	if lead.Status != "qualified" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Lead status must be 'qualified' to create customer",
		})
		return
	}
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
	// Save the customer to the database
	customer := models.Customer{
		LeadID:      lid,
		UserID:      uid,
		Address:     body.Address,
		CompanyName: body.CompanyName,
	}
	lidCnt := initializers.DB.Where("lead_id = ?", lid).Find(&models.Customer{}).RowsAffected
	if lidCnt > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Customer already exists for this lead",
		})
		return
	}
	result = initializers.DB.Create(&customer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error saving the customer",
		})
		return
	}
	// Join customer with lead data
	var updatedCustomer struct {
		CustomerID  uuid.UUID
		LeadID      uuid.UUID
		UserID      uuid.UUID
		Name        string
		Email       string
		Phone       string
		Status      string
		Source      string
		Address     string
		CompanyName string
	}
	x := initializers.DB.Table("customers").
		Select("customers.*, leads.*").
		Joins("JOIN leads ON leads.lead_id = customers.lead_id").
		Where("customers.customer_id = ?", customer.CustomerID).
		Scan(&updatedCustomer)
	if x.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching the customer",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "Customer created successfully",
		"customer": updatedCustomer,
	})
}

func GetCustomers(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
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
	var updatedCustomer []struct {
		CustomerID  uuid.UUID
		LeadID      uuid.UUID
		UserID      uuid.UUID
		Name        string
		Email       string
		Phone       string
		Status      string
		Source      string
		Address     string
		CompanyName string
	}
	result := initializers.DB.Table("customers").
		Select("customers.*, leads.*").
		Joins("JOIN leads ON leads.lead_id = customers.lead_id").
		Scan(&updatedCustomer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching customers",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"customers": updatedCustomer,
	})
}

func GetCustomerByID(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
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

	// var customer models.Customer
	id := uuid.MustParse(c.Param("id"))
	var updatedCustomer struct {
		CustomerID  uuid.UUID
		LeadID      uuid.UUID
		UserID      uuid.UUID
		Name        string
		Email       string
		Phone       string
		Status      string
		Source      string
		Address     string
		CompanyName string
	}
	result := initializers.DB.Table("customers").
		Select("customers.*, leads.*").
		Joins("JOIN leads ON leads.lead_id = customers.lead_id").
		Where("customers.customer_id = ?", id).
		Scan(&updatedCustomer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching customers",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"customer": updatedCustomer,
	})
}

func GetCustomersOfUser(c *gin.Context) {
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
	uid := uuid.MustParse(claims.UserID)
	var updatedCustomer []struct {
		CustomerID  uuid.UUID
		LeadID      uuid.UUID
		UserID      uuid.UUID
		Name        string
		Email       string
		Phone       string
		Status      string
		Source      string
		Address     string
		CompanyName string
	}
	result := initializers.DB.Table("customers").
		Select("customers.*, leads.*").
		Joins("JOIN leads ON leads.lead_id = customers.lead_id").
		Where("customers.user_id = ?", uid).
		Scan(&updatedCustomer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching customers",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"customers": updatedCustomer,
	})
}

func GetCustomersByUserID(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
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
	uid := uuid.MustParse(c.Param("id"))
	var updatedCustomer []struct {
		CustomerID  uuid.UUID
		LeadID      uuid.UUID
		UserID      uuid.UUID
		Name        string
		Email       string
		Phone       string
		Status      string
		Source      string
		Address     string
		CompanyName string
	}
	fmt.Println("GetCustomersByUserID")
	result := initializers.DB.Table("customers").
		Select("customers.*, leads.*").
		Joins("JOIN leads ON leads.lead_id = customers.lead_id").
		Where("customers.user_id = ?", uid).
		Scan(&updatedCustomer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching customers",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"customers": updatedCustomer,
	})
}

func UpdateCustomerInfo(c *gin.Context) {
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
	uid := uuid.MustParse(claims.UserID)
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
	var customer models.Customer
	result := initializers.DB.Where("customer_id = ?", id).First(&customer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching the customer",
		})
		return
	}
	if customer.UserID != uid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to update this customer",
		})
		return
	}
	if body.Address != "" {
		customer.Address = body.Address
	}
	if body.CompanyName != "" {
		customer.CompanyName = body.CompanyName
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "Customer updated successfully",
		"customer": customer,
	})
}

func DeleteCustomer(c *gin.Context) {
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
	uid := uuid.MustParse(claims.UserID)
	id := uuid.MustParse(c.Param("id"))
	var customer models.Customer
	result := initializers.DB.Where("customer_id = ?", id).First(&customer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching the customer",
		})
		return
	}
	if customer.UserID != uid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to delete this customer",
		})
		return
	}
	result = initializers.DB.Delete(&customer)
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
