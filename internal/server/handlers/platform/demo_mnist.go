package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// Mnist demo
func Mnist(c *gin.Context) {
	c.Status(200)
	tmpl.DemoMnist.Execute(c.Writer, &data.Main{
		Version:   internal.VERSION,
		PageTitle: "Handwritten digits classification",
	})
}
