package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// MedtunePresentation .
func MedtunePresentation(c *gin.Context) {
	c.Status(200)
	tmpl.SlideMedtunePresentation.Execute(c.Writer, &data.Slide{
		Title: "Slide: MedTune Presentation",
	})
}
