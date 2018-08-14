package api

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/session"
)

// DatahubUpload .
func DatahubUpload(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}

	//name := c.PostForm("name")
	file, err := c.FormFile("file")
	if err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	// Save file
	if err := c.SaveUploadedFile(file, "./static/demos/inception/images/"+file.Filename); err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	c.Redirect(302, "/datahub")
}
