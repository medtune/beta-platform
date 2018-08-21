package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/session"
)

// Test .
func Test(c *gin.Context) {
	// Check session
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied")))
		return
	}

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
