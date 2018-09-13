package api

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/demo"
)

// PathologyAL .
func PathologyAL(c *gin.Context) {
	var req *jsonutil.GetPathologyAL
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	response, err := demo.GetCXPBAForPathology(req)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	c.JSON(200, jsonutil.SuccessData(response))
}

// SpecAnalysisPool .
func SpecAnalysisPool(c *gin.Context) {
	req := jsonutil.GetSpecPoolGrid{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	result, err := demo.GetSpecPoolGrid(&req)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	c.JSON(200, jsonutil.SuccessData(result))
}
