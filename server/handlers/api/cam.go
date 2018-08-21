package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/capsul"
	"github.com/medtune/beta-platform/pkg/session"
)

// MuraRunCam .
func MuraRunCam(c *gin.Context) {
	// Check session
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip")))
		return
	}

	// Parse data from body
	camData := jsonutil.RunImageCam{}
	err := c.ShouldBindJSON(&camData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	ctx := context.Background()
	// Run inference
	resp, err := capsul.RunMuraCAM(ctx, &camData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// TODO cam file creation
	// should be used as static

	// Success
	c.JSON(200, jsonutil.SuccessData(resp))
}
