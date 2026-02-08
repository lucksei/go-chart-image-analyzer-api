package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucksei/go-chart-image-analyzer-api/internal/utils"
)

func HelmChartPost(c *gin.Context) {
	// Validating the request
	var jsonBody utils.HelmChartSource
	err := c.BindJSON(&jsonBody)
	if err != nil {
		c.Error(err)
	}

	// Create an ID for the specific helm chart source
	helmChartSource := jsonBody
	helmChartId, err := helmChartSource.ToBase64Id()
	if err != nil {
		c.Error(err)
	}

	// Loading the result store
	rs, ok := c.MustGet("result_store").(*utils.ResultStore)
	if !ok {
		c.Error(err)
	}

	// If the helm chart is being processed, accept (202)
	_, status := rs.Get(helmChartId)
	if status == utils.StatusInProgress {
		c.Writer.Header().Set("Location", fmt.Sprintf("/api/helm-chart/%s", helmChartId))
		c.Status(http.StatusAccepted)
		return
	}
	// If the helm chart is already in the store, redirect (303)
	if status == utils.StatusSuccess {
		c.Writer.Header().Set("Location", fmt.Sprintf("/api/helm-chart/%s", helmChartId))
		c.Status(http.StatusSeeOther)
		return
	}

	// Runs in the background, processes the helm chart + images if they are not already inside the store
	rs.SetPending(helmChartId)
	go func() {
		fmt.Printf("Processing helm chart %s\n", helmChartSource.ChartRef)
		rendered, err := utils.RenderHelmTemplate(helmChartSource)
		if err != nil {
			panic(err)
		}

		images := utils.GetImagesFromRendered(rendered)
		fmt.Printf("Found %d images\n", len(images))
		fmt.Printf("Processing images\n")
		result := utils.HelmChartAnalysis{
			RepoURL:  helmChartSource.RepoURL,
			ChartRef: helmChartSource.ChartRef,
		}
		imagesAnalysis := []utils.ImageAnalysis{}
		for i, image := range images {
			imageAnalysis, err := utils.PullImageAndParseAPIInfo(image)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Image %d: %s\n", i, imageAnalysis.Name)
			imagesAnalysis = append(imagesAnalysis, imageAnalysis)
		}
		result.Images = imagesAnalysis
		fmt.Printf("%v\n", result)

		rs.Put(helmChartId, result)
		fmt.Printf("Done processing helm chart %s\n", helmChartSource.ChartRef)
	}()

	c.Writer.Header().Set("Location", fmt.Sprintf("/api/helm-chart/%s", helmChartId))
	c.Status(http.StatusAccepted)
}

func HelmChartGet(c *gin.Context) {
	id := c.Param("id")

	// Loading the result store
	rs, ok := c.MustGet("result_store").(*utils.ResultStore)
	if !ok {
		c.Error(fmt.Errorf("Failed to retrieve ResultStore"))
	}
	result, status := rs.Get(id)
	if status == utils.StatusInProgress {
		// TODO: Change response to show that the request is in progress
		return
	}
	if status == utils.StatusNotFound {
		c.Status(http.StatusNotFound)
		// TODO: Change the response to show that the helm chart was not found, and that it has to be processed by the POST /apu/helm-chart endpoint
		return
	}

	c.JSON(http.StatusOK, result)
}
