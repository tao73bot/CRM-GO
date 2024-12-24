package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tao73bot/A_simple_CRM/helpers"
	"github.com/tao73bot/A_simple_CRM/initializers"
	"github.com/tao73bot/A_simple_CRM/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name, email, and password are required",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error hashing the password",
		})
		return
	}
	// Save the user to the database
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error saving the user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Sign up successful",
		"user":    user,
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email and password are required",
		})
		return
	}
	var user models.User
	result := initializers.DB.Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	token, refreshToken, err := helpers.GenerateAllTokens(user.Email, user.Name, user.Role, user.UserID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	// Store the refresh token in the cookie
	c.SetCookie("Authorization", refreshToken, 60*60*24*7, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

func Logout(c *gin.Context) {
	accesstoken, exists := c.Request.Header["Authorization"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No token found",
		})
		return
	}
	fmt.Println(accesstoken[0])
	helpers.InvalidateToken(accesstoken[0])
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Error logging out",
	// 	})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}

func IsUserLoggedIN(c *gin.Context) {
	accessToken, exists := c.Request.Header["Authorization"]
	fmt.Println(accessToken)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No token found",
		})
		return
	}
	fmt.Println("Gelo1")
	_, err := jwt.Parse(accessToken[0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}
	fmt.Println("Gelo")
	for _, token := range helpers.BlockList {
		if token == accessToken[0] {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User is not logged in",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User is logged in",
	})
}

func GetUser(c *gin.Context) {
	uid := uuid.MustParse(c.Param("uid"))
	if err := helpers.MatchUserRoleToUid(c, uid); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	result := initializers.DB.Where("user_id = ?", uid).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUsers(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching users",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetAnotherUserInfo(c *gin.Context) {
	uid := uuid.MustParse(c.Param("uid"))
	var user models.User
	result := initializers.DB.Where("user_id = ?", uid).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}

func UpdateUserDetails(c *gin.Context) {
	uid := uuid.MustParse(c.Param("uid"))
	if err := helpers.MatchUserRoleToUid(c, uid); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	result := initializers.DB.Where("user_id = ?", uid).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name and emailare required",
		})
		return
	}
	user.Name = body.Name
	user.Email = body.Email
	result = initializers.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error saving the user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User details updated",
		"user":    user,
	})
}

func UpdateUserPassword(c *gin.Context) {
	uid := uuid.MustParse(c.Param("uid"))
	if err := helpers.MatchUserRoleToUid(c, uid); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	result := initializers.DB.Where("user_id = ?", uid).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	var body struct {
		CurrentPassword    string `json:"current_password"`
		NewPassword        string `json:"new_password"`
		ConfirmNewPassword string `json:"confirm_new_password"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.CurrentPassword))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid current password",
		})
		return
	}
	if body.NewPassword != body.ConfirmNewPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Passwords do not match",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error hashing the password",
		})
		return
	}
	user.Password = string(hash)
	result = initializers.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error saving the user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User password updated",
	})
}

func UpdateUserRole(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	uid := uuid.MustParse(c.Param("uid"))
	var user models.User
	result := initializers.DB.Where("user_id = ?", uid).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	var body struct {
		Role string `json:"role"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Role is required",
		})
		return
	}
	user.Role = body.Role
	result = initializers.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error saving the user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User role updated",
		"user":    user,
	})
}

func DeleteUser(c *gin.Context) {
	if err := helpers.CheckUserRoles(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	uid := uuid.MustParse(c.Param("uid"))
	var user models.User
	result := initializers.DB.Where("user_id = ?", uid).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	result = initializers.DB.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error deleting the user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
