package api

import (
	"github.com/gin-gonic/gin"
)

// DemoDataUpload .
func DemoDataUpload(c *gin.Context) {
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
