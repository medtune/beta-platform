package api

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/jsonutil"
)

// Copyrights description handler
func Copyrights(c *gin.Context) {
	c.JSON(200, jsonutil.SuccessData(pkg.GetCopyrights()))
}
