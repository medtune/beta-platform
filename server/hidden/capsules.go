package hidden

import (
	"github.com/gin-gonic/gin"
	"github.com/iallabs/medtune-trials/pkg/tmpl"
)

func Inception(c *gin.Context) {
	c.Status(200)
	tmpl.Inception.Execute(c.Writer, nil)
}
