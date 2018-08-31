package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/session"
)

// HeartBeat .
func HeartBeat(c *gin.Context) {
	// Check session
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied")))
		return
	}

	c.JSON(200, jsonutil.SuccessData(jsonutil.ServiceStatus{
		Healthy:  true,
		UnixTime: time.Now().Unix(),
	}))
}
