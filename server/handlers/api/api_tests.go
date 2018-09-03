package api

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/jsonutil"
)

// Test .
func Test(c *gin.Context) {
	// Parse data from body
	testData := jsonutil.TestRequest{}
	err := c.ShouldBindJSON(&testData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	if testData.Sleep != 0 {
		time.Sleep(time.Duration(testData.Sleep) * time.Second)
	}

	c.JSON(200, jsonutil.SuccessData(jsonutil.TestResponse{
		Test: true,
	}))
}
