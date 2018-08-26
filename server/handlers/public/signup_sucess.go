package public

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// SignupSuccess .
func SignupSuccess(c *gin.Context) {
	if !session.GetLoginStatus(c) {
		c.Redirect(302, "error/401")
		return
	}
	c.Status(200)
	tmpl.SignupSuccess.Execute(c.Writer, data.Main{
		Version:   pkg.VERSION,
		PageTitle: "Signup ~ Success",
	})
}
