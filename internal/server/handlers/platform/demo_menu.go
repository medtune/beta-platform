package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// DemosMenu page
func DemosMenu(c *gin.Context) {
	c.Status(200)
	tmpl.DemosMenu.Execute(c.Writer, &data.Main{
		Version:   internal.VERSION,
		PageTitle: "Demonstrations",
	})
}
