package platform

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// Datahub .
func Datahub(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}
	c.Status(200)
	tmpl.DataHub.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Data Hub",
	})
}

// DatahubUpload .
func DatahubUpload(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.Redirect(302, "/index")
		return
	}

	demo := c.PostForm("demo")
	if demo == "" {
		// Default demo
		demo = "inception"
	}

	// read file
	file, err := c.FormFile("file")
	if err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	ext := filepath.Ext(file.Filename)
	// Verify supported extentions
	if ext != ".png" && ext != ".jpeg" && ext != ".jpg" {
		c.Redirect(302, "/error/401")
		return
	}

	// Use new name if not empty else old filename
	name := c.PostForm("name")
	if name == "" {
		name = file.Filename
	} else {
		name += ext
	}

	savePath := fmt.Sprintf("./static/demos/%s/images/%s", demo, name)

	// Save file
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	c.Redirect(302, "/datahub")
}
