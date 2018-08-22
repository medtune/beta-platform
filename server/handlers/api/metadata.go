package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/session"
)

// Version description handler
func Version(c *gin.Context) {
	// Check session
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied")))
		return
	}

	c.JSON(200, jsonutil.SuccessData(jsonutil.PackageVersion{
		Major:   0,
		Minor:   1,
		Patch:   2,
		Version: pkg.VERSION,
	}))
}
