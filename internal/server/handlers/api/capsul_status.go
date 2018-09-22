package api

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/beta-platform/internal/service/capsul"
)

// CapsulStatus .
func CapsulStatus(c *gin.Context) {
	// Parse data from body
	getStatus := jsonutil.GetStatus{}
	err := c.ShouldBindJSON(&getStatus)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// get model status
	status, err := capsul.GetStatus(ctx, &getStatus)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	c.JSON(200, jsonutil.SuccessData(status))
}
