package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// MedtunePresentation .
func MedtunePresentation(c *gin.Context) {
	c.Status(200)
	tmpl.SlideMedtunePresentation.Execute(c.Writer, &data.Slide{
		Title: "Slide: MedTune Presentation",
	})
}
