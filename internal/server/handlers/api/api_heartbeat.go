package api

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/jsonutil"
)

// HeartBeat .
func HeartBeat(c *gin.Context) {
	c.JSON(200, jsonutil.SuccessData(jsonutil.ServiceStatus{
		Healthy:  true,
		UnixTime: time.Now().Unix(),
	}))
}
