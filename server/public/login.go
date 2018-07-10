package public

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/iallabs/medtune-trials/pkg/jsonutil"
	"github.com/iallabs/medtune-trials/pkg/tmpl"
	"github.com/iallabs/medtune-trials/pkg/tmpl/data"
)

func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		s := sessions.Default(c)
		log := s.Get("logged")
		if log != nil && log.(bool) {
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
			s := sessions.Default(c)
			s.Set("username", loginData.Username)
			s.Set("logged", true)
			s.Save()
			c.JSON(200, jsonutil.Success())
		} else {
			c.JSON(200, jsonutil.Fail())
		}
	}
}
