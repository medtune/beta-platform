package api

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/capsul"
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

	// get model status
	status, err := capsul.GetStatus(&getStatus)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	c.JSON(200, jsonutil.SuccessData(status))
}
