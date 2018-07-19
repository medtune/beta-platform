package public

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/secret"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/store"
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

		// Parse json error
		if err := c.ShouldBindJSON(&signupData); err != nil {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		// Password confirmation error
		if signupData.Password != signupData.PasswordConfirm {
			c.JSON(200, jsonutil.Fail(fmt.Errorf("password confirmation error")))
			return
		}

		// Signup secret unknown
		if ok, err := secret.Check("signup", signupData.Secret); err != nil || !ok {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		if err := store.Agent.CreateUser(signupData.Email, signupData.Username, signupData.Password); err == nil {
			session.SetLoginStatus(c, true)
			c.JSON(200, jsonutil.Success())
		} else {
			// Create user error
			c.JSON(200, jsonutil.Fail(fmt.Errorf("email or username already taken")))

		}
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
