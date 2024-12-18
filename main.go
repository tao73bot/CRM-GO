package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.Migrate()
}

func main() {
	fmt.Println("Starting the CRM application ...")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the CRM application",
		})
	})

	r.Run()
}
