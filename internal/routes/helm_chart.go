package routes

import (
	"fmt"
	"net/http"
	"strings"
	"test/go_helm_chart_image_api/internal/models"
	"test/go_helm_chart_image_api/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"
)

func HelmChartPost(c *gin.Context) {
	var jsonBody models.HelmChartRequest
	err := c.BindJSON(&jsonBody)
	if err != nil {
		c.Error(err)
	}

	helmChartPath := utils.HelmChartPath{
		RepoURL:   jsonBody.RepoURL,
		ChartPath: jsonBody.ChartPath,
	}
	helmChartId, err := helmChartPath.ToBase64Id()
	if err != nil {
		c.Error(err)
	}

	go func() {
		rendered, err := utils.RenderHelmTemplate(helmChartPath)
		if err != nil {
			panic(err)
		}

		for key, value := range rendered {
			if strings.Contains(key, ".yaml") {
				fmt.Printf("* Template: %s\n", key)
				var template map[string]interface{}
				yaml.Unmarshal([]byte(value), &template)

				var containersSpecList []any
				utils.ContainersSpecSearch(template, &containersSpecList)
				for _, containers := range containersSpecList {
					for _, container := range containers.([]any) {
						fmt.Printf("%v\n", container.(map[string]any)["image"])
					}
				}
			}
		}
	}()

	c.Writer.Header().Set("Location", fmt.Sprintf("/api/helm-chart/%s", helmChartId))
	c.Status(http.StatusAccepted)
}

func HelmChartGet(c *gin.Context) {
	id := c.Param("id")
	helmChartPath, err := utils.Base64StringToHelmChart(id)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, models.HelmChartResponse{
		RepoURL:   helmChartPath.RepoURL,
		ChartPath: helmChartPath.ChartPath,
		Images:    []models.ChartImage{},
	})
}
