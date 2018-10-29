package api

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/jsonutil"
	"github.com/medtune/beta-platform/internal/service/demo"
)

// DemoDataDrop .
func DemoDataDrop(c *gin.Context) {
	demoname := c.Param("demo")

	// Parse data from body
	infData := jsonutil.RunImageInference{}
	if err := c.ShouldBindJSON(&infData); err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	path := staticPath(demoname, infData.File)

	err := demo.DropImagePath(path)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
	}

	c.JSON(200, jsonutil.Success())
}
