package public

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/access"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

func Signup(c *gin.Context) {
	if c.Request.Method == "GET" {
		// Render signup page
		tmpl.Signup.Execute(c.Writer, &data.Main{
			PageTitle: "Signup",
		})

	} else if c.Request.Method == "POST" {
		signupData := &jsonutil.SignupData{}
		var err error

		// Parse request body (json)
		err = c.ShouldBindJSON(&signupData)
		if err != nil {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		ok, err := access.CreateUser(signupData)
		if err != nil || !ok {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		session.SetLoginStatus(c, true)
		c.JSON(200, jsonutil.Success())
	}
}

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
