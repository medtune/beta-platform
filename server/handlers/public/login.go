package public

import (
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/service/global"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		// Check session
		// If logged redirect to home
		if session.GetLoginStatus(c) {
			c.Redirect(302, "/home")
			return
		}

		// Render login view
		tmpl.Login.Execute(c.Writer, &data.Main{
			Version:   pkg.VERSION,
			PageTitle: "login",
		})

	} else if c.Request.Method == "POST" {
		logins := jsonutil.LoginData{}
		var err error

		// Parse request body
		err = c.ShouldBindJSON(&logins)
		if err != nil {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		// Service authentificate user
		ok, err := global.AuthUser(&logins)
		if err != nil || !ok {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		// Success
		session.SetLoginStatus(c, true)
		c.JSON(200, jsonutil.Success())
	}
}
