package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lucksei/go-chart-image-analyzer-api/internal/utils"
)

// Simply stores the ResultStore item in the Gin context for use inside the routes
// Can be accessed using `c.Get("result_store")`
func ResultStore(r *utils.ResultStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("result_store", r)
		c.Next()
	}
}
