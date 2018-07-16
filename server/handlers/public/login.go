package public

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/session"
	"github.com/medtune/beta-platform/pkg/store"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		if session.GetLoginStatus(c) {
			c.Redirect(302, "/home")
		} else {
			tmpl.Login.Execute(c.Writer, &data.Main{
				PageTitle: "login",
			})
		}
	} else if c.Request.Method == "POST" {
		logins := jsonutil.LoginData{}
		if err := c.ShouldBindJSON(&logins); err != nil {
			c.JSON(200, jsonutil.Fail(err))
		}

		if ok, err := store.Agent.AuthentificateUser(logins.Username, logins.Password); err == nil && ok {
			session.SetLoginStatus(c, true)
			c.JSON(200, jsonutil.Success())
		} else {
			c.JSON(200, jsonutil.Fail(fmt.Errorf("username or password incorect")))
		}
	}
}
