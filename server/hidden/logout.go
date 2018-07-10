package hidden

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Set("logged", false)
	s.Save()
	c.Redirect(302, "/index")
}
