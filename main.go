package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/initializers"
	"github.com/tao73bot/A_simple_CRM/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.Migrate()
}

func main() {
	fmt.Println("Starting the CRM application ...")

	r := gin.Default()
	routes.UserRoutes(r)
	r.Run()
}
