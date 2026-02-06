package middleware

import (
	"fmt"
	"test/go_helm_chart_image_api/internal/utils"

	"github.com/gin-gonic/gin"
)

func ResultStore(r *utils.ResultStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("result_store", r)
		// TODO: Delete the print later
		fmt.Printf("Using result storage")
		c.Next()
	}
}
