package api

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/beta-platform/internal/service/capsul"
)

// InceptionImagenetRunInference .
func InceptionImagenetRunInference(c *gin.Context) {
	// Parse data from body
	infData := jsonutil.RunImageInference{}
	err := c.ShouldBindJSON(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	if infData.File != "" {
		infData.File = staticPath("inception", infData.File)
	}

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
