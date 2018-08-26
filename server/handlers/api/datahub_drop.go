package api

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/demo"
	"github.com/medtune/beta-platform/pkg/session"
)

// DemoDataDrop .
func DemoDataDrop(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}

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
