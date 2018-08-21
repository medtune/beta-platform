package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/demos"
	"github.com/medtune/beta-platform/pkg/session"
)

func staticPath(demo string, file string) string {
	return fmt.Sprintf("static/demos/%s/images/%s", demo, file)
}

// DatahubUpload .
func DatahubUpload(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}

	demo := c.PostForm("demo")
	if demo == "" {
		demo = "inception"
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	path := staticPath(demo, file.Filename)

	// Save file
	if err := c.SaveUploadedFile(file, "./"+path); err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	c.Redirect(302, "/datahub")
}

// DatahubDemoUpload .
func DatahubDemoUpload(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}

	demo := c.Param("demo")
	file, err := c.FormFile("file")
	if err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	path := staticPath(demo, file.Filename)

	// Save file
	if err := c.SaveUploadedFile(file, "./"+path); err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	c.Redirect(302, "/demos/"+demo)
}

// DatahubDemoDrop .
func DatahubDemoDrop(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}

	demo := c.Param("demo")

	// Parse data from body
	infData := jsonutil.RunImageInference{}
	if err := c.ShouldBindJSON(&infData); err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	path := staticPath(demo, infData.File)

	err := demos.DropImagePath(path)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
	}

	c.JSON(200, jsonutil.Success())
}
