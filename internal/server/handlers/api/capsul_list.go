package api

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/beta-platform/internal/service/dashboard"
)

// CapsulList .
func CapsulList(c *gin.Context) {

	// get capsuls config
	list, err := dashboard.GetCapsulesList()
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	c.JSON(200, jsonutil.SuccessData(list))
}
