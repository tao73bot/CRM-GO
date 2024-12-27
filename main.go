package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tao73bot/A_simple_CRM/initializers"
	"github.com/tao73bot/A_simple_CRM/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.Migrate()
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]

func main() {
	fmt.Println("Starting the CRM application ...")

	r := gin.Default()
	r.StaticFS("/docs", http.Dir("./docs"))
	r.GET("swagger/*any", ginSwagger.CustomWrapHandler(
		&ginSwagger.Config{
			URL: "/docs/swagger.yaml",
		}, swaggerFiles.Handler))
	routes.UserRoutes(r)
	routes.LeadRoutes(r)
	routes.CustomerRoutes(r)
	routes.InteractionRoutes(r)
	routes.AuthRoutes(r)
	r.Run()
}
