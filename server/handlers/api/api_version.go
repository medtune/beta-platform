package api

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/jsonutil"
)

// Version description handler
func Version(c *gin.Context) {
	// Check session

	c.JSON(200, jsonutil.SuccessData(jsonutil.PackageVersion{
		Major:   0,
		Minor:   1,
		Patch:   2,
		Version: pkg.VERSION,
	}))
}
