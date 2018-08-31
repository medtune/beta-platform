package api

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/session"
)

// DemoDataUpload .
func DemoDataUpload(c *gin.Context) {
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
