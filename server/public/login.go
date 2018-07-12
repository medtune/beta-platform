package public

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/session"
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
		loginData := jsonutil.LoginData{}
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(200, jsonutil.Fail())
		}
		//TODO: Authentification & Validation
		var auth = func(s1, s2 string) (bool, error) {
			fmt.Printf("authentificating %v %v\n", s1, s2)
			return false, nil
		}

		if _, err := auth(loginData.Username, loginData.Password); err == nil {
			session.SetLoginStatus(c, true)
			c.JSON(200, jsonutil.Success())
		} else {
			c.JSON(200, jsonutil.Fail())
		}
	}
}
