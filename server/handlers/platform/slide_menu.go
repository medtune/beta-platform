package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// SlidesMenu .
func SlidesMenu(c *gin.Context) {
	c.Status(200)
	tmpl.SlidesMenu.Execute(c.Writer, &data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Slides Menu",
	})
}
