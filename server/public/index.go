package public

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

func Index(c *gin.Context) {
	s := sessions.Default(c)
	if v := s.Get("logged"); v != nil && v.(bool) {
		c.Redirect(302, "home")
	} else {
		c.Status(200)
		inject := data.Main{
			PageTitle: "Index",
		}
		tmpl.Index.Execute(c.Writer, &inject)
	}
}
