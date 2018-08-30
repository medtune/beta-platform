package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/capsul"
	"github.com/medtune/beta-platform/pkg/session"
)

// InceptionImagenetRunInference .
func InceptionImagenetRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip")))
		return
	}

	// Parse data from body
	infData := jsonutil.RunImageInference{}
	err := c.ShouldBindJSON(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	infData.File = staticPath("inception", infData.File)

	ctx := context.Background()
	// Run inference
	result, err := capsul.RunInceptionInference(ctx, &infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Success
	c.JSON(200, jsonutil.SuccessData(result))
}
