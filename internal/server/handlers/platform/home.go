package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// Home page handler
func Home(c *gin.Context) {
	inject := data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Home",
	}
	c.Status(200)
	tmpl.Home.Execute(c.Writer, &inject)
}
