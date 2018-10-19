package api

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/jsonutil"
)

// Version description handler
func Version(c *gin.Context) {
	c.JSON(200, jsonutil.SuccessData(internal.GetVersion()))
}
