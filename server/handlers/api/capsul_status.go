package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/capsul"
	"github.com/medtune/beta-platform/pkg/session"
)

// CapsulStatus .
func CapsulStatus(c *gin.Context) {
	// Check session
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip")))
		return
	}

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
