package public

import (
	"github.com/gin-gonic/gin"

	"github.com/medtune/beta-platform/internal"
	"github.com/medtune/beta-platform/internal/session"
	"github.com/medtune/beta-platform/internal/tmpl"
	"github.com/medtune/beta-platform/internal/tmpl/data"
)

// SignupSuccess .
func SignupSuccess(c *gin.Context) {
	if !session.GetLoginStatus(c) {
		c.Redirect(302, "error/401")
		return
	}
	c.Status(200)
	tmpl.SignupSuccess.Execute(c.Writer, data.Main{
		Version:   internal.VERSION,
		PageTitle: "Signup ~ Success",
	})
}
