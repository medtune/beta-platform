package platform

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/service/demo"
	"github.com/medtune/beta-platform/internal/store"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// Chexray demo
func Chexray(c *gin.Context) {
	c.Status(200)
	tmpl.DemoChexray.Execute(c.Writer, &data.Main{
		Version:   internal.VERSION,
		PageTitle: "Demo: Chest X-Ray Classification",
	})
}

// ChexrayV2 demo
func ChexrayV2(c *gin.Context) {
	images, err := demo.CollectImagesData("chexray")
	if err != nil {
		images = nil
	}

	cxpbaData, err := store.Agent.GetSpecs(nil)
	if err != nil {
		cxpbaData = nil
	}

	c.Status(200)
	tmpl.DemoChexrayV2.Execute(c.Writer, &data.ChexrayV2Demo{
		Main: data.Main{
			Version:   internal.VERSION,
			PageTitle: "Demo V2: Chest X-Ray Classification",
		},
		Samples:    images,
		Cxpbatable: cxpbaData,
	})
}

// ChexrayUpload .
func ChexrayUpload(c *gin.Context) {
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

	savePath := fmt.Sprintf("./static/demos/%s/images/%s", "chexray", name)

	// Save file
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.Redirect(302, "/error/500")
		return
	}

	c.Redirect(302, "/demos/chexray.v2")
}
