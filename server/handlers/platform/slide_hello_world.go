package platform

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// HelloWorld .
func HelloWorld(c *gin.Context) {
	c.Status(200)
	tmpl.SlideHelloWorld.Execute(c.Writer, &data.Slide{
		Title: "Slide: Hello world",
	})
}
