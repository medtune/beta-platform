package public

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/iallabs/medtune-trials/pkg/jsonutil"
	"github.com/iallabs/medtune-trials/pkg/tmpl"
	"github.com/iallabs/medtune-trials/pkg/tmpl/data"
)

func Signup(c *gin.Context) {
	if c.Request.Method == "GET" {
		tmpl.Signup.Execute(c.Writer, &data.Main{
			PageTitle: "Signup",
		})
	} else if c.Request.Method == "POST" {
		signupData := &jsonutil.SignupData{}
		if err := c.ShouldBindJSON(&signupData); err != nil {
			c.JSON(200, jsonutil.Fail())
		}

		var sign = func(s1, s2, s3, s4 string) (bool, error) {
			fmt.Printf("signup %v %v %v %v\n", s1, s2, s3, s4)
			return false, nil
		}
		if _, err := sign(signupData.Username,
			signupData.Password,
			signupData.PasswordConfirm,
			signupData.Secret); err == nil {

			s := sessions.Default(c)
			s.Set("username", signupData.Username)
			s.Set("logged", true)
			s.Save()
			c.JSON(200, jsonutil.Success())
		} else {
			c.JSON(200, jsonutil.Fail())
		}
	}
}
