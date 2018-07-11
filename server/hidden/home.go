package hidden

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

func Home(c *gin.Context) {
	inject := data.Main{
		PageTitle: "Home",
	}
	c.Status(200)
	tmpl.Home.Execute(c.Writer, &inject)
}
