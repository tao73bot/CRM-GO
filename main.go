package main

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
	"github.com/tao73bot/A_simple_CRM/initializers"
	"github.com/tao73bot/A_simple_CRM/routes"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.Migrate()
}

func main() {
	fmt.Println("Starting the CRM application ...")

	r := gin.Default()
	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.LeadRoutes(r)
	routes.CustomerRoutes(r)
	url := ginSwagger.URL("http://localhost:3333/swagger/doc.json")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()
}
