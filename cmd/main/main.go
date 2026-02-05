package main

import (
	"test/go_helm_chart_image_api/internal/routes"
	"test/go_helm_chart_image_api/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	apiGroup := router.Group("/api")
	apiGroup.GET("/health", routes.Health)

	helmChartGroup := apiGroup.Group("/helm-chart")
	helmChartGroup.POST("", routes.HelmChartPost)
	helmChartGroup.GET("/:id", routes.HelmChartGet)

	// NOTE: Very important to initialize helm sdk settings before running API
	err := utils.InitHelmSettings()
	if err != nil {
		panic(err)
	}

	router.Run()

}
