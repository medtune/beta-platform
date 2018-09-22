package api

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/beta-platform/internal/service/capsul"
)

// MnistRunInference .
func MnistRunInference(c *gin.Context) {
	// Parse data from body
	infData := jsonutil.RunImageInference{}
	err := c.ShouldBindJSON(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	ctx := context.Background()

	if infData.File != "" {
		infData.File = staticPath("mnist", infData.File)
	}

	// Run inference
	resp, err := capsul.RunMnistInference(ctx, &infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	// Success
	c.JSON(200, jsonutil.SuccessData(resp))
}
